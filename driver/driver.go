package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB ...
type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ...
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		uname,
		pass,
		host,
		port,
		dbname,
	)
	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}