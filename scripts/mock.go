package main

import (
	"douyin/app/models"
	Init "douyin/init"
	"douyin/utils/log"

	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	fmt.Println("----------Read Config...------------")

	Init.InitConfig()
	Init.InitMysql()
	// reset logger
	db = models.DB
	// db.Logger = db.Logger.LogMode(logger.Silent)

	fmt.Println("----------Start Mock...-------------")

	mockUsers()
	mockVideos()
}

func create[T any](model T) {
	err := db.Create(&model).Error

	if err != nil {
		log.Error(err)
	} else {
		log.Info(fmt.Sprintf("create success %+v", model))
	}
}

func mockUsers() {
	u := models.User{
		Name:     "Gypsophlia",
		Password: "123abC",
	}
	create(u)
}

func mockVideos() {
	v := models.Video{
		UserId:        1,
		Title:         "123",
		PlayUrl:       "123",
		CoverUrl:      "123",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}

	create(v)
}
