package models

import (
	"bytes"
	"database/sql"
	"math/rand"
	"strings"
	"time"

	db "github.com/kesilent/exam-web/database"
)

type ExamClass struct {
	Id          int64     `json:"id" form:"id"`
	Class_id    int64     `json:"class_id" form:"class_id"`
	Exam_id     int64     `json:"exam_id" form:"exam_id"`
	Grade_id    int64     `json:"grade_id" form:"grade_id"`
	Create_time time.Time `json:"create_time" form:"create_time"`
}

func GetExamClass(exam_id int64) bool {
	var id string
	err := db.SqlDB.QueryRow("SELECT id FROM edu_exam_classes where exam_id=?", exam_id).Scan(&id)
	if err == sql.ErrNoRows {
		return true
	} else {
		return false
	}
}

func AddExamClass(exam_id int64) int64 {
	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	stmt, err := db.SqlDB.Prepare("INSERT INTO edu_exam_classes(id,class_id,exam_id,grade_id,create_time) VALUES (?,?,?,?,?)")
	db.CheckErr(err)

	idin := RandomString(16, "0")

	res, err := stmt.Exec(idin, "233238698316005376", exam_id, "1447314780711813127", datetime)
	db.CheckErr(err)

	id, err := res.LastInsertId()
	db.CheckErr(err)

	return id
}
func RandomString(randLength int, randType string) (result string) {
	var num string = "0123456789"
	var lower string = "abcdefghijklmnopqrstuvwxyz"
	var upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := bytes.Buffer{}
	if strings.Contains(randType, "0") {
		b.WriteString(num)
	}
	if strings.Contains(randType, "a") {
		b.WriteString(lower)
	}
	if strings.Contains(randType, "A") {
		b.WriteString(upper)
	}
	var str = b.String()
	var strLen = len(str)
	if strLen == 0 {
		result = ""
		return
	}

	sfe := rand.Int63()
	rand.Seed(sfe)
	b = bytes.Buffer{}
	for i := 0; i < randLength; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	result = b.String()
	return
}
