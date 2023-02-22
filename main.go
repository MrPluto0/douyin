package main

import (
	Init "douyin/init"
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	pprof.Register(r)

	Init.InitConfig()
	Init.InitMysql()
	Init.InitRedis()
	Init.InitRouter(r)

	addr := fmt.Sprintf("%v:%v", viper.Get("server.host"), viper.Get("server.port"))
	r.Run(addr)
}
