package persist

import "database/sql"

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/20 20:53
 * @description:
 ***************************************************************/

type DB struct {
	rawDB *sql.DB
}

func Open(driverName string, dataSourceName string) *DB {
	rawDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil
	}
	db := &DB{}
	db.rawDB = rawDB
	return db
}
