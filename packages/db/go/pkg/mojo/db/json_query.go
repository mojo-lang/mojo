// The MIT License (MIT)

// Copyright (c) 2013-NOW  Jinzhu <wosmvp@gmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

package db

import (
    "fmt"
    "strings"

    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

// JSONQueryExpression json query expression, implements clause.Expression interface to use as querier
type JSONQueryExpression struct {
    column      string
    keys        []string
    hasKeys     bool
    equals      bool
    equalsValue interface{}
}

// JSONQuery query column as json
func JSONQuery(column string) *JSONQueryExpression {
    return &JSONQueryExpression{column: column}
}

// HasKey returns clause.Expression
func (e *JSONQueryExpression) HasKey(keys ...string) *JSONQueryExpression {
    e.keys = keys
    e.hasKeys = true
    return e
}

// Equals Keys returns clause.Expression
func (e *JSONQueryExpression) Equals(value interface{}, keys ...string) *JSONQueryExpression {
    e.keys = keys
    e.equals = true
    e.equalsValue = value
    return e
}

// Build implements clause.Expression
func (e *JSONQueryExpression) Build(builder clause.Builder) {
    if stmt, ok := builder.(*gorm.Statement); ok {
        switch stmt.Dialector.Name() {
        case "mysql", "sqlite":
            switch {
            case e.hasKeys:
                if len(e.keys) > 0 {
                    builder.WriteString("JSON_EXTRACT(" + stmt.Quote(e.column) + ",")
                    builder.AddVar(stmt, "$."+strings.Join(e.keys, "."))
                    builder.WriteString(") IS NOT NULL")
                }
            case e.equals:
                if len(e.keys) > 0 {
                    builder.WriteString("JSON_EXTRACT(" + stmt.Quote(e.column) + ",")
                    builder.AddVar(stmt, "$."+strings.Join(e.keys, "."))
                    builder.WriteString(") = ")
                    if _, ok := e.equalsValue.(bool); ok {
                        builder.WriteString(fmt.Sprint(e.equalsValue))
                    } else {
                        stmt.AddVar(builder, e.equalsValue)
                    }
                }
            }
        case "postgres":
            switch {
            case e.hasKeys:
                if len(e.keys) > 0 {
                    stmt.WriteQuoted(e.column)
                    stmt.WriteString("::jsonb")
                    for _, key := range e.keys[0 : len(e.keys)-1] {
                        stmt.WriteString(" -> ")
                        stmt.AddVar(builder, key)
                    }

                    stmt.WriteString(" ? ")
                    stmt.AddVar(builder, e.keys[len(e.keys)-1])
                }
            case e.equals:
                if len(e.keys) > 0 {
                    builder.WriteString(fmt.Sprintf("json_extract_path_text(%v::json,", stmt.Quote(e.column)))

                    for idx, key := range e.keys {
                        if idx > 0 {
                            builder.WriteByte(',')
                        }
                        stmt.AddVar(builder, key)
                    }
                    builder.WriteString(") = ")

                    if _, ok := e.equalsValue.(string); ok {
                        stmt.AddVar(builder, e.equalsValue)
                    } else {
                        stmt.AddVar(builder, fmt.Sprint(e.equalsValue))
                    }
                }
            }
        }
    }
}
