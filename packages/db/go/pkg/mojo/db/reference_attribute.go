package db

import (
    "bytes"
    "fmt"
    "strings"
)

const ReferenceAttributeName = "reference"
const ReferenceAttributeFullName = "db.reference"

const (
    foreignKey = "foreign_key"
    joinTable  = "join_table"
)

type ReferenceAttribute struct {
    ForeignKey string
    JoinTable  string
}

func (r ReferenceAttribute) Format() string {
    var kvs [][]string

    if len(r.ForeignKey) > 0 {
        kvs = append(kvs, []string{foreignKey, r.ForeignKey})
    }
    if len(r.JoinTable) > 0 {
        kvs = append(kvs, []string{joinTable, r.JoinTable})
    }

    buffer := bytes.NewBuffer(nil)
    for i, kv := range kvs {
        if i > 0 {
            buffer.WriteByte(',')
        }
        buffer.WriteString(kv[0])
        buffer.WriteByte(':')
        buffer.WriteString(kv[1])
    }
    return buffer.String()
}

func (r *ReferenceAttribute) Parse(value string) error {
    if r != nil {
        segments := strings.Split(value, ",")
        for _, segment := range segments {
            kv := strings.Split(segment, ":")
            if len(kv) != 2 {
                return fmt.Errorf("failed to parse reference attribute: %s", value)
            }
            switch kv[0] {
            case foreignKey:
                r.ForeignKey = kv[1]
            case joinTable:
                r.JoinTable = kv[1]
            }
        }
    }
    return nil
}
