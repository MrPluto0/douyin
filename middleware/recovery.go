package middleware

import (
	"douyin/utils/response"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		response.Resp(c, err)
	})
}
