package models

import (
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	ID        uint           `gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

var DB *gorm.DB
