package models

import (
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

const EmptyID = 0

var DB *gorm.DB
