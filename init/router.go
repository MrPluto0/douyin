package init

import (
	API "douyin/app/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	mainRouter := r.Group("/douyin")

	// User Group
	userGroup := mainRouter.Group("user")
	{
		userGroup.POST("login", API.UserApi.Login)
		userGroup.POST("register")
		userGroup.GET("")
	}

	// Feed
	mainRouter.GET("feed")

	// Publish Group
	publishGroup := mainRouter.Group("/publish")
	{
		publishGroup.POST("action")
		publishGroup.GET("list")
	}

	// Favorite Group
	favoriteGroup := mainRouter.Group("/favorite")
	{
		favoriteGroup.POST("action")
		favoriteGroup.GET("list")
	}

	// Comment Group
	commentGroup := mainRouter.Group("/comment")
	{
		commentGroup.POST("action")
		commentGroup.GET("list")
	}
}
