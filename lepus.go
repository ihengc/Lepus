package Lepus

import "database/sql"

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/14 20:57
 * @description:
 ***************************************************************/

type DB struct {
	raw *sql.DB
}

func Open(driverName, dataSourceName string) *DB {
	raw, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil
	}
	lepusDB := &DB{}
	lepusDB.raw = raw
	return lepusDB
}
