package init

import (
	API "douyin/app/controller"
	"douyin/app/models"
	"douyin/middleware"
	"douyin/utils/response"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// Global Middleware
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.AuthMiddleware())
	r.Use(middleware.LogMiddleware())

	mainRouter := r.Group("/douyin")

	// User Group
	{
		userGroup := mainRouter.Group("user")
		userGroup.POST("login", API.UserApi.Login)
		userGroup.POST("register", API.UserApi.Register)
		userGroup.GET("")
	}

	// Feed
	mainRouter.GET("feed")

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

	// test route for token
	mainRouter.GET("test", func(ctx *gin.Context) {
		if user, ok := ctx.Get("user"); !ok {
			panic(*response.ErrUserNotFound)
		} else {
			u := user.(models.User)
			u.Name = "test panic and token"
			panic(u)
		}
	})
}
