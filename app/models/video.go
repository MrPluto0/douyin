package models

import (
	"sync"
)

// Table = Gorm Model
type Video struct {
	CommonModel
	User          User   `json:"auther" gorm:"foreignKey:UserId"`
	UserId        int    `json:"user_id"`
	Title         string `json:"title"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

// User Dao
type VideoDao struct{}

const VIDEO_LIMIT = 30

var (
	video     *VideoDao
	videoOnce sync.Once
)

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(func() {
		video = &VideoDao{}
	})
	return video
}

func (vD *VideoDao) QueryByCreateTime(createdAt string) (videos []Video, err error) {
	err = DB.Order("created_at desc").Limit(VIDEO_LIMIT).Where("created_at <= ?", createdAt).Find(&videos).Error
	return videos, err
}
