package models

import (
	db "github.com/kesilent/shuhua-api/database"
)

//商品分类
type SH_Type struct {
	Id           int        `json:"typeId" form:"typeId"`
	Name         string     `json:"typeName" form:"typeName"`
	ParentId     int        `json:"typeParentId" form:"typeParentId"`
	ChildrenType []*SH_Type `json:"children"`
}
type TypeList []*SH_Type

//添加分类
func AddType(model *SH_Type) (int64, error) {
	stmt, err := db.SqlDB.Prepare("INSERT INTO `shuhuawang`.`sh_type`(`typeId`, `typeName`, `typeParentId`) VALUES (?, ?, ?)")
	db.CheckErr(err)
	res, err := stmt.Exec(model.Id, model.Name, model.ParentId)
	db.CheckErr(err)
	num, err := res.RowsAffected()
	db.CheckErr(err)
	return num, err
}

//通过父类ID获取分类数据
func GetTypeForParId(parentId int64) (TypeList, error) {
	typeList := make(TypeList, 0)

	stmt, err := db.SqlDB.Prepare("SELECT * FROM sh_type WHERE typeParentId=?")
	db.CheckErr(err)
	row, err := stmt.Query(parentId)
	db.CheckErr(err)
	defer stmt.Close()

	for row.Next() {
		item := new(SH_Type)
		err := row.Scan(&item.Id, &item.Name, &item.ParentId)
		db.CheckErr(err)
		typeList = append(typeList, item)
	}
	return typeList, err
}

//更新分类信息
func UpdateType(model *SH_Type) (int64, error) {
	stmt, err := db.SqlDB.Prepare("UPDATE `shuhuawang`.`sh_type` SET `typeName` = ?, `typeParentId` = ? WHERE `typeId` =?;")
	db.CheckErr(err)
	res, err := stmt.Exec(model.Name, model.ParentId, model.Id)
	db.CheckErr(err)
	num, err := res.RowsAffected()
	db.CheckErr(err)
	return num, err
}
