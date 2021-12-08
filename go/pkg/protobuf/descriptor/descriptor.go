package descriptor

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Each type we import as a protocol buffer (other than FileDescriptorProto) needs
// a pointer to the FileDescriptorProto that represents it.  These types achieve that
// wrapping by placing each Proto inside a struct with the pointer to its File. The
// structs have the same names as their contents, with "Proto" removed.
// FileDescriptor is used to store the things that it points to.

// The File and package name method are Common to Messages and Enums.
type Common struct {
	File     *FileDescriptor // File this object comes from.
	Comments *descriptor.SourceCodeInfo_Location
}

func (c *Common) LeadingComments() *string {
	if c.Comments != nil {
		return c.Comments.LeadingComments
	}
	return nil
}

func (c *Common) TrailingComments() *string {
	if c.Comments != nil {
		return c.Comments.TrailingComments
	}
	return nil
}
