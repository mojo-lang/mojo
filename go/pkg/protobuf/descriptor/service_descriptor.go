package descriptor

import "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// ServiceDescriptor describes an service.
type ServiceDescriptor struct {
	Common
	*descriptor.ServiceDescriptorProto

	Index int // The Index into the container, whether the File or a message.
}

// Return a slice of all the EnumDescriptors defined within this File
func WrapServiceDescriptors(file *FileDescriptor) []*ServiceDescriptor {
	sl := make([]*ServiceDescriptor, 0, len(file.Service)+10)
	// Top-level Enums.
	for i, service := range file.Service {
		sl = append(sl, newServiceDescriptor(service, nil, file, i))
	}
	return sl
}

// Construct the ServiceDescriptor
func newServiceDescriptor(desc *descriptor.ServiceDescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) *ServiceDescriptor {
	sd := &ServiceDescriptor{
		Common:                 Common{file, nil},
		ServiceDescriptorProto: desc,
		Index:                  index,
	}
	if parent == nil {
		//ed.Path = fmt.Sprintf("%d,%d", enumPath, Index)
	} else {
		//ed.Path = fmt.Sprintf("%s,%d,%d", Parent.Path, messageEnumPath, Index)
	}
	return sd
}
