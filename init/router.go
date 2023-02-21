package init

import (
	API "douyin/app/controller"
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// Global Middleware
	r.Use(gin.Recovery())
	r.Use(middleware.AuthMiddleware())
	r.Use(middleware.LogMiddleware())

	r.Static("/static", "./assets")

	mainRouter := r.Group("/douyin")

	// User Group
	{
		userGroup := mainRouter.Group("user")
		userGroup.POST("login", API.UserApi.Login)
		userGroup.POST("register", API.UserApi.Register)
		userGroup.GET("", API.UserApi.UserInfo)
	}

	// Feed
	mainRouter.GET("feed", API.VideoApi.Feed)

	// Publish Group
	{
		publishGroup := mainRouter.Group("/publish")
		publishGroup.POST("action")
		publishGroup.GET("list")
	}

	// Favorite Group
	{
		favoriteGroup := mainRouter.Group("/favorite")
		favoriteGroup.POST("action")
		favoriteGroup.GET("list")
	}

	// Comment Group
	{
		commentGroup := mainRouter.Group("/comment")
		commentGroup.POST("action")
		commentGroup.GET("list")
	}
}
