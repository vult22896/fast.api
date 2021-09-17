package server

import (
	"fast.bibabo.vn/controllers"
	"fast.bibabo.vn/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("api/v1")
	{
		userGroup := v1.Group("user").Use(middlewares.Auth)
		{
			user := controllers.UserController{}
			userGroup.GET("", user.Index)
			userGroup.GET("me", user.Me)
			userGroup.GET(":id", user.Show)
			userGroup.GET(":id/posts", user.ListPost)
		}
		event := controllers.EventController{}
		v1.GET("event", event.Index)
	}
	return router

}
