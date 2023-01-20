package init

import (
	"douyin/app/models"
	"douyin/utils"

	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql() {
	config := viper.Sub("database")
	host := config.Get("host")
	port := config.Get("port")
	user := config.Get("username")
	pwd := config.Get("password")
	dbname := config.Get("dbname")

	// Create logger
	writer, err := utils.OpenFile_A(viper.GetString("root") + config.GetString("log_path"))
	utils.CheckPanicErr(err)
	newLogger := logger.New(
		log.New(writer, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)

	// Create database instance
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	utils.CheckPanicErr(err)

	// initialize sql config
	sqlDB, err := db.DB()
	utils.CheckPanicErr(err)
	sqlDB.SetMaxIdleConns(config.GetInt("max_idle"))
	sqlDB.SetMaxOpenConns(config.GetInt("max_open"))

	models.DB = db
}
