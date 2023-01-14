package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Age      uint8  `json:"age" binding:"gte=1,lte=120"`
}

func main() {
	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		var req Request
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, req)
	})
	server.Run()
}
