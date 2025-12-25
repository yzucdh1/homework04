package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAT time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string         `gorm:"not null" json:"title"`
	Content   string         `gorm:"not null" json:"content"`
	UserID    uint           `json:"user_id"`
	User      User           `gorm:"foreignkey:UserID" json:"user"`
	Comments  []Comment      `gorm:"foreignkey:PostID" json:"comments"`
}
