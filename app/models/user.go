package models

import (
	"context"
	"encoding/json"
	"sync"
	"time"
)

// User Table = Gorm Model
type User struct {
	CommonModel
	Name            string `json:"name" gorm:"type:varchar(32);not null;uniqueIndex"`
	Password        string `json:"-" gorm:"type:varchar(32);not null"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       string `json:"work_count"`
	FavoriteCount   string `json:"favorite_count"`
}

// User Dao
type UserDao struct{}

var (
	userDao  *UserDao
	userOnce sync.Once
	ctx      = context.Background()
	expire   = time.Hour * 24
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (uD *UserDao) QueryByName(name string) (u User, err error) {
	userKey := "user-" + name
	userStr, err := Redis.Get(ctx, userKey).Result()

	if err != nil {
		err = DB.Where("name = ?", name).First(&u).Error // if user not found, err occurs
		if err == nil {
			userByte, _ := json.Marshal(u)
			err = Redis.Set(ctx, userKey, string(userByte), expire).Err()
		}
	} else {
		err = json.Unmarshal([]byte(userStr), &u)
	}

	return u, err
}

// This function is for benchmark
func (uD *UserDao) QueryByPwd(pwd string) (u User, err error) {
	err = DB.Where("password = ?", pwd).First(&u).Error
	return u, err
}

func (uD *UserDao) QueryById(id uint) (u User, err error) {
	err = DB.First(&u, id).Error
	return u, err
}

func (uD *UserDao) Create(name string, pwd string) (rowsAffected int64, err error) {
	user := User{Name: name, Password: pwd}
	result := DB.Create(&user)
	return result.RowsAffected, result.Error
}

func (uD *UserDao) Delete(name string) error {
	return DB.Where("name = ?", name).Delete(&User{}).Error
}
