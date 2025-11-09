package query

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"regexp"
	"strings"
)

var fieldPattern = regexp.MustCompile(`[a-zA-Z][0-9a-zA-Z\-._]*`)

func normalizeFiledMask(fieldMask *core.FieldMask) []string {
	var fields []string
	for _, p := range fieldMask.Paths {
		if len(p) > 0 && fieldPattern.MatchString(p) {
			f := strings.ReplaceAll(p, ".", "_")
			fields = append(fields, f)
		}
	}
	return fields
}

//func ApplyFieldMask(d *gorm.DB, fieldMask *core.FieldMask) (*gorm.DB, error) {
//	if d != nil && fieldMask != nil && len(fieldMask.Paths) > 0 {
//		fields := normalizeFiledMask(fieldMask)
//		if len(fields) > 0 {
//			return d.Select(fields), nil
//		}
//	}
//	return d, nil
//}
