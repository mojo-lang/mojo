package query

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
)

// Query
// field mask + unique  => Uniques
// Uniques will shadow the field mask field
type Query struct {
	FieldQueries []*FieldQuery    `json:"fieldQueries,omitempty"`
	Filter       *lang.Expression `json:"filter,omitempty"`
	Order        *core.Ordering   `json:"order,omitempty"`
	Unique       bool             `json:"unique,omitempty"`
	Uniques      []string         `json:"uniques,omitempty"`
	FieldMask    *core.FieldMask  `json:"fieldMask,omitempty"`

	Projections []*FieldProjection       `json:"projections,omitempty"`
	Groups      []string                 `json:"groups,omitempty"`
	Fields      map[string]*ProjectField `json:"-"`

	PageSize  int32  `json:"pageSize,omitempty"`
	PageToken string `json:"pageToken,omitempty"`
	Skip      int32  `json:"skip,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

func (q *Query) GetField(name string) *ProjectField {
	if q != nil {
		if field, ok := q.Fields[name]; ok {
			return field
		}
	}
	return nil
}

func (q *Query) Normalize() error {
	if q != nil {
		if q.FieldMask != nil {
			fields := normalizeFiledMask(q.FieldMask)

			if q.Unique {
				if len(q.Uniques) > 0 {
					fieldIndex := make(map[string]bool)
					for _, f := range fields {
						fieldIndex[f] = true
					}
					for _, u := range q.Uniques {
						if ok := fieldIndex[u]; !ok {
							return fmt.Errorf("uniques (%v) should be same with fieldmask (%v) or empty when unique is true", q.Uniques, fields)
						}
					}
					q.Uniques = []string{}
				}

				for _, p := range fields {
					q.Uniques = append(q.Uniques, p)
				}
			} else {
				if len(q.Uniques) > 0 {
					q.FieldMask = nil
				}
			}

			if len(q.Uniques) > 0 {
				if len(q.Projections) > 0 {
					return fmt.Errorf("uniques (%v) should be exclusive to field projections", q.Uniques)
				}

				q.Projections = append(q.Projections, &FieldProjection{
					Name:      "",
					Uniques:   q.Uniques,
					Functions: nil,
					Alias:     nil,
				})
			} else {
				for _, field := range fields {
					q.Projections = append(q.Projections, &FieldProjection{
						Name:      field,
						Functions: nil,
						Alias:     nil,
					})
				}
			}
		} else {
			if len(q.Uniques) > 0 {
				if len(q.Projections) > 0 {
					return fmt.Errorf("uniques (%v) should be exclusive to field projections", q.Uniques)
				}

				if q.Unique {
					q.Unique = false
				}

				q.Projections = append(q.Projections, &FieldProjection{
					Name:      "",
					Uniques:   q.Uniques,
					Functions: nil,
					Alias:     nil,
				})
			}
		}
	}

	return nil
}

func (q *Query) TotalCount(d *gorm.DB, meta Fields) *gorm.DB {
	if q == nil {
		return d
	}

	projections := q.Projections
	ps, pt := q.PageSize, q.PageToken

	if len(q.Uniques) > 0 {
		q.Projections = []*FieldProjection{{Uniques: q.Uniques, Functions: []string{"COUNT"}, Alias: []string{"count"}}}
	} else {
		q.Projections = []*FieldProjection{{Name: "*", Functions: []string{"COUNT"}, Alias: []string{"count"}}}
	}
	q.PageSize = 0
	q.PageToken = ""
	tx := q.Apply(d, meta)

	q.PageSize = ps
	q.PageToken = pt
	q.Projections = projections
	return tx
}

func (q *Query) Apply(d *gorm.DB, meta Fields) *gorm.DB {
	if q == nil {
		return d
	}

	var err error
	tx := d

	if len(q.Projections) > 0 {
		if tx, err = q.applyProjections(tx, meta); err != nil {
			_ = tx.AddError(err)
			return tx
		}
	}

	if len(q.Groups) > 0 {
		tx = tx.Group(strings.Join(q.Groups, ","))
	}

	if tx, err = q.applyQuery(tx, meta); err != nil {
		_ = tx.AddError(err)
		return tx
	}

	return tx
}

func (q *Query) applyProjections(d *gorm.DB, meta Fields) (tx *gorm.DB, err error) {
	builder := &strings.Builder{}

	tx = d
	if len(q.Projections) == 1 && len(q.Projections[0].Functions) == 0 {
		projection := q.Projections[0]
		q.Fields = make(map[string]*ProjectField)
		if len(projection.Name) > 0 {
			q.Fields[projection.Name] = &ProjectField{
				Projection: projection,
				Index:      0,
			}

			return d.Select(projection.Name), nil
		} else if len(projection.Uniques) > 0 {
			for i, unique := range projection.Uniques {
				q.Fields[unique] = &ProjectField{
					Projection: projection,
					Index:      i,
				}
			}
			return tx.Distinct(q.Projections[0].Uniques), nil
		} else {
			return nil, errors.New("invalid projection, has no name and uniques")
		}
	}

	for _, group := range q.Groups {
		found := false
		for _, prj := range q.Projections {
			if prj.Name == group {
				found = true
				break
			}
		}
		if !found {
			q.Projections = append(q.Projections, &FieldProjection{
				Name: group,
			})
		}
	}

	projections := make(map[string]*ProjectField)
	for i, projection := range q.Projections {
		if i > 0 {
			builder.WriteString(", ")
		}
		if len(projection.Functions) == 0 {
			builder.WriteString(projection.Name)
			projections[projection.Name] = &ProjectField{
				Projection: projection,
				Index:      0,
			}
			continue
		}

		for j, fun := range projection.Functions {
			if j > 0 {
				builder.WriteString(", ")
			}
			if len(projection.Alias) < len(projection.Functions) {
				projection.Alias = append(projection.Alias, "")
			}

			if len(projection.Name) == 0 && len(projection.Uniques) == 0 && strings.ToLower(fun) == "count" {
				projection.Name = "*"
			}

			if projection.Name == "*" {
				if len(projection.Alias[j]) == 0 {
					alias := fun

					projection.Alias[j] = alias
					projections[alias] = &ProjectField{
						Projection: projection,
						Index:      j,
					}
				}
			} else {
				alias := projection.Alias[j]
				if len(alias) == 0 {
					if len(projection.Name) == 0 {
						if len(projection.Uniques) > 0 {
							alias = projection.Functions[j]
						} else {
							err = errors.New("name and uniques should not be empty both in the projection")
							return
						}
					} else {
						alias = projection.Name + "_" + strings.ToLower(fun)
					}

					projection.Alias[j] = alias
				}

				projections[alias] = &ProjectField{
					Projection: projection,
					Index:      j,
				}
			}

			if len(projection.Name) > 0 {
				builder.WriteString(fmt.Sprintf("%s(%s) as %s", fun, projection.Name, projection.Alias[j]))
			} else {
				unique := fmt.Sprintf("DISTINCT(%s)", strings.Join(projection.Uniques, ","))
				builder.WriteString(fmt.Sprintf("%s(%s) as %s", fun, unique, projection.Alias[j]))
			}
		}
	}
	q.Fields = projections

	return d.Select(builder.String()), nil
}

func (q *Query) applyQuery(d *gorm.DB, meta Fields) (tx *gorm.DB, err error) {
	tx = d

	if len(q.FieldQueries) > 0 {
		for _, query := range q.FieldQueries {
			if tx, err = ApplyFieldQuery(tx, query, meta); err != nil {
				return
			}
		}
	}
	if q.Filter != nil {
		if tx, err = ApplyFilter(tx, q.Filter, meta); err != nil {
			return
		}
	}
	if q.Order != nil {
		if tx, err = ApplyOrder(tx, q.Order); err != nil {
			return
		}
	}

	if q.PageSize > 0 {
		if tx, err = ApplyPagination(tx, q.PageSize, q.PageToken, q.Skip); err != nil {
			return
		}
	} else {
		if q.Skip > 0 {
			tx = d.Offset(int(q.Skip))
		}
		if q.Limit > 0 {
			tx = d.Limit(q.Limit)
		}
	}

	return
}

func ApplyOrder(d *gorm.DB, order *core.Ordering) (*gorm.DB, error) {
	if d != nil && order != nil && len(order.Orders) > 0 {
		return d.Order(order.Format()), nil
	}
	return d, nil
}

func ApplyPagination(d *gorm.DB, pageSize int32, pageToken string, skip int32) (*gorm.DB, error) {
	if d != nil && pageSize > 0 {
		var index int64 = 0
		var err error
		if len(pageToken) > 0 {
			index, err = strconv.ParseInt(pageToken, 10, 64)
			if err != nil {
				return nil, err
			}
		}

		offset := pageSize*int32(index) + skip
		return d.Offset(int(offset)).Limit(int(pageSize)), nil
	}
	return d, nil
}
