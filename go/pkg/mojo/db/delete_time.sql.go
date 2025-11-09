// Copyright (c) 2013-NOW  Jinzhu <wosmvp@gmail.com>
//
// The MIT License (MIT)
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// source: https://github.com/go-gorm/soft_delete

package db

import (
	"database/sql/driver"
	"fmt"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"reflect"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

func (x *DeleteTime) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return core.Now().ToTime(), nil
}

func (x *DeleteTime) Scan(src interface{}) error {
	if v := reflect.ValueOf(src); !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch val := src.(type) {
	case time.Time:
		x.Val = core.FromTime(val)
	default:
		return fmt.Errorf("could NOT decode type %T -> %T", src, x)
	}

	return nil
}

func (*DeleteTime) QueryClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{SoftDeleteQueryClause{Field: f}}
}

func (*DeleteTime) UpdateClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{SoftDeleteUpdateClause{Field: f}}
}

func (*DeleteTime) DeleteClauses(f *schema.Field) []clause.Interface {
	settings := schema.ParseTagSetting(f.TagSettings["SOFTDELETE"], ",")
	softDeleteClause := SoftDeleteDeleteClause{
		Field: f,
		Flag:  settings["FLAG"] != "",
	}

	// flag is much more priority
	if v := settings["DeleteTimeFIELD"]; v != "" { // DeleteTimeField
		softDeleteClause.StateField = f.Schema.LookUpField(v)
	}

	return []clause.Interface{softDeleteClause}
}

type SoftDeleteQueryClause struct {
	Field *schema.Field
}

func (sd SoftDeleteQueryClause) Name() string {
	return ""
}

func (sd SoftDeleteQueryClause) Build(clause.Builder) {
}

func (sd SoftDeleteQueryClause) MergeClause(*clause.Clause) {
}

func (sd SoftDeleteQueryClause) ModifyStatement(stmt *gorm.Statement) {
	if _, ok := stmt.Clauses["soft_delete_enabled"]; !ok && !stmt.Statement.Unscoped {
		if c, ok := stmt.Clauses["WHERE"]; ok {
			if where, ok := c.Expression.(clause.Where); ok && len(where.Exprs) > 1 {
				for _, expr := range where.Exprs {
					if orCond, ok := expr.(clause.OrConditions); ok && len(orCond.Exprs) == 1 {
						where.Exprs = []clause.Expression{clause.And(where.Exprs...)}
						c.Expression = where
						stmt.Clauses["WHERE"] = c
						break
					}
				}
			}
		}

		stmt.AddClause(clause.Where{Exprs: []clause.Expression{
			clause.Eq{Column: clause.Column{Table: clause.CurrentTable, Name: sd.Field.DBName}, Value: nil},
		}})

		stmt.Clauses["soft_delete_enabled"] = clause.Clause{}
	}
}

type SoftDeleteUpdateClause struct {
	Field *schema.Field
}

func (sd SoftDeleteUpdateClause) Name() string {
	return ""
}

func (sd SoftDeleteUpdateClause) Build(clause.Builder) {
}

func (sd SoftDeleteUpdateClause) MergeClause(*clause.Clause) {
}

func (sd SoftDeleteUpdateClause) ModifyStatement(stmt *gorm.Statement) {
	if stmt.SQL.String() == "" {
		if _, ok := stmt.Clauses["WHERE"]; stmt.DB.AllowGlobalUpdate || ok {
			SoftDeleteQueryClause(sd).ModifyStatement(stmt)
		}
	}
}

type SoftDeleteDeleteClause struct {
	Field      *schema.Field
	Flag       bool
	StateField *schema.Field
}

func (sd SoftDeleteDeleteClause) Name() string {
	return ""
}

func (sd SoftDeleteDeleteClause) Build(clause.Builder) {
}

func (sd SoftDeleteDeleteClause) MergeClause(*clause.Clause) {
}

func (sd SoftDeleteDeleteClause) ModifyStatement(stmt *gorm.Statement) {
	if stmt.SQL.Len() == 0 && !stmt.Statement.Unscoped {
		curTime := stmt.DB.NowFunc()
		//if sd.Flag {
		//    set := clause.Set{{Column: clause.Column{Name: sd.Field.DBName}, Value: 1}}
		//    stmt.SetColumn(sd.Field.DBName, 1, true)
		//
		//    if sd.Field != nil {
		//        set = append(set, clause.Assignment{Column: clause.Column{Name: sd.Field.DBName}, Value: curTime.Unix()})
		//        stmt.SetColumn(sd.Field.DBName, curTime, true)
		//    }
		//
		//    stmt.AddClause(set)
		//} else {
		//    var curUnix int64 = 0
		//    if sd.TimeType == schema.UnixNanosecond {
		//        curUnix = curTime.UnixNano()
		//    } else if sd.TimeType == schema.UnixMillisecond {
		//        curUnix = curTime.UnixNano() / 1e6
		//    } else {
		//        curUnix = curTime.Unix()
		//    }
		stmt.AddClause(clause.Set{{Column: clause.Column{Name: sd.Field.DBName}, Value: curTime}})
		stmt.SetColumn(sd.Field.DBName, curTime, true)
		//}

		if stmt.Schema != nil {
			_, queryValues := schema.GetIdentityFieldValuesMap(stmt.Context, stmt.ReflectValue, stmt.Schema.PrimaryFields)
			column, values := schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)

			if len(values) > 0 {
				stmt.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
			}

			if stmt.ReflectValue.CanAddr() && stmt.Dest != stmt.Model && stmt.Model != nil {
				_, queryValues = schema.GetIdentityFieldValuesMap(stmt.Context, reflect.ValueOf(stmt.Model), stmt.Schema.PrimaryFields)
				column, values = schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)

				if len(values) > 0 {
					stmt.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
				}
			}
		}

		SoftDeleteQueryClause{Field: sd.Field}.ModifyStatement(stmt)
		stmt.AddClauseIfNotExists(clause.Update{})
		stmt.Build(stmt.DB.Callback().Update().Clauses...)
	}
}
