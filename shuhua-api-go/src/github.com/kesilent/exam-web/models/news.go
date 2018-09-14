package models

import (
	db "github.com/kesilent/exam-web/database"
)

type Exam_Info struct {
	Exam_id   int64  `json:"exam_id" form:"exam_id"`
	Exam_name string `json:"exam_name" form:"exam_name"`
	Exam_path string `json:"exam_path" form:"sheet_pdf_path"`
}
type Exam_InfoList []*Exam_Info

//获取最新时间的试卷
func GetNewsForNewTime() (Exam_InfoList, error) {
	news := make(Exam_InfoList, 0)
	var temp_path string
	stmt, err := db.SqlDB.Prepare("SELECT edu_exam_info.exam_id, exam_name, sheet_pdf_path FROM edu_exam_info JOIN edu_exam_paper_content ON edu_exam_info.exam_id = edu_exam_paper_content.exam_id where edu_exam_info.create_time=(select create_time from edu_exam_info order by create_time desc limit 1)")
	db.CheckErr(err)
	rows, err := stmt.Query()
	db.CheckErr(err)
	defer stmt.Close()

	for rows.Next() {
		item := new(Exam_Info)
		err := rows.Scan(&item.Exam_id, &item.Exam_name, &temp_path)
		db.CheckErr(err)
		item.Exam_path = "http://120.27.17.206" + temp_path

		news = append(news, item)
		UpdateStatus(item.Exam_id)
		if GetExamClass(item.Exam_id) {
			AddExamClass(item.Exam_id)
		}
	}
	db.CheckErr(err)
	return news, err
}

//更新状态
func UpdateStatus(examid int64) {
	stmt, err := db.SqlDB.Prepare("UPDATE edu_exam_info SET `status`=1 WHERE exam_id=?")
	db.CheckErr(err)
	stmt.Query(examid)
}
