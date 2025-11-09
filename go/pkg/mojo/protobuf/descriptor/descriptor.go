package descriptor

import "google.golang.org/protobuf/reflect/protoreflect"

// Each type we import as a protocol buffer (other than FileDescriptorProto) needs
// a pointer to the FileDescriptorProto that represents it.  These types achieve that
// wrapping by placing each Proto inside a struct with the pointer to its File. The
// structs have the same names as their contents, with "Proto" removed.
// FileDescriptor is used to store the things that it points to.

// Descriptor The File and package name method are Descriptor to Messages and Enums.
type Descriptor struct {
    File *File // File this object comes from.

    Path     protoreflect.SourcePath
    Comments CommentSet // comments associated with this descriptor
}

func (d *Descriptor) LeadingComments() Comments {
    if d != nil {
        return d.Comments.Leading
    }
    return ""
}

func (d *Descriptor) TrailingComments() Comments {
    if d != nil {
        return d.Comments.Trailing
    }
    return ""
}

// AppendPath add elements to a Location's path, returning a new Location.
func (d *Descriptor) AppendPath(num protoreflect.FieldNumber, idx int) *Descriptor {
    d.Path = append(protoreflect.SourcePath(nil), d.Path...) // make copy
    d.Path = append(d.Path, int32(num), int32(idx))
    return d
}
