package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "swd:swd12345@tcp(120.27.17.206:3306)/bigdata1?charset=utf8")
	CheckErr(err)
	SqlDB.SetMaxOpenConns(2000)
	SqlDB.SetMaxIdleConns(1000)
	err = SqlDB.Ping()
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}
