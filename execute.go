package Lepus

import "fmt"

/********************************************************
* @author: Ihc
* @date: 2022/6/15 0015 14:43
* @version: 1.0
* @description:
*********************************************************/

type Executor struct {
}

func (exec *Executor) Execute(lepus *Lepus) (*Lepus, error) {
	SQL := lepus.Stmt.SQL.String()
	if SQL == "" {
		return nil, fmt.Errorf("null sql error")
	}
	// TODO STOP
	return nil, nil
	if ret, err := lepus.raw.Exec(SQL); err != nil {
		return nil, err
	} else {
		if rowsAffected, err := ret.RowsAffected(); err != nil {
			return nil, err
		} else {
			lepus.RowsAffected = rowsAffected
		}
	}
	return lepus, nil
}
