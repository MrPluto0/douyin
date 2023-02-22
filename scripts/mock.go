package main

import (
	"douyin/app/models"
	Init "douyin/init"
	"douyin/utils/log"
	"math/rand"
	"sync"
	"time"

	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB
var userTotal = 10000

func main() {
	fmt.Println("----------Read Config...------------")

	Init.InitConfig()
	Init.InitMysql()
	db = models.DB

	fmt.Println("----------Start Mock...-------------")

	// mockUsers()
	mockVideos()
}

func create[T any](model T) {
	err := db.Create(&model).Error

	if err != nil {
		log.Error(err)
	} else {
		log.Info(fmt.Sprintf("create success %+v\n", model))
	}
}

func randStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func mockUsers() {
	var wg sync.WaitGroup
	wg.Add(userTotal)
	for i := 0; i < userTotal; i++ {
		go func() {
			create(models.User{
				Name:     randStr(6),
				Password: randStr(6),
			})
			wg.Done()
		}()
	}
	wg.Wait()
}

func mockVideos() {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		t, _ := time.ParseDuration(fmt.Sprintf("-%dm", i))
		go func() {
			create(models.Video{
				CommonModel: models.CommonModel{
					CreatedAt: time.Now().Add(t),
				},
				Title:         randStr(6),
				PlayUrl:       "/static/video/1.mp4",
				CoverUrl:      "/static/img/1_cover.png",
				FavoriteCount: rand.Intn(10000),
				CommentCount:  rand.Intn(1000),
				IsFavorite:    false,
				UserId:        rand.Intn(50),
			})
			wg.Done()
		}()
	}
	wg.Wait()
}
