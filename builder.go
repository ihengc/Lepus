package Lepus

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

}
