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
	// Run userDao
	dao := models.NewUserDaoInstance()
	user := dao.QueryUser(req.Username)

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

func (s *userService) Register(req define.RegisterReq) define.RegisterRes {
	dao := models.NewUserDaoInstance()
	user := dao.QueryUser(req.Username)

	if user.ID != models.EmptyID {
		return define.RegisterRes{
			Errno: *response.ErrUserExisted,
		}
	}

	rows := dao.CreateUser(req.Username, req.Password)
	if rows != 1 {
		return define.RegisterRes{
			Errno: *response.ErrCreateFailed,
		}
	}

	return define.RegisterRes{
		UserId: 0,
		Token:  "123",
		Errno:  *response.OK,
	}
}
