package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAT time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `gorm:"not null"`
	Content   string         `gorm:"not null"`
	UserID    uint
	User      User      `gorm:"foreignkey:UserID"`
	Comments  []Comment `gorm:"foreignkey:PostID"`
}
