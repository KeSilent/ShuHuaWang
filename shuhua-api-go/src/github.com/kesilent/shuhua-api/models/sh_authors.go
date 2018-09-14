package models

import (
	"time"

	db "github.com/kesilent/shuhua-api/database"
)

type SH_Authors struct {
	Id           int       `json:"aId"`
	Name         string    `json:"aName"`
	Title        string    `json:"aTitle"`
	Photo        string    `json:"aPhoto"`
	TitleImage   string    `json:"aTitleImage"`
	Introduction string    `json:"aIntroduction"`
	CreateTime   time.Time `json:"aCreateTime"`
	ModifyTime   time.Time `json:"aModifyTime"`
	Status       int       `json:"aStatus"`
	IsBest       int       `json:"aIsBest"`
	IsSort       int       `json:"aIsSort"`
}
type AuthorList []*SH_Authors

//添加作者
func AddAuthor(model *SH_Authors) (int64, error) {
	stmt, err := db.SqlDB.Prepare("INSERT INTO `shuhuawang`.`sh_authors`(`aId`, `aName`, `aTitle`, `aPhoto`, `aTitleImage`, `aIntroduction`, `aCreateTime`, `aModifyTime`, `aStatus`, `aIsBest`, `aIsSort`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	db.CheckErr(err)
	res, err := stmt.Exec(model.Id, model.Name, model.Title, model.Photo, model.TitleImage, model.Introduction, model.CreateTime, model.ModifyTime, model.Status, model.IsBest, model.IsSort)
	db.CheckErr(err)

	num, err := res.RowsAffected()
	db.CheckErr(err)
	return num, err
}

//通过ID获取作者
func GetAuthorForId(id int) (*SH_Authors, error) {
	model := new(SH_Authors)
	stmt, err := db.SqlDB.Prepare("SELECT aName, aId, aTitle, aPhoto, aTitleImage, aIntroduction, aCreateTime, aModifyTime, aStatus, aIsBest, aIsSort FROM shuhuawang.sh_authors WHERE aId=?")
	db.CheckErr(err)
	rows, err := stmt.Query(id)
	db.CheckErr(err)

	if rows.Next() {
		err = rows.Scan(&model.Id, &model.Name, &model.Title, &model.Photo, &model.TitleImage, &model.Introduction, &model.CreateTime, &model.ModifyTime, &model.Status, &model.IsBest, &model.IsSort)
		db.CheckErr(err)

	}
	return model, err
}

//通过分页获取作者
//page页数  size每页显示数量
func GetAuthorForPage(page int, size int) (AuthorList, error) {
	result := make(AuthorList, 0)
	stmt, err := db.SqlDB.Prepare("SELECT aName, aId, aTitle, aPhoto, aTitleImage, aIntroduction, aCreateTime, aModifyTime, aStatus, aIsBest, aIsSort FROM shuhuawang.sh_authors ORDER BY aCreateTime desc LIMIT ?,?")
	db.CheckErr(err)
	rows, err := stmt.Query(page, size)
	db.CheckErr(err)

	if rows.Next() {
		model := new(SH_Authors)
		err = rows.Scan(&model.Id, &model.Name, &model.Title, &model.Photo, &model.TitleImage, &model.Introduction, &model.CreateTime, &model.ModifyTime, &model.Status, &model.IsBest, &model.IsSort)
		db.CheckErr(err)
		result = append(result, model)
	}
	return result, err
}
