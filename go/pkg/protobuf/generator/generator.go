package generator

import (
    "bytes"
    "fmt"
    "log"
    "os"
    path2 "path"
    "sort"
    "strings"

    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/types/pluginpb"
)

// Generator is the type whose methods generate the output, stored in the associated response structure.
type Generator struct {
    *bytes.Buffer

    Request  *pluginpb.CodeGeneratorRequest
    Response *pluginpb.CodeGeneratorResponse // The output.

    Param        map[string]string // Command-line parameters.
    ImportPrefix string            // String to prefix to imported package file names.
    ImportMap    map[string]string // Mapping from .proto file name to import path

    Pkg map[string]string // The names under which we import support packages

    genFiles []*descriptor.File // Those files we will generate output for.

    allFiles       []*descriptor.File          // All files in the tree
    allFilesByName map[string]*descriptor.File // All files by filename.

    file *descriptor.File // The file we are compiling now.

    indent      string
    writeOutput bool
}

// New creates a new generator and allocates the request and response protobufs.
func New(files []*descriptor.File) *Generator {
    g := new(Generator)
    g.Buffer = new(bytes.Buffer)
    g.Request = new(pluginpb.CodeGeneratorRequest)
    g.Response = new(pluginpb.CodeGeneratorResponse)

    if len(files) > 0 {
        for _, file := range files {
            g.allFiles = append(g.allFiles, file)
        }
        g.genFiles = g.allFiles
    }

    return g
}

func (g *Generator) GetGeneratedFiles() []*descriptor.File {
    fileIndex := make(map[string]bool)
    for _, file := range g.Response.File {
        fileIndex[*file.Name] = true
    }

    var files []*descriptor.File
    for _, file := range g.genFiles {
        if fileIndex[file.GetName()] {
            files = append(files, file)
        }
    }

    return files
}

// Error reports a problem, including an error, and exits the program.
func (g *Generator) Error(err error, msgs ...string) {
    s := strings.Join(msgs, " ") + ":" + err.Error()
    log.Print("protoc-gen-proto: error:", s)
    os.Exit(1)
}

// Fail reports a problem and exits the program.
func (g *Generator) Fail(msgs ...string) {
    s := strings.Join(msgs, " ")
    log.Print("protoc-gen-proto: error:", s)
    os.Exit(1)
}

// printAtom prints the (atomic, non-annotation) argument to the generated output.
func (g *Generator) printAtom(v interface{}) {
    switch v := v.(type) {
    case string:
        g.WriteString(v)
    case *string:
        g.WriteString(*v)
    case bool:
        fmt.Fprint(g, v)
    case *bool:
        fmt.Fprint(g, *v)
    case int:
        fmt.Fprint(g, v)
    case *int32:
        fmt.Fprint(g, *v)
    case *int64:
        fmt.Fprint(g, *v)
    case float64:
        fmt.Fprint(g, v)
    case *float64:
        fmt.Fprint(g, *v)
    default:
        g.Fail(fmt.Sprintf("unknown type in printer: %T", v))
    }
}

// P prints the arguments to the generated output.  It handles strings and int32s, plus
// handling indirections because they may be *string, etc.  Any inputs of type AnnotatedAtoms may emit
// annotations in a .meta file in addition to outputting the atoms themselves (if g.annotateCode
// is true).
func (g *Generator) P(strs ...interface{}) {
    if !g.writeOutput {
        return
    }
    g.WriteString(g.indent)
    g.S(strs...)
    g.WriteByte('\n')
}

func (g *Generator) S(strs ...interface{}) {
    if !g.writeOutput {
        return
    }
    for _, str := range strs {
        switch v := str.(type) {
        default:
            g.printAtom(v)
        }
    }
}

// deprecationComment is the standard comment added to deprecated
// messages, fields, enums, and enum values.
var deprecationComment = "// Deprecated: Do not use."

// PrintComments prints any comments from the source .proto file.
// The path is a comma-separated list of integers.
// It returns an indication of whether any comments were printed.
// See descriptor.proto for its format.
func (g *Generator) PrintComments(path string) bool {
    if !g.writeOutput {
        return false
    }
    if c, ok := g.makeComments(path); ok {
        g.P(c)
        return true
    }
    return false
}

// makeComments generates the comment string for the field, no "\n" at the end
func (g *Generator) makeComments(path string) (string, bool) {
    //loc, ok := g.file.Comments[path]
    //if !ok {
    //return "", false
    //}
    w := new(bytes.Buffer)
    //nl := ""
    //for _, line := range strings.Split(strings.TrimSuffix(loc.GetLeadingComments(), "\n"), "\n") {
    //    fmt.Fprintf(w, "%s//%s", nl, line)
    //    nl = "\n"
    //}
    return w.String(), true
}

// In Indents the output one tab stop.
func (g *Generator) In() { g.indent += "    " }

// Out unindents the output one tab stop.
func (g *Generator) Out() {
    if len(g.indent) > 0 {
        g.indent = g.indent[4:]
    }
}

// WrapTypes walks the incoming data, wrapping DescriptorProtos, EnumProtos
// and FileProtos into file-referenced objects within the Generator.
// It also creates the list of files to generate and so should be called before GenerateAllFiles.
func (g *Generator) WrapTypes() {
    g.allFiles = make([]*descriptor.File, 0, len(g.Request.ProtoFile))
    g.allFilesByName = make(map[string]*descriptor.File, len(g.allFiles))
    genFileNames := make(map[string]bool)
    for _, n := range g.Request.FileToGenerate {
        genFileNames[n] = true
    }
    for _, f := range g.Request.ProtoFile {
        file := descriptor.NewFileFrom(f)
        g.allFiles = append(g.allFiles, file)
    }

    g.genFiles = make([]*descriptor.File, 0, len(g.Request.FileToGenerate))
    for _, fileName := range g.Request.FileToGenerate {
        fd := g.allFilesByName[fileName]
        if fd == nil {
            g.Fail("could not find file named", fileName)
        }
        g.genFiles = append(g.genFiles, fd)
    }
}

func (g *Generator) removeGeneratedDir(dir string) {
    var paths []string
    for _, file := range g.Response.File {
        path := path2.Dir(path2.Join(dir, *file.Name))
        paths = append(paths, path)
    }

    sort.Strings(paths)
    rootPath := ""
    for _, path := range paths {
        if len(rootPath) == 0 || !strings.HasPrefix(path, rootPath) {
            rootPath = path
            if core.IsExist(rootPath) {
                util.ClearFiles(rootPath, ".proto")
            }
        }
    }
}

func (g *Generator) WriteAllFiles(dir string) error {
    if g.Response == nil {
        return nil
    }

    var files util.GeneratedFiles
    for _, file := range g.Response.File {
        if file.Name != nil && file.Content != nil {
            files = append(files, &util.GeneratedFile{
                Name:              *file.Name,
                Content:           *file.Content,
                SkipNoneGenerated: true,
            })
        } else {
            logs.Warn("meet an empty file!")
        }
    }
    guard := &util.PathGuard{
        OnlyClearGenerated: true,
        Suffixes:           []string{".proto"},
    }

    for _, file := range files {
        if err := file.WriteTo(dir, guard); err != nil {
            return err
        }
    }

    return nil
}

// GenerateAllFiles generates the output for all the files we're outputting.
func (g *Generator) GenerateAllFiles() *Generator {
    // Generate the output. The generator runs for every file, even the files
    // that we don't generate output for, so that we can collate the full list
    // of exported symbols to support public imports.
    genFileMap := make(map[*descriptor.File]bool, len(g.genFiles))
    for _, file := range g.genFiles {
        genFileMap[file] = true
    }
    for _, file := range g.allFiles {
        g.Reset()
        g.writeOutput = genFileMap[file]
        if hasContent := g.generate(file); !hasContent {
            continue
        }
        if !g.writeOutput {
            continue
        }
        g.Response.File = append(g.Response.File, &pluginpb.CodeGeneratorResponse_File{
            Name:    proto.String(file.GetName()),
            Content: proto.String(g.String()),
        })
    }

    return g
}

// Fill the response protocol buffer with the generated output for all the files we're
// supposed to generate.
func (g *Generator) generate(file *descriptor.File) bool {
    g.file = file

    isEmpty := true
    for _, enum := range g.file.Enums {
        if enum.Parent == nil {
            g.P()
            g.generateEnum(enum)
            isEmpty = false
        }
    }

    for _, message := range g.file.Messages {
        if message.Parent == nil {
            if message.IsMapEntry() {
                continue
            }

            if isSystemMessage(message) && file.GetPackageName() == "mojo.core" {
                continue
            }

            g.P()
            g.generateMessage(message)
            isEmpty = false
        }
    }

    for _, service := range g.file.Services {
        g.P()
        g.generateService(service)
        isEmpty = false
    }

    // Run the plugins before the imports so we know which imports are necessary.
    //g.runPlugins(file)

    // Generate header and imports last, though they appear first in the output.
    rem := g.Buffer
    g.Buffer = new(bytes.Buffer)
    g.generateHeader()

    if len(g.file.GetDependencies()) > 0 {
        g.P()
        g.generateImports()
    }

    if g.file.HasOptions() {
        options := g.file.GetOptions()
        if g.file.HasGoOptions() {
            g.P()
            g.WriteString("option go_package = \"")
            g.WriteString(*options.GoPackage)
            g.WriteString("\";")
        }

        if g.file.HasJavaOptions() {
            if options.JavaMultipleFiles != nil {
                g.P()
                g.WriteString("option java_multiple_files = ")
                if *options.JavaMultipleFiles {
                    g.WriteString("true;")
                } else {
                    g.WriteString("false;")
                }
            }

            if options.JavaOuterClassname != nil {
                g.P()
                g.WriteString("option java_outer_classname = \"")
                g.WriteString(*options.JavaOuterClassname)
                g.WriteString("\";")
            }

            if options.JavaPackage != nil {
                g.P()
                g.WriteString("option java_package = \"")
                g.WriteString(*options.JavaPackage)
                g.WriteString("\";")
            }
        }
    }

    g.P()
    g.Write(rem.Bytes())

    return !isEmpty
}

// Generate the header, including package definition
func (g *Generator) generateHeader() {
    g.P("// Code generated by mojo. DO NOT EDIT.")
    if g.file.IsDeprecated() {
        g.P("//")
        g.P("// ", g.file.GetName(), " is a deprecated file.")
    }

    g.P()

    syntax := "proto3"
    if !g.file.IsProto3() {
        syntax = "proto2"
    }
    g.P("syntax = \"", syntax, "\";")

    if g.file.Package != nil {
        g.P()
        g.P("package ", g.file.Package, ";")
    } else {
        // error
    }
}

func (g *Generator) generateImports() {
    for _, imp := range g.file.GetDependencies() {
        g.P("import \"", imp, "\";")
    }
}

// Generate the enum definitions for this Enum.
func (g *Generator) generateEnum(enum *descriptor.Enum) {
    if enum.IsDeprecated() {
        g.P("// ", deprecationComment)
    }

    g.P("enum ", enum.GetName(), " {")
    g.In()
    for _, value := range enum.Values {
        if value.IsDeprecated() {
            g.P()
            g.P("// ", deprecationComment)
        }

        g.P(value.GetName(), "=", value.GetNumber(), ";")
    }
    g.Out()
    g.P("}")
}

// Generate the type, methods and default constant definitions for this Descriptor.
func (g *Generator) generateMessage(message *descriptor.Message) {
    g.P("message ", message.GetName(), " {")
    g.In()

    // Build a structure more suitable for generating the text in one pass
    for _, enum := range message.Enums {
        g.P()
        g.generateEnum(enum)
    }

    for _, msg := range message.Messages {
        if msg.IsMapEntry() {
            continue
        }

        g.P()
        g.generateMessage(msg)
    }

    printField := func(field *descriptor.Field) {
        g.WriteString(g.indent)
        if field.IsMessageType() {
            desc := message.GetInnerMessage(field.GetTypeName())
            if desc.IsMapEntry() {
                keyType := desc.Fields[0].GetTypeName()
                valType := desc.Fields[1].GetTypeName()
                g.S("map<", keyType, ", ", valType, "> ", field.GetName(), " = ", field.GetNumber())
            } else {
                if field.IsRepeated() {
                    g.S("repeated ", field.GetTypeName(), " ", field.GetName(), " = ", field.GetNumber())
                } else {
                    g.S(field.GetTypeName(), " ", field.GetName(), " = ", field.GetNumber())
                }
            }
        } else if field.IsRepeated() {
            g.S("repeated ", field.GetTypeName(), " ", field.GetName(), " = ", field.GetNumber())
        } else {
            g.S(field.GetTypeName(), " ", field.GetName(), " = ", field.GetNumber())
        }

        if field.HasOption() {
            fieldOptions := mojo.FieldOptionsExtensions()
            first := true
            buffer := bytes.NewBuffer(nil)
            for _, option := range fieldOptions {
                switch option.TypeDescriptor().Kind() {
                case protoreflect.BoolKind:
                    if v := field.GetBoolOption(option); v != nil {
                        if !first {
                            buffer.WriteString(", ")
                        }
                        if *v {
                            buffer.WriteString(fmt.Sprint("(", mojo.GetOptionFullName(option), ")=true"))
                        } else {
                            buffer.WriteString(fmt.Sprint("(", mojo.GetOptionFullName(option), ")=false"))
                        }
                        first = false
                    }
                case protoreflect.Int32Kind, protoreflect.Int64Kind,
                    protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
                    protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
                    if v := field.GetInt64Option(option); v != nil {
                        if !first {
                            buffer.WriteString(", ")
                        }
                        buffer.WriteString(fmt.Sprint("(", mojo.GetOptionFullName(option), ")=", *v))
                        first = false
                    }
                case protoreflect.StringKind:
                    if v := field.GetStringOption(option); len(v) > 0 {
                        if !first {
                            buffer.WriteString(", ")
                        }
                        buffer.WriteString(fmt.Sprint("(", mojo.GetOptionFullName(option), ")=", "\"", v, "\""))
                        first = false
                    }
                }
            }
            if buffer.Len() > 0 {
                g.S(" [")
                g.S(buffer.String())
                g.S("];\n")
            } else {
                g.S(";\n")
            }
        } else {
            g.S(";\n")
        }

        //fieldDeprecated := ""
        //if field.GetOptions().IsDeprecated() {
        //	fieldDeprecated = deprecationComment
        //}

        //fieldFullPath := fmt.Sprintf("%s,%d,%d", message.path, messageFieldPath, i)
        //c, ok := g.makeComments(fieldFullPath)
        //if ok {
        //	c += "\n"
        //}
    }

    printedOneofs := make(map[string]bool)
    for _, field := range message.Fields {
        if field.Oneof == nil {
            printField(field)
        } else {
            if _, ok := printedOneofs[field.Oneof.GetName()]; !ok {
                g.P("oneof ", field.Oneof.GetName(), " {")
                g.In()
                fields := field.Oneof.Fields
                for _, f := range fields {
                    printField(f)
                }
                g.Out()
                g.P("}")

                printedOneofs[field.Oneof.GetName()] = true
            }
        }
    }

    g.Out()
    g.P("}")
}

// Generate the type, methods and default constant definitions for this Descriptor.
func (g *Generator) generateService(service *descriptor.Service) {
    g.P("service ", service.GetName(), " {")
    g.In()

    for _, method := range service.Methods {
        g.P("rpc ", method.GetName(), "(", method.GetInput().GetName(), ") returns (", method.GetOutpu().GetName(), ");")
    }

    g.Out()
    g.P("}")
}
