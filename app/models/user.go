package models

import (
	"fmt"
	"sync"
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
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (uD *UserDao) QueryByName(name string) (u User, err error) {
	userKey := "username-" + name
	err = ReadRedis(userKey, &u, func(value *User) error {
		return DB.Where("name = ?", name).First(&value).Error // if user not found, err occurs
	})
	return u, err
}

func (uD *UserDao) QueryById(id uint) (u User, err error) {
	userKey := fmt.Sprintf("userid-%d", id)
	err = ReadRedis(userKey, &u, func(value *User) error {
		return DB.First(&value, id).Error
	})
	return u, err
}

// This function is for benchmark
func (uD *UserDao) QueryByPwd(pwd string) (u User, err error) {
	err = DB.Where("password = ?", pwd).First(&u).Error
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
