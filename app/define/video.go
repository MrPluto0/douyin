package define

import (
	"douyin/app/models"
	"douyin/utils/response"
)

type FeedReq struct {
	Token      string `form:"token"  url:"token"`
	LatestTime int64  `form:"latest_time"  url:"latest_time"`
}

type FeedRes struct {
	response.Errno
	NextTime  int64          `json:"next_time,omitempty"` // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []models.Video `json:"video_list"`          // 视频列表
}
