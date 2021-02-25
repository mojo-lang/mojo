package descriptor

import "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// EnumDescriptor describes an enum. If it's at top level, its Parent will be nil.
// Otherwise it will be the descriptor of the message in which it is defined.
type EnumDescriptor struct {
	Common
	*descriptor.EnumDescriptorProto

	Parent *MessageDescriptor // The containing message, if any.
	Index  int                // The Index into the container, whether the File or a message.
	Path   string             // The SourceCodeInfo Path as comma-separated integers.
}

func NewEnumDescriptor(file *FileDescriptor) *EnumDescriptor {
	return &EnumDescriptor{
		Common: Common{
			File: file,
		},
		EnumDescriptorProto: &descriptor.EnumDescriptorProto{},
	}
}

// Return a slice of all the EnumDescriptors defined within this File
func WrapEnumDescriptors(file *FileDescriptor, descs []*MessageDescriptor) []*EnumDescriptor {
	sl := make([]*EnumDescriptor, 0, len(file.EnumType)+10)
	// Top-level Enums.
	for i, enum := range file.EnumType {
		sl = append(sl, newEnumDescriptor(enum, nil, file, i))
	}
	// Enums within Messages. Enums within embedded Messages appear in the outer-most message.
	for _, nested := range descs {
		for i, enum := range nested.EnumType {
			sl = append(sl, newEnumDescriptor(enum, nested, file, i))
		}
	}
	return sl
}

// Construct the EnumDescriptor
func newEnumDescriptor(desc *descriptor.EnumDescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) *EnumDescriptor {
	ed := &EnumDescriptor{
		Common:              Common{file, nil},
		EnumDescriptorProto: desc,
		Parent:              parent,
		Index:               index,
	}
	if parent == nil {
		//ed.Path = fmt.Sprintf("%d,%d", enumPath, Index)
	} else {
		//ed.Path = fmt.Sprintf("%s,%d,%d", Parent.Path, messageEnumPath, Index)
	}
	return ed
}
