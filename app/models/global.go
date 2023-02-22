package models

import (
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CommonModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

const EmptyID = 0

var DB *gorm.DB

var Redis *redis.Client
