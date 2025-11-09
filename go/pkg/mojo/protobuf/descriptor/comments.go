package descriptor

import (
    "google.golang.org/protobuf/reflect/protoreflect"
    "strings"
)

// Comments is a comments string as provided by protoc.
type Comments string

// String formats the comments by inserting // to the start of each line,
// ensuring that there is a trailing newline.
// An empty comment is formatted as an empty string.
func (c Comments) String() string {
    if c == "" {
        return ""
    }
    var b []byte
    for _, line := range strings.Split(strings.TrimSuffix(string(c), "\n"), "\n") {
        b = append(b, "//"...)
        b = append(b, line...)
        b = append(b, "\n"...)
    }
    return string(b)
}

// CommentSet is a set of leading and trailing comments associated
// with a .proto descriptor declaration.
type CommentSet struct {
    LeadingDetached []Comments
    Leading         Comments
    Trailing        Comments
}

func makeCommentSet(loc protoreflect.SourceLocation) CommentSet {
    var leadingDetached []Comments
    for _, s := range loc.LeadingDetachedComments {
        leadingDetached = append(leadingDetached, Comments(s))
    }
    return CommentSet{
        LeadingDetached: leadingDetached,
        Leading:         Comments(loc.LeadingComments),
        Trailing:        Comments(loc.TrailingComments),
    }
}
