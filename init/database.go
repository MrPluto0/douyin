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

func newLogger(logPath string) logger.Interface {
	// TODO: close writer
	writer, err := utils.OpenFile_A(logPath)
	utils.CheckPanicErr(err)
	return logger.New(
		log.New(writer, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)
}

func InitMysql() {
	// Read Config
	config := viper.Sub("database")
	host := config.Get("host")
	port := config.Get("port")
	user := config.Get("username")
	pwd := config.Get("password")
	dbname := config.Get("dbname")

	// Create database instance
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, host, port, dbname)
	logPath := viper.GetString("root") + config.GetString("log_path")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger(logPath),
	})
	utils.CheckPanicErr(err)

	// Initialize sql config
	sqlDB, err := db.DB()
	utils.CheckPanicErr(err)
	sqlDB.SetMaxIdleConns(config.GetInt("max_idle"))
	sqlDB.SetMaxOpenConns(config.GetInt("max_open"))

	// Migrate Models
	db.AutoMigrate(&models.User{})

	models.DB = db
}
