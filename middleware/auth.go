package middleware

import (
	"douyin/utils/arr"
	"douyin/utils/jwt"
	"douyin/utils/response"

	"github.com/gin-gonic/gin"
)

var AuthRoutes = []string{
	"/douyin/user",
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !arr.IsContain(c.Request.URL.Path, AuthRoutes) {
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
