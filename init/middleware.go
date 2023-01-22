package init

import (
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func InitMiddleWare(r *gin.Engine) {
	r.Use(middleware.LogMiddleWare())
}