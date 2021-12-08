package descriptor

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"strconv"
	"strings"
)

// FileDescriptor describes an protocol buffer descriptor File (.proto).
// It includes slices of all the Messages and Enums defined within it.
// Those slices are constructed by WrapTypes.
type FileDescriptor struct {
	*descriptor.FileDescriptorProto

	Messages []*MessageDescriptor // All the Messages defined in this File.
	Enums    []*EnumDescriptor    // All the Enums defined in this File.
	Services []*ServiceDescriptor // All the top-level extensions defined in this File.

	// Comments, stored as a map of Path (comma-separated integers) to the comment.
	Comments map[string]*descriptor.SourceCodeInfo_Location

	Proto3 bool // whether to generate proto3 code for this File
}

func NewFileDescriptor() *FileDescriptor {
	return &FileDescriptor{
		FileDescriptorProto: &descriptor.FileDescriptorProto{},

		Messages: make([]*MessageDescriptor, 0),
	}
}

func (f *FileDescriptor) HasOptions() bool {
	if f == nil || f.Options == nil {
		return false
	}

	options := f.Options
	return options.GoPackage != nil ||
		options.JavaPackage != nil ||
		options.JavaMultipleFiles != nil ||
		options.Deprecated != nil ||
		options.JavaOuterClassname != nil
}

func (f *FileDescriptor) HasService() bool {
	return f != nil && len(f.Services) > 0
}

func (f *FileDescriptor) IsMessageExist(name string) bool {
	if f == nil {
		return false
	}
	for _, msg := range f.Messages {
		if name == *msg.Name {
			return true
		}
	}
	return false
}

func NewServiceDescriptor(file *FileDescriptor) *ServiceDescriptor {
	return &ServiceDescriptor{
		Common: Common{
			File: file,
		},
		ServiceDescriptorProto: &descriptor.ServiceDescriptorProto{},
	}
}

func FileIsProto3(file *descriptor.FileDescriptorProto) bool {
	return file.GetSyntax() == "Proto3"
}

func ExtractComments(file *FileDescriptor) {
	file.Comments = make(map[string]*descriptor.SourceCodeInfo_Location)
	for _, loc := range file.GetSourceCodeInfo().GetLocation() {
		if loc.LeadingComments == nil {
			continue
		}
		var p []string
		for _, n := range loc.Path {
			p = append(p, strconv.Itoa(int(n)))
		}
		file.Comments[strings.Join(p, ",")] = loc
	}
}

// Return a slice of all the Descriptors defined within this File
func WrapMessageDescriptors(file *FileDescriptor) []*MessageDescriptor {
	sl := make([]*MessageDescriptor, 0, len(file.MessageType)+10)
	for i, desc := range file.MessageType {
		sl = wrapThisMessageDescriptor(sl, desc, nil, file, i)
	}
	return sl
}

func HasJavaOptions(options *descriptor.FileOptions) bool {
	return options != nil && (options.JavaPackage != nil ||
		options.JavaMultipleFiles != nil ||
		options.JavaOuterClassname != nil ||
		options.JavaGenericServices != nil ||
		options.JavaStringCheckUtf8 != nil)
}

func UnmarshalFiles(bytes []byte) ([]*descriptor.FileDescriptorProto, error) {
	fileDesc := &descriptor.FileDescriptorSet{}
	if err := proto.Unmarshal(bytes, fileDesc); err != nil {
		return nil, err
	}
	return fileDesc.File, nil
}
