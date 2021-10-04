package server

import (
	"fast.bibabo.vn/controllers"
	"fast.bibabo.vn/database"
	"fast.bibabo.vn/middlewares"
	"fast.bibabo.vn/repositories"
	puService "fast.bibabo.vn/services/post_services"
	uService "fast.bibabo.vn/services/user_services"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	db := database.GetInstanceMysql().Connect()
	myCache := database.GetInstanceRedis().Caching()
	mongo := database.GetInstanceMongo().Connect().DB("bibabo")
	v1 := router.Group("api/v1")
	{
		userGroup := v1.Group("user").Use(middlewares.Auth(db, myCache))
		{
			userService := uService.GetIntanceUserService(db, myCache)
			postUserService := puService.GetPostUserService(db, myCache, mongo)
			userController := controllers.NewUserController(db, myCache, userService, postUserService)
			userGroup.GET("", userController.Index)
			userGroup.GET("me", userController.Me)
			userGroup.GET(":id", userController.Show)
			userGroup.GET(":id/posts", userController.ListPost)
		}

		groupGroup := v1.Group("group")
		{
			groupRepo := repositories.InstanceGroupRepository(db, myCache)
			groupController := controllers.InstanceGroupController(groupRepo)
			groupGroup.GET("list", groupController.FindAll)
		}
	}
	return router

}
