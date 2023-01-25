package middleware

import (
	"douyin/utils"
	"douyin/utils/jwt"
	"douyin/utils/response"

	"github.com/gin-gonic/gin"
)

var noAuthRoutes = []string{
	"/douyin/feed",
	"/douyin/user/login",
	"/douyin/user/register",
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// noAuth
		if utils.IsContain(c.Request.URL.Path, noAuthRoutes) {
			return
		}

		token := c.Query("token")
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.Resp(c, response.ErrToken.Extend(err))
			c.Abort()
		} else {
			c.Set("user", claims.User)
		}
	}
}
