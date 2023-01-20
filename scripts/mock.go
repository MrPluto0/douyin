package main

import (
	"douyin/app/models"
	Init "douyin/init"
	"douyin/utils/log"

	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func main() {
	fmt.Println("----------Read Config...-------------")

	Init.InitConfig()
	Init.InitMysql()
	// reset logger
	db = models.DB
	db.Logger = db.Logger.LogMode(logger.Silent)

	fmt.Println("----------Start Mock...-------------")

	mockUsers()
}

func mockUsers() {
	u := models.User{
		CommonModel: models.CommonModel{ID: 1},
		Name:        "Gypsophlia",
		Password:    "123456",
	}
	err := db.Create(&u).Error

	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info(fmt.Sprintf("create success %+v", u))
	}
}
