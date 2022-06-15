package Lepus

import (
	"database/sql"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/14 20:57
 * @description:
 ***************************************************************/

// Lepus 对应数据库
type Lepus struct {
	raw          *sql.DB
	sqlBuilder   *SQLBuilder
	RowsAffected int64
	Stmt         *Statement
}

// Open 获取数据库连接
func Open(driverName, dataSourceName string) (*Lepus, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	lepus := &Lepus{}
	lepus.raw = db
	lepus.Stmt = &Statement{}
	lepus.sqlBuilder = &SQLBuilder{}
	return lepus, nil
}

// Create 新建数据
func (lepus *Lepus) Create(value interface{}) {
	lepus.Stmt.Dest = value
	lepus.sqlBuilder.Create(lepus).Execute(lepus)
}
