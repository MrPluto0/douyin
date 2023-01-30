package main

import (
	Init "douyin/init"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.New()

	Init.InitConfig()
	Init.InitMysql()
	Init.InitRouter(r)

	addr := fmt.Sprintf("%v:%v", viper.Get("server.host"), viper.Get("server.port"))
	r.Run(addr)
}
