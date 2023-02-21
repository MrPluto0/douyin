package controller

import (
	"douyin/app/define"
	"douyin/app/models"
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
		response.Resp(c, service.UserService.Register(req))
	}
}

func (u *userApi) UserInfo(c *gin.Context) {
	var req define.UserInfoReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response.Resp(c, define.UserInfoRes{
			Errno: response.ErrValidation.Extend(err),
		})
		return
	}

	userAny, _ := c.Get("user")
	user := userAny.(models.User)
	if user.ID != req.UserId {
		response.Resp(c, define.UserInfoRes{
			Errno: *response.ErrToken,
		})
		return
	}

	response.Resp(c, service.UserService.UserInfo(req))
}
