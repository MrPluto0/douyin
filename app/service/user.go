package service

import (
	"douyin/app/define"
	"douyin/app/models"
	"douyin/utils/jwt"
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
	user, err := dao.QueryUser(req.Username)
	if err != nil {
		return define.LoginRes{
			Errno: response.ErrUserNotFound.Extend(err),
		}
	}

	// Core
	if req.Password != user.Password {
		return define.LoginRes{
			Errno: *response.ErrPwdWrong,
		}
	}

	// Generate token
	token, err := jwt.GenerateToken(user)
	if err != nil {
		return define.LoginRes{
			Errno: response.ErrToken.Extend(err),
		}
	}

	// return successfully
	return define.LoginRes{
		UserId: user.ID,
		Token:  token,
		Errno:  *response.OK,
	}
}
