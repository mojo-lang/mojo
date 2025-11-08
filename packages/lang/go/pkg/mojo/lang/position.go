package lang

import (
	"reflect"
	"sort"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

type PositionGetter interface {
	StartPositionGetter
	EndPositionGetter
}

type PositionSetter interface {
	StartPositionSetter
	EndPositionSetter
}

const (
	StartPositionFieldName = "StartPosition"
	EndPositionFieldName   = "EndPosition"
)

type PositionGetters []PositionGetter

func (p PositionGetters) Len() int {
	return len(p)
}
func (p PositionGetters) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PositionGetters) Less(i, j int) bool {
	return p[i].GetStartPosition().Compare(p[j].GetStartPosition()) < 0
}

type StartPositionGetter interface {
	GetStartPosition() *Position
}

type StartPositionSetter interface {
	SetStartPosition(position *Position)
}

func SetStartPosition(object interface{}, position *Position) {
	if setter, ok := object.(StartPositionSetter); ok {
		setter.SetStartPosition(position)
	}
}

type EndPositionGetter interface {
	GetEndPosition() *Position
}

type EndPositionSetter interface {
	SetEndPosition(position *Position)
}

func SetEndPosition(object interface{}, position *Position) {
	if setter, ok := object.(EndPositionSetter); ok {
		setter.SetEndPosition(position)
	}
}

// PatchPosition patch the set value to target
func PatchPosition(target *Position, patch *Position) *Position {
	if target == nil {
		return patch
	}
	if patch == nil {
		return target
	}

	// patch position has set
	if len(patch.Filename) > 0 || patch.Line > 0 || patch.Column > 0 || patch.Offset > 0 {
		target.Filename = patch.Filename
		target.Line = patch.Line
		target.Column = patch.Column
		target.Offset = patch.Offset
	}

	if len(patch.LeadingComments) > 0 {
		target.LeadingComments = append(target.LeadingComments, patch.LeadingComments...)
	}
	if len(patch.TailingComments) > 0 {
		target.TailingComments = append(target.TailingComments, patch.TailingComments...)
	}

	return target
}

func GetUnionPosition(union interface{}, name string) *Position {
	if union != nil {
		if pos, ok := core.GetUnionField(union, name).Interface().(*Position); ok {
			return pos
		}
	}
	return nil
}

func SetUnionPosition(union interface{}, name string, position *Position) {
	if union != nil {
		value := reflect.ValueOf(union)
		for {
			if _, ok := value.Interface().(core.IsUnion); ok {
				value = reflect.Indirect(value).Field(core.GetUnionFieldIndex(value)).Elem()
				value = reflect.Indirect(value).Field(0)
			} else {
				break
			}
		}

		value = reflect.Indirect(value).FieldByName(name)
		if value.CanSet() {
			if value.IsNil() {
				value.Set(reflect.ValueOf(position))
			} else {
				if pos, ok := value.Interface().(*Position); ok {
					value.Set(reflect.ValueOf(PatchPosition(pos, position)))
				}
			}
		}
	}
}

func (x *Position) Compare(position *Position) int {
	if x == nil {
		if position == nil {
			return 0
		} else {
			return -1
		}
	} else {
		if position == nil {
			return 1
		} else {
			if x.Line == position.Line {
				if x.Column == position.Column {
					return 0
				} else if x.Column < position.Column {
					return -1
				} else {
					return 1
				}
			} else if x.Line < position.Line {
				return -1
			} else {
				return 1
			}
		}
	}
}

func (x *Position) AppendLeadingComment(comment interface{}) {
	if x != nil {
		switch c := comment.(type) {
		case *Comment:
			x.LeadingComments = append(x.LeadingComments, c)
		case *BlockComment:
			x.LeadingComments = append(x.LeadingComments, NewBlockCommentComment(c))
		case *MultiLineComment:
			x.LeadingComments = append(x.LeadingComments, NewMultiLineCommentComment(c))
		case *Document:
			x.LeadingComments = append(x.LeadingComments, NewDocumentComment(c))
		}
		x.LeadingComments = SortComments(x.LeadingComments)
	}
}

func (x *Position) AppendTailingComment(comment interface{}) {
	if x != nil {
		switch c := comment.(type) {
		case *Comment:
			x.TailingComments = append(x.TailingComments, c)
		case *BlockComment:
			x.TailingComments = append(x.TailingComments, NewBlockCommentComment(c))
		case *MultiLineComment:
			x.TailingComments = append(x.TailingComments, NewMultiLineCommentComment(c))
		case *Document:
			x.TailingComments = append(x.TailingComments, NewDocumentComment(c))
		}
		x.TailingComments = SortComments(x.TailingComments)
	}
}

func SortComments(comments []*Comment) []*Comment {
	var gs PositionGetters
	for _, comment := range comments {
		gs = append(gs, comment)
	}
	sort.Sort(gs)

	var sortedComments []*Comment
	for _, g := range gs {
		sortedComments = append(sortedComments, g.(*Comment))
	}
	return sortedComments
}
