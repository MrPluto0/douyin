package tests

import (
	"douyin/app/define"
	"douyin/app/models"
	"douyin/utils/response"

	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	e := NewHttpExcept(t)

	// Exp1: empty query params
	e.POST("/douyin/user/login").
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.ErrValidation.Code)

	// Exp2: password validate error
	e.POST("/douyin/user/login").WithQueryObject(define.LoginReq{Username: "Pluto", Password: "123abc"}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.ErrValidation.Code)

	// Exp3: user not found
	e.POST("/douyin/user/login").WithQueryObject(define.LoginReq{Username: "damn", Password: "123abC"}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.ErrDatabase.Code)

	// Exp4: wrong password
	e.POST("/douyin/user/login").WithQueryObject(define.LoginReq{Username: "Gypsophlia", Password: "123aC"}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.ErrPwdWrong.Code)

	// Exp5: right request
	e.POST("/douyin/user/login").WithQueryObject(define.LoginReq{Username: "Gypsophlia", Password: "123abC"}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.OK.Code).
		ContainsKey("user_id").
		ContainsKey("token")
}

func TestRegister(t *testing.T) {
	e := NewHttpExcept(t)

	e.POST("/douyin/user/register").WithQueryObject(define.RegisterReq{Username: "abc", Password: "123abc"}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.ErrValidation.Code)

	e.POST("/douyin/user/register").WithQueryObject(define.RegisterReq{Username: "abc", Password: "123Abc"}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.OK.Code)

	e.POST("/douyin/user/register").WithQueryObject(define.RegisterReq{Username: "abc", Password: "123abC"}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.ErrCreateUser.Code)

	// idempotent: recovery the data before
	dao := models.NewUserDaoInstance()
	dao.Delete("abc")
}

func TestUserInfo(t *testing.T) {
	e := NewHttpExcept(t)

	username := "Pluto"
	password := "123abC"
	user := e.POST("/douyin/user/login").WithQueryObject(define.LoginReq{Username: username, Password: password}).
		Expect().Status(http.StatusOK).JSON().Object()
	token := user.Value("token").String().Raw()
	userId := user.Value("user_id").Number().Raw()

	e.GET("/douyin/user").Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.ErrToken.Code)

	e.GET("/douyin/user").WithQueryObject(define.UserInfoReq{Token: token, UserId: uint(userId) + 1}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.ErrToken.Code)

	e.GET("/douyin/user").WithQueryObject(define.UserInfoReq{Token: token, UserId: uint(userId)}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.OK.Code).
		Value("user").Object().ValueEqual("name", username)
}
