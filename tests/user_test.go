package tests

import (
	"douyin/app/define"
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
	e.POST("/douyin/user/login").WithQueryObject(define.LoginReq{Username: "Pluto", Password: "123abC"}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.ErrUserNotFound.Code)

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
