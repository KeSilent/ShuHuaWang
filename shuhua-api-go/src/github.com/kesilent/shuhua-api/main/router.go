package main

import (
	"github.com/gin-gonic/gin"
	apis "github.com/kesilent/shuhua-api/apis"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	//添加中间件
	router.Use(Cors())
	v1 := router.Group("v1")
	{
		v1.GET("/", apis.IndexApi)
		v1.POST("/news", apis.AddNewsApi)
		v1.GET("/news", apis.GetNewsApi)
	}
	return router
}
