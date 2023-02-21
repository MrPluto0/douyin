package tests

import (
	"douyin/app/define"
	"douyin/utils/response"
	"net/http"
	"testing"
	"time"
)

func TestFeed(t *testing.T) {
	e := NewHttpExcept(t)

	e.GET("/douyin/feed").
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.OK.Code).
		Value("video_list").Array().Length().Equal(30)

	e.GET("/douyin/feed").WithQueryObject(define.FeedReq{LatestTime: time.Now().Unix()}).
		Expect().Status(http.StatusOK).JSON().Object().
		ValueEqual("status_code", response.OK.Code).
		Value("video_list").Array().Length().Equal(30)
}
