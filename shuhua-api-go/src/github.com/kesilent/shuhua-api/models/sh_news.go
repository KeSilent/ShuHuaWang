package models

import (
	"time"

	db "github.com/kesilent/shuhua-api/database"
)

type SH_News struct {
	NewsId   int       `json:"newsid" form:"newsId"`
	NTitle   string    `json:"ntitle" form:"nTitle"`
	NContent string    `json:"ncontent" form:"nContent"`
	NAuthor  string    `json:"nauthor" form:"nAuthor"`
	NTime    time.Time `json:"ntime" form:"nTime"`
}
type NewList []*SH_News

//添加新闻
func AddNews(p *SH_News) (int64, error) {
	stmt, err := db.SqlDB.Prepare("INSERT INTO sh_news(nTitle, nContent,nAuthor,nTime) VALUES (?, ?, ?, ?)")
	db.CheckErr(err)

	res, err := stmt.Exec(p.NTitle, p.NContent, p.NAuthor, p.NTime)
	db.CheckErr(err)

	id, err := res.LastInsertId()
	db.CheckErr(err)

	return id, err
}

//获取新闻
//quantity： 数量
func GetNewsForQuantity(quantity int) (NewList, error) {
	news := make(NewList, 0)
	stmt, err := db.SqlDB.Prepare("SELECT newsId,nContent,nTitle,nAuthor,nTime FROM sh_news ORDER BY nTime DESC LIMIT ? ")
	db.CheckErr(err)
	rows, err := stmt.Query(quantity)
	db.CheckErr(err)
	defer stmt.Close()

	for rows.Next() {
		var nTime string
		item := new(SH_News)
		err := rows.Scan(&item.NewsId, &item.NContent, &item.NTitle, &item.NAuthor, &nTime)
		db.CheckErr(err)
		item.NTime, _ = time.Parse("2006-01-02 15:04:05", nTime)

		news = append(news, item)
	}
	db.CheckErr(err)
	return news, err
}
