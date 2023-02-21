package middleware

import (
	"douyin/utils/check"
	"douyin/utils/file"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func logFormat(param gin.LogFormatterParams) string {
	// your custom format
	return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC3339Nano),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func LogMiddleware() gin.HandlerFunc {
	logPath := viper.GetString("root") + viper.GetString("server.log_path")
	writer, err := file.OpenFile_A(logPath)
	check.CheckPanicErr(err)

	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(writer, os.Stdout)
	logMdw := gin.LoggerWithFormatter(logFormat)
	return logMdw
}
