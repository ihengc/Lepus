package Lepus

import "fmt"

/********************************************************
* @author: Ihc
* @date: 2022/6/15 0015 14:35
* @version: 1.0
* @description:
*********************************************************/

type SQLBuildType byte

const (
	Create SQLBuildType = iota + 1
	Insert
	Update
	Delete
)

type SQLBuilder struct {
}

func (s *SQLBuilder) Create(lepus *Lepus) *Executor {
	BuildInsertSQL(lepus)
	return &Executor{}
}

// BuildInsertSQL 构建插入SQL
func BuildInsertSQL(lepus *Lepus) {
	// INSERT INTO `TABLE` (`FIELD`, `FIELD`, `FIELD`) VALUES(),();
	stmt := lepus.Stmt
	stmt.WriteString("INSERT INTO ")
	fmt.Println(stmt.Schema.Name)
	stmt.WriteQuoted(stmt.Schema)
	stmt.WriteByte('(')
	for idx, field := range stmt.Fields {
		if idx > 0 {
			stmt.WriteByte(',')
		}
		stmt.WriteQuoted(field)
	}
	stmt.WriteByte(')')
	stmt.WriteString(" VALUES")
	fmt.Println(stmt.SQL.String())
}

func (s *SQLBuilder) Delete(lepus *Lepus) *Executor {
	return &Executor{}
}
