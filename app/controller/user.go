package controller

import (
	"douyin/app/define"
	"douyin/app/service"
	"douyin/utils/response"

	"github.com/gin-gonic/gin"
)

type userApi struct{}

var UserApi *userApi

func init() {
	UserApi = &userApi{}
}

func (u *userApi) Login(c *gin.Context) {
	var req define.LoginReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response.Resp(c, define.LoginRes{
			Errno: response.ErrValidation.Extend(err),
		})
	} else {
		req.Validate()
		response.Resp(c, service.UserService.Login(req))
	}
}

func (u *userApi) Register(c *gin.Context) {
	var req define.RegisterReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response.Resp(c, define.LoginRes{
			Errno: response.ErrValidation.Extend(err),
		})
	} else {
		req.Validate()
		response.Resp(c, service.UserService.Register(req))
	}
}
