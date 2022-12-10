package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type CommentMerger []*lang.Comment

func (c CommentMerger) Merge(file *lang.SourceFile) {
	if file == nil || len(c) == 0 {
		return
	}

	index := 0
	statements := file.Statements
	for i := 0; i < len(statements); {
		statement := statements[i]
		for index < len(c) {
			relative := c[index].ComparePosition(statement.GetStartPosition(), statement.GetEndPosition())
			if relative == lang.RelativeLeading {
				statement.GetStartPosition().AppendLeadingComment(c[index])
				index++
			} else if relative == lang.RelativeInternal {
				_, _ = statement.MergeComment(c[index])
				index++
			} else {
				if c[index].IsFollowing() && c[index].GetStartPosition().GetLine() == statement.GetEndPosition().GetLine() {
					statement.GetEndPosition().AppendTailingComment(c[index])
					index++
				}
				i++
				break
			}
		}
		if index == len(c) {
			break
		}
	}

	file.TailingComments = append(file.TailingComments, c[index:]...)
	file.TailingComments = lang.SortComments(file.TailingComments)
}
