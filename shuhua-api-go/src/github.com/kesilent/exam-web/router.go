package main

import (
	"github.com/gin-gonic/gin"
	apis "github.com/kesilent/exam-web/apis"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	//添加中间件
	router.Use(Cors())
	v1 := router.Group("v1")
	{
		v1.GET("/", apis.GetNewsApi)
	}
	return router
}
