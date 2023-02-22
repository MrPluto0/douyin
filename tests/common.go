package tests

import (
	Init "douyin/init"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	engine = gin.New()
	Init.InitConfig()
	Init.InitRedis()
	Init.InitMysql()
	Init.InitRouter(engine)
}

// Create httpexpect instance
func NewHttpExcept(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(engine),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		// Printers: []httpexpect.Printer{
		// 	httpexpect.NewDebugPrinter(t, true),
		// },
	})
}
