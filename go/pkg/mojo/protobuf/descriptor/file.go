package descriptor

import (
	"sort"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	Proto2Syntax = "proto2"
	Proto3Syntax = "proto3"
)

// File describes a protocol buffer descriptor File (.proto).
// It includes slices of all the Messages and Enums defined within it.
// Those slices are constructed by WrapTypes.
type File struct {
	Packages *Packages

	Proto *descriptorpb.FileDescriptorProto

	Messages []*Message // All the Messages defined in this File.
	Enums    []*Enum    // All the Enums defined in this File.
	Services []*Service // All the top-level extensions defined in this File.

	cursor protoreflect.SourcePath
}

func NewFile() *File {
	return &File{
		Proto: &descriptorpb.FileDescriptorProto{},
	}
}

func NewFileWithName(name string, pkg string) *File {
	return &File{
		Proto: &descriptorpb.FileDescriptorProto{
			Name:    &name,
			Package: &pkg,
			Syntax:  &Proto3Syntax,
		},
	}
}

func NewFileFrom(proto *descriptorpb.FileDescriptorProto) *File {
	file := &File{
		Proto: proto,
	}

	for _, msg := range proto.MessageType {
		m := NewMessageFrom(file, msg)
		file.Messages = append(file.Messages, m)
	}

	for _, enum := range proto.EnumType {
		e := NewEnumFrom(file, enum)
		file.Enums = append(file.Enums, e)
	}

	for _, service := range proto.Service {
		s := NewServiceFrom(file, service)
		file.Services = append(file.Services, s)
	}

	return file
}

func (f *File) proto() *descriptorpb.FileDescriptorProto {
	if f != nil {
		return f.Proto
	}
	return nil
}

func (f *File) GetPackages() *Packages {
	if f != nil {
		return f.Packages
	}
	return nil
}

func (f *File) IsDeprecated() bool {
	return f.proto().GetOptions().GetDeprecated()
}

func (f *File) GetOptions() *descriptorpb.FileOptions {
	if f != nil && f.Proto != nil {
		if f.Proto.Options == nil {
			f.Proto.Options = &descriptorpb.FileOptions{}
		}
		return f.Proto.Options
	}
	return nil
}

func (f *File) HasOptions() bool {
	if options := f.proto().GetOptions(); options != nil {
		if options.Deprecated != nil ||
			options.OptimizeFor != nil {
			return true
		}
	}
	if f.HasJavaOptions() || f.HasGoOptions() || f.HasCcOptions() || f.HasPyOptions() {
		return true
	}
	return false
}

func (f *File) HasGoOptions() bool {
	if options := f.proto().GetOptions(); options != nil {
		return options.GoPackage != nil
	}
	return false
}

func (f *File) HasCcOptions() bool {
	if options := f.proto().GetOptions(); options != nil {
		return options.CcEnableArenas != nil || options.CcGenericServices != nil
	}
	return false
}

func (f *File) HasPyOptions() bool {
	if options := f.proto().GetOptions(); options != nil {
		return options.PyGenericServices != nil
	}
	return false
}

func (f *File) HasJavaOptions() bool {
	if options := f.proto().GetOptions(); options != nil {
		return options.JavaPackage != nil ||
			options.JavaMultipleFiles != nil ||
			options.JavaOuterClassname != nil ||
			options.JavaGenericServices != nil ||
			options.JavaStringCheckUtf8 != nil
	}
	return false
}

func (f *File) IsProto3() bool {
	return f != nil && (len(f.proto().GetSyntax()) == 0 || f.proto().GetSyntax() == Proto3Syntax)
}

func (f *File) SetProto3(value bool) *File {
	if f != nil && f.Proto != nil {
		if value {
			f.Proto.Syntax = &Proto3Syntax
		} else {
			f.Proto.Syntax = &Proto2Syntax
		}
	}
	return f
}

func (f *File) GetName() string {
	return f.proto().GetName()
}

func (f *File) SetName(name string) *File {
	if f != nil && f.Proto != nil {
		f.Proto.Name = &name
	}
	return f
}

func (f *File) GetPackageName() string {
	return f.proto().GetPackage()
}

func (f *File) SetPackageName(name string) *File {
	if f != nil && f.Proto != nil {
		f.Proto.Package = &name
	}
	return f
}

func (f *File) HasMessage() bool {
	return f != nil && len(f.Messages) > 0
}

func (f *File) GetMessage(name string) *Message {
	if f != nil {
		for _, msg := range f.Messages {
			if name == msg.GetName() {
				return msg
			}
		}
	}
	return nil
}

func (f *File) IsMessageExist(name string) bool {
	return f.GetMessage(name) != nil
}

func (f *File) HasEnum() bool {
	return f != nil && len(f.Enums) > 0
}

func (f *File) GetEnum(name string) *Enum {
	if f != nil {
		for _, e := range f.Enums {
			if name == e.GetName() {
				return e
			}
		}
	}
	return nil
}

func (f *File) IsEnumExist(name string) bool {
	return f.GetEnum(name) != nil
}

func (f *File) HasService() bool {
	return f != nil && len(f.Services) > 0
}

func (f *File) GetService(name string) *Service {
	if f != nil {
		for _, s := range f.Services {
			if name == s.GetName() {
				return s
			}
		}
	}
	return nil
}

func (f *File) IsServiceExist(name string) bool {
	return f.GetService(name) != nil
}

func (f *File) IsEmpty() bool {
	return !(f.HasMessage() || f.HasEnum() || f.HasService())
}

func (f *File) AppendMessage(message *Message) *File {
	if f != nil && f.Proto != nil {
		f.Messages = append(f.Messages, message)
		f.proto().MessageType = append(f.proto().MessageType, message.Proto)
	}
	return f
}

func (f *File) AppendEnum(enum *Enum) *File {
	if f != nil && f.Proto != nil {
		f.Enums = append(f.Enums, enum)
		f.proto().EnumType = append(f.proto().EnumType, enum.Proto)
	}
	return f
}

func (f *File) AppendService(service *Service) *File {
	if f != nil && f.Proto != nil {
		f.Services = append(f.Services, service)
		f.proto().Service = append(f.proto().Service, service.Proto)
	}
	return f
}

func (f *File) GetDependencies() []string {
	return f.proto().GetDependency()
}

func (f *File) AppendDependency(dependency string) *File {
	if f != nil && f.Proto != nil {
		f.Proto.Dependency = append(f.Proto.Dependency, dependency)
	}
	return f
}

func (f *File) CleanDependency() *File {
	if f != nil && f.Proto != nil && len(f.Proto.Dependency) > 0 {
		vals := core.NewStringValues(f.Proto.Dependency...).Unique().Vals
		f.Proto.Dependency = []string{}
		for _, val := range vals {
			if val != f.GetName() {
				f.Proto.Dependency = append(f.Proto.Dependency, val)
			}
		}
		sort.Strings(f.Proto.Dependency)
	}
	return f
}

func (f *File) ExtractComments() {
	// file.Comments = make(map[string]*descriptorpb.SourceCodeInfo_Location)
	// for _, loc := range file.GetSourceCodeInfo().GetLocation() {
	//    if loc.LeadingComments == nil {
	//        continue
	//    }
	//    var p []string
	//    for _, n := range loc.Path {
	//        p = append(p, strconv.Itoa(int(n)))
	//    }
	//    file.Comments[strings.Join(p, ",")] = loc
	// }
}

// func UnmarshalFiles(bytes []byte) ([]*descriptorpb.FileProto, error) {
//    fileDesc := &descriptorpb.FileDescriptorProto{}
//    if err := proto.Unmarshal(bytes, fileDesc); err != nil {
//        return nil, err
//    }
//    return fileDesc.File, nil
// }
