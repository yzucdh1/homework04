package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAT time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Content   string         `gorm:"not null"`
	UserID    uint
	User      User `gorm:"foreignkey:UserID"`
	PostID    uint
	Post      Post `gorm:"foreignkey:PostID"`
}
