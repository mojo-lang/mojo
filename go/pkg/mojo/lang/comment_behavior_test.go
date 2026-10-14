package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommentVariantsAndRelativePositions(t *testing.T) {
	for _, node := range []interface{}{&Document{}, &BlockComment{}, &MultiLineComment{}} {
		comment := NewComment(node)
		require.NotNil(t, comment)
		comment.SetStartPosition(&Position{Line: 3, Column: 2})
		comment.SetEndPosition(&Position{Line: 3, Column: 8})
		require.Equal(t, int64(2), comment.GetStartPosition().Column)
		require.Equal(t, int64(8), comment.GetEndPosition().Column)
		require.Equal(t, RelativeType(RelativeLeading), comment.ComparePosition(&Position{Line: 4}, &Position{Line: 5}))
		require.Equal(t, RelativeType(RelativeTailing), comment.ComparePosition(&Position{Line: 1}, &Position{Line: 2}))
		require.Equal(t, RelativeType(RelativeInternal), comment.ComparePosition(&Position{Line: 3}, &Position{Line: 4}))
	}
	require.Nil(t, NewComment("unsupported"))
	require.Len(t, NewComments(&Document{}, &BlockComment{}, &MultiLineComment{}, nil, "unsupported"), 3)
	block := &BlockComment{StartPosition: &Position{Line: 1}, EndPosition: &Position{Line: 1}}
	require.True(t, block.IsInLine())
	block.EndPosition.Line = 2
	require.False(t, block.IsInLine())
	var absent *Comment
	require.Nil(t, absent.GetStartPosition())
	require.Nil(t, absent.GetEndPosition())
	require.False(t, absent.IsFollowing())
	require.NotPanics(t, func() { absent.SetStartPosition(&Position{}); absent.SetEndPosition(&Position{}) })
}

func TestMergeCommentsAttachesToTheCorrectTerm(t *testing.T) {
	term := NewSymbolTerm(&Position{Line: 3, Column: 1}, TermTypeName, "field")
	leading := NewBlockCommentComment(&BlockComment{StartPosition: &Position{Line: 1}, EndPosition: &Position{Line: 2}})
	merged, err := MergeCommentToTerms(leading, term)
	require.NoError(t, err)
	require.True(t, merged)
	require.Equal(t, []*Comment{leading}, term.StartPosition.LeadingComments)
	following := NewMultiLineCommentComment(&MultiLineComment{StartPosition: &Position{Line: 3, Column: 10}, EndPosition: &Position{Line: 3, Column: 20}, Lines: []*LineComment{{Following: true}}})
	merged, err = MergeCommentToTerms(following, term)
	require.NoError(t, err)
	require.True(t, merged)
	require.Equal(t, []*Comment{following}, term.EndPosition.TailingComments)
	internal := NewBlockCommentComment(&BlockComment{StartPosition: &Position{Line: 3, Column: 2}, EndPosition: &Position{Line: 3, Column: 3}})
	merged, err = MergeCommentToTerms(internal, term)
	require.Error(t, err)
	require.False(t, merged)
	merged, err = MergeCommentToTerms(leading, "unsupported")
	require.Error(t, err)
	require.False(t, merged)
	later := NewSymbolTerm(&Position{Line: 4, Column: 1}, TermTypeName, "next")
	following.GetMultiLineComment().Lines[0].Following = false
	merged, err = MergeCommentToTerms(following, term, later)
	require.NoError(t, err)
	require.True(t, merged)
	require.Equal(t, []*Comment{following}, later.StartPosition.LeadingComments)
	_, err = MergeCommentToTerms(following)
	require.Error(t, err)
}

func TestDocumentContentAndMarkdownParsing(t *testing.T) {
	document := &Document{Lines: []*Document_Line{{Content: ""}, {Content: ""}, {Content: "# Title"}, {Content: ""}, {Content: "Body"}}}
	require.Equal(t, "# Title\n\nBody", document.GetContent())
	require.Same(t, document, document.Parse())
	require.NotNil(t, document.Structured)
	var absent *Document
	require.Empty(t, absent.GetContent())
	require.Nil(t, absent.Parse())
	require.Empty(t, (&Document{Lines: []*Document_Line{{}, {}}}).GetContent())
}
