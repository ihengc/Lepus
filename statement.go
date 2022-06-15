package Lepus

import (
	"fmt"
	reflect "reflect"
	"strings"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/15 0015 14:36
* @version: 1.0
* @description:
*********************************************************/

type Field struct {
	Name string
	Type reflect.Type
}

type Schema struct {
	Name string
}

type Statement struct {
	Dest   interface{}     // 操作的目标
	SQL    strings.Builder // SQL
	Fields []*Field
	Schema *Schema
}

func (stmt *Statement) ParseField() {
	refValue := reflect.Indirect(reflect.ValueOf(stmt.Dest))
	refType := refValue.Type()

	switch refValue.Kind() {
	case reflect.Struct:
		for i := 0; i < refType.NumField(); i++ {
			stmt.Fields = append(stmt.Fields, &Field{refType.Field(i).Name, refType.Field(i).Type})
		}
		stmt.Schema = &Schema{Name: refType.Name()}
	}
}

func (stmt *Statement) WriteByte(c byte) {
	stmt.SQL.WriteByte(c)
}

func (stmt *Statement) WriteString(s string) {
	stmt.SQL.WriteString(s)
}

func (stmt *Statement) WriteQuoted(v interface{}) {
	switch v.(type) {
	case *Field:
		field := v.(*Field)
		stmt.WriteString(fmt.Sprintf("`%s`", strings.ToLower(field.Name)))
	case *Schema:
		schema := v.(*Schema)
		stmt.WriteString(fmt.Sprintf("`%s`", strings.ToLower(schema.Name)))
	}
}
