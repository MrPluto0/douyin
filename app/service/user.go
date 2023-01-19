package service

import (
	"douyin/app/define"
	"douyin/app/models"
	"douyin/utils/response"
)

type userService struct{}

var UserService *userService

func init() {
	UserService = &userService{}
}

func (s *userService) Login(req define.LoginReq) define.LoginRes {
	// Validate
	if matched, _ := req.Validate(); !matched {
		return define.LoginRes{
			Errno: *response.ErrValidation,
		}
	}

	// Run userDao
	dao := models.NewUserDaoInstance()
	user := dao.QueryUser(req.Username)

	if user != nil {
		return define.LoginRes{
			UserId: user.ID,
			Token:  "12",
			Errno:  *response.OK,
		}
	} else {
		return define.LoginRes{
			Errno: *response.ErrUserNotFound,
		}
	}
}
