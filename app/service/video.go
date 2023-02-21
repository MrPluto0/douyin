package service

import (
	"douyin/app/define"
	"douyin/app/models"
	"douyin/utils/response"
	"time"
)

type videoService struct{}

var VideoService *videoService

const FMT = "2006-01-02 15:04:05"

func init() {
	VideoService = &videoService{}
}

func (v *videoService) Feed(req define.FeedReq) define.FeedRes {
	var createdAt string
	if req.LatestTime == nil {
		createdAt = time.Now().Format(FMT)
	} else {
		createdAt = time.Unix(*req.LatestTime, 0).Format(FMT)
	}

	dao := models.NewVideoDaoInstance()
	videos, err := dao.QueryByCreateTime(createdAt)
	if err != nil {
		return define.FeedRes{
			Errno: response.ErrDatabase.Extend(err),
		}
	}

	var nextTime int64
	if len(videos) == 0 {
		nextTime = time.Now().Unix()
	} else {
		nextTime = videos[len(videos)-1].CreatedAt.Unix()
	}

	return define.FeedRes{
		Errno:     *response.OK,
		VideoList: videos,
		NextTime:  nextTime,
	}
}
