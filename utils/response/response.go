// Unused module in this project
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"status_code"`
	Msg  string      `json:"status_msg"`
	Data interface{} `json:"data"`
}

func Resp(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func RespJSON(ctx *gin.Context, err Errno, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: OK.Code,
		Msg:  OK.Msg,
		Data: data,
	})
}

func RespSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: OK.Code,
		Msg:  OK.Msg,
		Data: data,
	})
}
