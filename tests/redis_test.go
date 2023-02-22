package tests

import (
	"context"
	"douyin/app/models"
	"testing"
)

func TestRedis(t *testing.T) {
	rds := models.Redis
	ctx := context.Background()
	val, err := rds.Get(ctx, "user-gypsophlia").Result()
	if err != nil {
		panic(err)
	}
	t.Fatal("key", val)
}

func BenchmarkRedis(b *testing.B) {
	rds := models.Redis
	ctx := context.Background()
	key := "user-Gypsophlia"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rds.Get(ctx, key).Result()
	}

}

func BenchmarkMysql(b *testing.B) {
	var u models.User
	db := models.DB
	name := "Gypsophlia"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.Where("name = ?", name).First(&u).Error
	}
}
