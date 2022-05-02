package generator

import (
    "bytes"
    "context"
    "fmt"
    "github.com/fatih/structtag"
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/db/go/pkg/mojo/db"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/injection"
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
    "go/ast"
    "io/fs"
    "io/ioutil"
    "os"
    "os/exec"
    path2 "path"
    "path/filepath"
    "strings"
)

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

func ProtocGenGo(path string, pkg *lang.Package, files []*descriptor.File) (util.GeneratedFiles, error) {
    if err := prepareEnv(); err != nil {
        return nil, err
    }

    cmd := exec.Command("protoc", "-I.")
    for _, dep := range pkg.ResolvedDependencies {
        wd := dep.GetExtraString("workingDir")
        p := dep.GetExtraString("path")
        cmd.Args = append(cmd.Args, "--proto_path="+path2.Join(wd, p, "protobuf"))
    }

    cmd.Dir = path2.Join(path, "protobuf")

    outDir := filepath.Join(path, "go.out")
    if core.IsExist(outDir) {
        if err := os.RemoveAll(outDir); err != nil {
            return nil, err
        }
    }
    if err := core.CreateDir(outDir); err != nil {
        return nil, err
    }
    defer func() {
        os.RemoveAll(outDir)
    }()

    cmd.Args = append(cmd.Args, fmt.Sprintf("--go_out=%s", outDir))
    cmd.Args = append(cmd.Args, "--go_opt=paths=source_relative")
    cmd.Args = append(cmd.Args, fmt.Sprintf("--go-grpc_out=%s", outDir))
    cmd.Args = append(cmd.Args, "--go-grpc_opt=paths=source_relative")

    for _, file := range files {
        cmd.Args = append(cmd.Args, file.GetName())
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
                if content, err := ioutil.ReadFile(path); err != nil {
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
                tags.Set(&structtag.Tag{
                    Key:     "json",
                    Name:    strcase.ToLowerCamel(alias),
                    Options: []string{"omitempty"},
                })
            } else {
                tags.Set(&structtag.Tag{
                    Key:     "json",
                    Name:    strcase.ToLowerCamel(fieldName),
                    Options: []string{"omitempty"},
                })
            }

            if ignore, _ := valDecl.GetBoolAttribute(db.IgnoreAttributeFullName); ignore {
                tags.Set(&structtag.Tag{
                    Key:  "db",
                    Name: "-",
                })
                tags.Set(&structtag.Tag{
                    Key:  "gorm",
                    Name: "-",
                })
            } else {
                if dbJson, _ := valDecl.GetBoolAttribute(db.JSONAttributeFullName); dbJson {
                    if tag, _ := tags.Get("db"); tag == nil {
                        tags.Set(&structtag.Tag{
                            Key:  "db",
                            Name: fieldName,
                        })
                    }
                    tags.AddOptions("db", "json")
                }
                if dbExplode, err := valDecl.GetStringAttribute(db.ExplodeAttributeFullName); err == nil {
                    if tag, _ := tags.Get("db"); tag == nil {
                        tags.Set(&structtag.Tag{
                            Key:  "db",
                            Name: fieldName,
                        })
                    }
                    if len(dbExplode) > 0 {
                        tags.AddOptions("db", "explode="+dbExplode)
                    } else {
                        tags.AddOptions("db", "explode")
                    }
                }

                addPrimaryKey := func(key string) {
                    if tag, _ := tags.Get("db"); tag == nil {
                        tags.Set(&structtag.Tag{
                            Key:  "db",
                            Name: fieldName,
                        })
                    }
                    tags.AddOptions("db", key+"=true")
                    tags.Set(&structtag.Tag{Key: "gorm", Name: "primaryKey"})
                }

                if v, err := valDecl.GetBoolAttribute(db.PrimaryKeyAttributeFullName); err == nil && v {
                    addPrimaryKey("primaryKey")
                } else if v, err = valDecl.GetBoolAttribute(db.KeyAttributeFullName); err == nil && v {
                    addPrimaryKey("key")
                } else if v, err = valDecl.GetBoolAttribute(core.KeyAttributeName); err == nil && v {
                    addPrimaryKey("key")
                }

                if valDecl.HasAttribute(db.IndexAttributeFullName) {
                    if tag, _ := tags.Get("db"); tag == nil {
                        tags.Set(&structtag.Tag{
                            Key:  "db",
                            Name: fieldName,
                        })
                    }
                    tags.AddOptions("db", "index=true")
                    tags.Set(&structtag.Tag{Key: "gorm", Name: "index"})
                }

                if foreignKey, _ := valDecl.GetStringAttribute(db.ForeignKeyAttributeFullName); len(foreignKey) > 0 {
                    if tag, _ := tags.Get("db"); tag == nil {
                        tags.Set(&structtag.Tag{
                            Key:  "db",
                            Name: fieldName,
                        })
                    }
                    tags.AddOptions("db", "foreignKey="+foreignKey)
                    tags.Set(&structtag.Tag{Key: "gorm", Name: "foreignKey:" + strcase.ToCamel(foreignKey)})
                }
            }
        })

        bytes, err := injector.Inject(f.Name, []byte(f.Content))
        if err == nil {
            f.Content = string(bytes)
        }

        bytes, err = injection.NewDbJSONInjector().Inject(f.Name, []byte(f.Content))
        if err == nil {
            f.Content = string(bytes)
        }
    }

    return genFiles, nil
}
