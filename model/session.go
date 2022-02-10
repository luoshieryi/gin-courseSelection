package model

import (
	"time"
)

type Session struct {
	ID        int64     `gorm:"auto_increment"`
	Auth      string    `gorm:""`
	UserID    int64     `gorm:""`
	Expires   time.Time `gorm:""`
	CreatedAt time.Time `gorm:""` // gorm 自动更新字段
}
