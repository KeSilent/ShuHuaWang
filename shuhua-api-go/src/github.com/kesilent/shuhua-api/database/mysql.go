package database

import (
	"bytes"
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/shuhuawang?charset=utf8")
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

//产生编号
func GenerateID(randLength int) (result string) {
	var num string = "0123456789"

	b := bytes.Buffer{}
	b.WriteString(num)
	var str = b.String()
	var strLen = len(str)
	if strLen == 0 {
		result = "0"
		return
	}

	sfe := time.Now().UnixNano()
	rand.Seed(sfe)
	b = bytes.Buffer{}
	for i := 0; i < randLength; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	result = b.String()
	return
}
