package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/structtag"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/go/generator/injection"
	"github.com/mojo-lang/mojo/go/pkg/mojo/mpm"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

func GenerateMojoPackageProtobuf(pkg *lang.Package) (string, error) {
	dir, err := os.MkdirTemp("", "protobuf-mojo-"+pkg.Name)
	if err != nil {
		return "", err
	}

	files := mpm.GetMojoPbFile(pkg.Name)
	if files == nil {
		return "", nil
	}

	for name, file := range files.Parts {
		fileName := path.Join(dir, name)
		err = core.CreateDir(path.Dir(fileName))
		if err != nil {
			return "", err
		}

		if err = os.WriteFile(fileName, file, fs.ModePerm); err != nil {
			return "", err
		}
	}
	return dir, nil
}

func isGoCommandExist(command string) bool {
	cmd := exec.Command("which", command)
	if out, err := cmd.Output(); err != nil {
		if goPath := os.Getenv("GOPATH"); len(goPath) > 0 {
			return core.IsExist(filepath.Join(goPath, "bin", command))
		} else {
			return false
		}
	} else {
		return len(out) > 0 && !strings.Contains(string(out), "not found")
	}
}

func prepareEnv() error {
	if !isGoCommandExist("protoc-gen-go") {
		err := exec.Command("go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go@latest").Run()
		if err != nil {
			return err
		}
	}
	if !isGoCommandExist("protoc-gen-go-grpc") {
		err := exec.Command("go", "install", "google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest").Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func ProtocGenGo(dir string, pkg *lang.Package, files []*descriptor.File) (util.GeneratedFiles, error) {
	if err := prepareEnv(); err != nil {
		return nil, err
	}

	var tempPbDirs []string
	cmd := exec.Command("protoc", "-I.")
	for _, dep := range pkg.ResolvedDependencies {
		wd := dep.GetExtraString("workingDir")
		p := dep.GetExtraString("path")
		if len(wd) == 0 && len(p) == 0 {
			pbDir, err := GenerateMojoPackageProtobuf(dep)
			if err != nil {
				logs.Errorw("failed to generate mojo package's protobuf files", "package", dep.FullName, "error", err)
				return nil, err
			}
			if len(pbDir) > 0 {
				tempPbDirs = append(tempPbDirs, pbDir)
				cmd.Args = append(cmd.Args, "--proto_path="+pbDir)
			}
		} else {
			cmd.Args = append(cmd.Args, "--proto_path="+path.Join(wd, p, "protobuf"))
		}
	}
	defer func() {
		for _, d := range tempPbDirs {
			if err := os.RemoveAll(d); err != nil {
				logs.Warnw("failed to remove all the files", "dir", d, "error", err)
			}
		}
	}()

	cmd.Dir = path.Join(dir, "protobuf")

	outDir := filepath.Join(dir, "go.out")
	if core.IsExist(outDir) {
		if err := os.RemoveAll(outDir); err != nil {
			return nil, err
		}
	}
	if err := core.CreateDir(outDir); err != nil {
		return nil, err
	}
	defer func() {
		if err := os.RemoveAll(outDir); err != nil {
			logs.Warnw("failed to remove all the files in the dir", "dir", outDir, "error", err)
		}
	}()

	cmd.Args = append(cmd.Args, fmt.Sprintf("--go_out=%s", outDir))
	cmd.Args = append(cmd.Args, "--go_opt=paths=source_relative")
	cmd.Args = append(cmd.Args, fmt.Sprintf("--go-grpc_out=%s", outDir))
	cmd.Args = append(cmd.Args, "--go-grpc_opt=paths=source_relative")

	for _, file := range files {
		if !file.IsEmpty() {
			cmd.Args = append(cmd.Args, file.GetName())
		}
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_, err := cmd.Output()
	if err != nil {
		logs.Errorw(fmt.Sprintf("failed to run protoc cmd %s", stderr.String()), "cmd", cmd.String())
		return nil, err
	}
	logs.Debugw("finish to run protoc cmd", "warning", stderr.String(), "cmd", cmd.String())

	var genFiles util.GeneratedFiles
	err = filepath.WalkDir(outDir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			if rel, err := filepath.Rel(outDir, path); err != nil {
				return err
			} else {
				if content, err := os.ReadFile(path); err != nil {
					return err
				} else {
					genFiles = append(genFiles, &util.GeneratedFile{
						Name:    filepath.Join("pkg", rel),
						Content: string(content),
					})
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	for _, f := range genFiles {
		xxxSkip := []string{"gorm", "xml", "bson"}
		injector := injection.NewTagInjector(xxxSkip, nil)
		injector.RegisterTagChanger(func(ctx context.Context, field *ast.Field, tags *structtag.Tags) {
			structName := injection.ToMojoStructName(injection.GetStructName(ctx))
			fieldName := injection.ToMojoFieldName(field.Names[0].Name)
			valDecl := injection.GetStructField(pkg, structName, fieldName)

			if alias, _ := valDecl.GetStringAttribute(core.AliasAttributeName); len(alias) > 0 {
				_ = tags.Set(&structtag.Tag{
					Key:     "json",
					Name:    strcase.ToLowerCamel(alias),
					Options: []string{"omitempty"},
				})
			} else {
				_ = tags.Set(&structtag.Tag{
					Key:     "json",
					Name:    strcase.ToLowerCamel(fieldName),
					Options: []string{"omitempty"},
				})
			}

			if ignore, _ := valDecl.GetBoolAttribute(db.IgnoreAttributeFullName); ignore {
				_ = AddTagsOptions(tags, "db", fieldName, "-")
				_ = AddTagsOptions(tags, "gorm", "", "-")
			} else {
				if dbJson, _ := valDecl.GetBoolAttribute(db.JSONAttributeFullName); dbJson {
					_ = AddTagsOptions(tags, "db", fieldName, "json")
				}
				if dbExplode, err := valDecl.GetStringAttribute(db.ExplodeAttributeFullName); err == nil {
					if len(dbExplode) > 0 {
						_ = AddTagsOptions(tags, "db", fieldName, "explode="+dbExplode)
					} else {
						_ = AddTagsOptions(tags, "db", "", "explode")
					}
				}

				addPrimaryKey := func(key string) {
					_ = AddTagsOptions(tags, "db", fieldName, key+"=true")
					_ = AddTagsOptions(tags, "gorm", "", "primaryKey")
				}

				if v, err := valDecl.GetBoolAttribute(db.PrimaryKeyAttributeFullName); err == nil && v {
					addPrimaryKey("primaryKey")
				} else if v, err = valDecl.GetBoolAttribute(db.KeyAttributeFullName); err == nil && v {
					addPrimaryKey("key")
				} else if v, err = valDecl.GetBoolAttribute(core.KeyAttributeName); err == nil && v {
					addPrimaryKey("key")
				}

				if valDecl.HasAttribute(db.IndexAttributeFullName) {
					_ = AddTagsOptions(tags, "db", fieldName, "index=true")
					_ = AddTagsOptions(tags, "gorm", "", "index")
				}

				if foreignKey, _ := valDecl.GetStringAttribute(db.ForeignKeyAttributeFullName); len(foreignKey) > 0 {
					_ = AddTagsOptions(tags, "db", fieldName, "foreignKey="+foreignKey)
					_ = AddTagsOptions(tags, "gorm", "", "foreignKey:"+strcase.ToCamel(foreignKey))
				}

				if maxLength, _ := valDecl.GetIntegerAttribute(core.MaxLengthAttributeName); maxLength > 0 {
					if valDecl.Type.GetFullName() == core.StringTypeFullName {
						if maxLength < 256 {
							_ = AddTagsOptions(tags, "gorm", "", "type:varchar("+strconv.Itoa(int(maxLength))+")")
						}
					}
				}
			}

			if tag, _ := tags.Get("gorm"); tag != nil {
				if len(tag.Options) > 0 {
					options := strings.Join(tag.Options, ";")
					tag.Name = tag.Name + ";" + options
					tag.Options = []string{}
				}
			}
		})

		bs, err := injector.Inject(f.Name, []byte(f.Content))
		if err == nil {
			f.Content = string(bs)
		}

		bs, err = injection.NewDbJSONInjector().Inject(f.Name, []byte(f.Content))
		if err == nil {
			f.Content = string(bs)
		}

		// disable the renamer, which will make the message failed to convert to Proto.Message.V1 interface.
		// bs, err = newStringMethodRenamer().Inject(f.Name, []byte(f.Content))
		// if err == nil {
		// 	f.Content = string(bs)
		// }
	}

	return genFiles, nil
}

func AddTagsOptions(tags *structtag.Tags, key string, name string, options ...string) error {
	if len(name) == 0 && len(options) == 0 {
		return nil
	}

	if tag, _ := tags.Get(key); tag == nil {
		if len(name) == 0 {
			name = options[0]
			options = options[1:]
		}

		err := tags.Set(&structtag.Tag{
			Key:  key,
			Name: name,
		})
		if err != nil {
			return err
		}

		if len(options) > 0 {
			tags.AddOptions(key, options...)
		}
	} else {
		tags.AddOptions(key, options...)
	}
	return nil
}

// type stringMethodRenamer struct {
// 	injection.Injector
// }
//
// func newStringMethodRenamer() *stringMethodRenamer {
// 	in := &stringMethodRenamer{}
// 	in.OnFunction = in.onFunction
// 	return in
// }
//
// func (r *stringMethodRenamer) onFunction(ctx context.Context, function *ast.FuncDecl, areaAppender func(area injection.Area)) {
// 	if function.Recv != nil && len(function.Recv.List) > 0 {
// 		name := function.Name
// 		if name.Name == "String" {
// 			areaAppender(&injection.TextArea{
// 				Start:   name.Pos(),
// 				End:     name.End(),
// 				Content: "ToText",
// 			})
// 		}
// 	}
// }
