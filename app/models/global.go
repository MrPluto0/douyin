package models

import (
	"context"
	"encoding/json"
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

var (
	ctx    = context.Background()
	expire = time.Hour * 24
)

func ReadRedis[T any](key string, value *T, fallback func(value *T) error) error {
	var err error
	var val string

	val, err = Redis.Get(ctx, key).Result()

	if err == redis.Nil {
		err = fallback(value)
		if err == nil {
			userByte, _ := json.Marshal(value)
			err = Redis.Set(ctx, key, string(userByte), expire).Err()
		}
	} else if err != nil {
		panic(err)
	} else {
		err = json.Unmarshal([]byte(val), value)
	}
	return err
}
