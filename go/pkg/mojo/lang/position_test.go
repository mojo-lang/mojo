package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPositionOrdering(t *testing.T) {
	positions := []*Position{nil, {Line: 1, Column: 1}, {Line: 1, Column: 2}, {Line: 2, Column: 1}}
	for i, left := range positions {
		for j, right := range positions {
			want := 0
			if i < j {
				want = -1
			} else if i > j {
				want = 1
			}
			require.Equal(t, want, left.Compare(right))
		}
	}
	require.Zero(t, (&Position{Line: 1, Column: 1, Offset: 100}).Compare(&Position{Line: 1, Column: 1}))
}

func TestPatchPositionKeepsAndMergesComments(t *testing.T) {
	leading := NewBlockCommentComment(&BlockComment{})
	trailing := NewBlockCommentComment(&BlockComment{})
	target := &Position{Filename: "old", Line: 1, Column: 2, LeadingComments: []*Comment{leading}}
	patch := &Position{Filename: "new", Line: 3, Column: 4, Offset: 5, TailingComments: []*Comment{trailing}}
	require.Same(t, target, PatchPosition(target, patch))
	require.Equal(t, "new", target.Filename)
	require.Equal(t, int64(3), int64(target.Line))
	require.Equal(t, []*Comment{leading}, target.LeadingComments)
	require.Equal(t, []*Comment{trailing}, target.TailingComments)
	PatchPosition(target, &Position{LeadingComments: []*Comment{trailing}})
	require.Equal(t, "new", target.Filename)
	require.Len(t, target.LeadingComments, 2)
	require.Same(t, target, PatchPosition(target, nil))
	require.Same(t, patch, PatchPosition(nil, patch))
	require.Nil(t, PatchPosition(nil, nil))
}

func TestSortAndAppendPositionComments(t *testing.T) {
	late := NewBlockCommentComment(&BlockComment{StartPosition: &Position{Line: 3}})
	early := NewBlockCommentComment(&BlockComment{StartPosition: &Position{Line: 1}})
	pos := &Position{}
	pos.AppendLeadingComment(late)
	pos.AppendLeadingComment(early)
	require.Equal(t, []*Comment{early, late}, pos.LeadingComments)
	pos.AppendTailingComment(late)
	pos.AppendTailingComment(early)
	require.Equal(t, []*Comment{early, late}, pos.TailingComments)
	pos.AppendLeadingComment("unsupported")
	require.Len(t, pos.LeadingComments, 2)
}
