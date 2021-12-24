package compiler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/fatih/structtag"
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	ggdescriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	_ "github.com/gogo/protobuf/protoc-gen-gogo/grpc"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/go/compiler/inject"
	"github.com/mojo-lang/mojo/go/pkg/mojo/util"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
	"go/ast"
	"os/exec"
	path2 "path"
	"strings"
)

func ProtocGo(path string, pkg *lang.Package, files []*descriptor.FileDescriptor) (util.GeneratedFiles, error) {
	if pkg == nil {
		return nil, errors.New("")
	}

	cmd := exec.Command("protoc", "-I.")
	for _, dep := range pkg.ResolvedDependencies {
		wd := dep.GetExtraString("workingDir")
		p := dep.GetExtraString("path")
		cmd.Args = append(cmd.Args, "--proto_path="+path2.Join(wd, p, "protobuf"))
	}

	cmd.Args = append(cmd.Args, "--include_source_info")
	cmd.Args = append(cmd.Args, "--include_imports")

	cmd.Dir = path2.Join(path, "protobuf")

	//cmd.Args = append(cmd.Args, "--go_out=.")
	//cmd.Args = append(cmd.Args, "--gogo_out=plugins=grpc:.")
	cmd.Args = append(cmd.Args, "--descriptor_set_out=/dev/stdout")

	var outFiles util.GeneratedFiles

	/// gogoproto will clean the plugin after the first run `command.Generate(req)`
	/// make sure the service files will use the grpc plugin
	fileGroups := make(map[string][]*descriptor.FileDescriptor)

	packageGroups := make(map[string]bool)
	for _, file := range files {
		if file.HasService() {
			packageGroups[*file.Package] = true
			fileGroups[*file.Package] = append(fileGroups[*file.Package], file)
		}
	}
	for _, file := range files {
		if file.HasService() {
			continue
		}

		if in, ok := packageGroups[*file.Package]; ok && in {
			fileGroups[*file.Package] = append(fileGroups[*file.Package], file)
		} else {
			fileGroups["-others-"] = append(fileGroups["-others-"], file)
		}
	}

	for _, group := range fileGroups {
		if len(group) == 0 {
			continue
		}

		fileCmd := &exec.Cmd{
			Path: cmd.Path,
			Args: cmd.Args,
			Dir:  cmd.Dir,
		}

		for _, file := range group {
			fileCmd.Args = append(fileCmd.Args, *file.Name)
		}

		var stderr bytes.Buffer
		fileCmd.Stderr = &stderr
		out, err := fileCmd.Output()
		if err != nil {
			logs.Errorw(fmt.Sprintf("failed to run protoc cmd %s", stderr.String()), "cmd", fileCmd.String())
			return nil, err
		}
		logs.Debugw("finish to run protoc cmd", "warning", stderr.String(), "cmd", fileCmd.String())

		fs, err := UnmarshalFiles(out)
		if err != nil {
			return nil, err
		}

		vanity.ForEachFile(fs, vanity.TurnOffGogoImport)
		//vanity.ForEachFile(fs, vanity.TurnOnSizerAll)
		//vanity.ForEachFile(fs, vanity.TurnOnUnmarshalerAll)
		//vanity.ForEachFile(fs, vanity.TurnOffGogoImport)
		vanity.ForEachFieldInFiles(fs, JsonTagLowerCamelCase)
		vanity.ForEachFieldInFiles(fs, CompileDBFieldOptions)

		parameter := "plugins=grpc"
		req := &plugin.CodeGeneratorRequest{
			ProtoFile: fs,
			Parameter: &parameter,
		}
		for _, f := range fs {
			for _, file := range group {
				if *f.Name == *file.Name {
					req.FileToGenerate = append(req.FileToGenerate, *f.Name)
				}
			}
		}
		resp := command.Generate(req)
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		for _, f := range resp.File {
			if f.Name != nil && f.Content != nil {
				for _, file := range group {
					goPath := "pkg/" + lang.PackageNameToPath(*file.Package)
					pos := strings.Index(*f.Name, goPath)
					if pos >= 0 {
						name := (*f.Name)[pos:]
						outFiles = append(outFiles, &util.GeneratedFile{
							Name:    name,
							Content: *f.Content,
						})
					}
				}
			}
		}
	}

	for _, f := range outFiles {
		xxxSkip := []string{"gorm", "xml", "bson"}
		injector := NewTagInjector(xxxSkip, nil)
		injector.RegisterTagChanger(func(ctx context.Context, field *ast.Field, tags *structtag.Tags) {
			if dbTag, _ := tags.Get("db"); dbTag != nil && inject.HasTagOption(dbTag, "ignore") {
				tags.Set(&structtag.Tag{
					Key:  "gorm",
					Name: "-",
				})
			}
		})

		bytes, err := injector.Inject(f.Name, []byte(f.Content))
		if err == nil {
			f.Content = string(bytes)
		}

		bytes, err = NewDbJSONInjector().Inject(f.Name, []byte(f.Content))
		if err == nil {
			f.Content = string(bytes)
		}
	}

	return outFiles, nil
}

func JsonTagLowerCamelCase(field *ggdescriptor.FieldDescriptorProto) {
	//if field.IsRepeated() || field.IsMessage() {
	//	return
	//}
	//if field.DefaultValue != nil {
	//	return
	//}

	value := strcase.ToLowerCamel(*field.Name)
	value += ",omitempty"

	CompileAliasFieldOptions(field)
	SetStringFieldOption(gogoproto.E_Jsontag, value)(field)
}

func CompileAliasFieldOptions(field *ggdescriptor.FieldDescriptorProto) {
	value, err := proto.GetUnsafeExtension(field.Options, mojo.E_Alias.Field)
	if err != nil || value == nil {
		return
	}
	if v := value.(*string); v != nil && len(*v) > 0 {
		SetStringFieldOption(gogoproto.E_Jsontag, *v+",omitempty")(field)
	}
}

func CompileDBFieldOptions(field *ggdescriptor.FieldDescriptorProto) {
	tag := "db:\""

	first := true
	if value, err := proto.GetUnsafeExtension(field.Options, mojo.E_DbJson.Field); err != nil || value == nil {
	} else {
		tag += "json"
		first = false
	}

	if value, err := proto.GetUnsafeExtension(field.Options, mojo.E_DbExplode.Field); err != nil || value == nil {
	} else {
		v := value.(*string)
		if !first {
			tag += ","
		}
		if v != nil && len(*v) > 0 {
			tag += "explode=" + *v
		} else {
			tag += "explode"
		}
		first = false
	}

	if value, err := proto.GetUnsafeExtension(field.Options, mojo.E_DbIgnore.Field); err != nil || value == nil {
	} else {
		if !first {
			tag += ","
		}
		tag += "ignore"
		first = false
	}

	tag += "\""

	if !first {
		SetStringFieldOption(gogoproto.E_Moretags, tag)(field)
	}
}

func SetStringFieldOption(extension *proto.ExtensionDesc, value string) func(field *ggdescriptor.FieldDescriptorProto) {
	return func(field *ggdescriptor.FieldDescriptorProto) {
		if FieldHasStringExtension(field, extension) {
			return
		}
		if field.Options == nil {
			field.Options = &ggdescriptor.FieldOptions{}
		}
		if err := proto.SetExtension(field.Options, extension, &value); err != nil {
			panic(err)
		}
	}
}

func FieldHasStringExtension(field *ggdescriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) bool {
	if field.Options == nil {
		return false
	}
	value, err := proto.GetExtension(field.Options, extension)
	if err != nil {
		return false
	}
	if value == nil {
		return false
	}
	if value.(*string) == nil {
		return false
	}
	return true
}

func UnmarshalFiles(bytes []byte) ([]*ggdescriptor.FileDescriptorProto, error) {
	fileDesc := &ggdescriptor.FileDescriptorSet{}
	if err := proto.Unmarshal(bytes, fileDesc); err != nil {
		return nil, err
	}
	return fileDesc.File, nil
}

var E_EnumAlias = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         72001,
	Name:          "mojo.enum_alias",
	Tag:           "varint,72001,opt,name=enum_alias",
	Filename:      "mojo/mojo.proto",
}

var E_EnumvalueAlias = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         76001,
	Name:          "mojo.enumvalue_alias",
	Tag:           "bytes,76001,opt,name=enumvalue_alias",
	Filename:      "mojo/mojo.proto",
}

var E_GettersAll = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         73001,
	Name:          "mojo.getters_all",
	Tag:           "varint,73001,opt,name=getters_all",
	Filename:      "mojo/mojo.proto",
}

var E_Getters = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         74001,
	Name:          "mojo.getters",
	Tag:           "varint,74001,opt,name=getters",
	Filename:      "mojo/mojo.proto",
}

var E_Alias = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         75001,
	Name:          "mojo.alias",
	Tag:           "bytes,75001,opt,name=alias",
	Filename:      "mojo/mojo.proto",
}

var E_DbIgnore = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         75100,
	Name:          "mojo.db_ignore",
	Tag:           "varint,75100,opt,name=db_ignore",
	Filename:      "mojo/mojo.proto",
}
var E_DbJson = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         75101,
	Name:          "mojo.db_json",
	Tag:           "varint,75101,opt,name=db_json",
	Filename:      "mojo/mojo.proto",
}
var E_DbIndex = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         75102,
	Name:          "mojo.db_index",
	Tag:           "bytes,75102,opt,name=db_index",
	Filename:      "mojo/mojo.proto",
}
var E_DbExplode = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         75103,
	Name:          "mojo.db_explode",
	Tag:           "bytes,75103,opt,name=db_explode",
	Filename:      "mojo/mojo.proto",
}
var E_DbReference = &proto.ExtensionDesc{
	ExtendedType:  (*ggdescriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         75104,
	Name:          "mojo.db_reference",
	Tag:           "bytes,75104,opt,name=db_reference",
	Filename:      "mojo/mojo.proto",
}

func init() {
	proto.RegisterExtension(E_EnumAlias)
	proto.RegisterExtension(E_EnumvalueAlias)
	proto.RegisterExtension(E_GettersAll)
	proto.RegisterExtension(E_Getters)
	proto.RegisterExtension(E_Alias)
	proto.RegisterExtension(E_DbIgnore)
	proto.RegisterExtension(E_DbJson)
	proto.RegisterExtension(E_DbIndex)
	proto.RegisterExtension(E_DbExplode)
	proto.RegisterExtension(E_DbReference)
}
