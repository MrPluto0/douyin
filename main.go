package main

import (
	Init "douyin/init"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	Init.InitConfig()
	Init.InitRouter(r)
	Init.InitMysql()

	r.Run()
}
