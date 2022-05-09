package db_client

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sql.DB

func InitializeDBConnection() {
	db, err := sqlx.Open(driverName: "mysql", dataSourceName: "root:@tcp(localhost:3306)/test_db?parseTime=true")
	if err != nil {
		panic(err.Error)
	}
	err = d.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}