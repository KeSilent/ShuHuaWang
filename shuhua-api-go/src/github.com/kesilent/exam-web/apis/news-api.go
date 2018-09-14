package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kesilent/exam-web/models"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

//获取新闻条数
func GetNewsApi(c *gin.Context) {
	newlist, err := models.GetNewsForNewTime()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
	}
	c.JSON(http.StatusOK, newlist)
}
