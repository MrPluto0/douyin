package controller

import (
	"douyin/app/define"
	"douyin/app/service"
	"douyin/utils/response"

	"github.com/gin-gonic/gin"
)

type videoApi struct{}

var VideoApi *videoApi

func init() {
	VideoApi = &videoApi{}
}

func (v *videoApi) Feed(c *gin.Context) {
	var req define.FeedReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response.Resp(c, define.FeedRes{
			Errno: response.ErrValidation.Extend(err),
		})
	} else {
		response.Resp(c, service.VideoService.Feed(req))
	}
}
