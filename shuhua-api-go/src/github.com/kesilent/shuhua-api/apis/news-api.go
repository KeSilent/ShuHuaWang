package apis

import (
	"net/http"
	"strconv"
	"time"

	"github.com/kesilent/shuhua-api/models"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

//添加新闻
func AddNewsApi(c *gin.Context) {
	news := new(models.SH_News)
	news.NTitle = c.PostForm("title")
	news.NContent = c.PostForm("content")
	news.NAuthor = c.PostForm("author")
	news.NTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	id := models.AddNews(news)
	c.String(http.StatusOK, strconv.FormatInt(id, 10))
}

//获取新闻条数
func GetNewsApi(c *gin.Context) {
	quantity, err := strconv.Atoi(c.Query("number"))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
	} else {
		newlist, err := models.GetNewsForQuantity(quantity)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, err)
		}
		c.JSON(http.StatusOK, newlist)
	}
}
