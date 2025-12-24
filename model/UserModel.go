package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAT time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	Email     string         `gorm:"unique;not null"`
	SecretKey string         `gorm:"not null"`
	Posts     []Post         `gorm:"foreignkey:UserID"`
	Comments  []Comment      `gorm:"foreignkey:UserID"`
}
