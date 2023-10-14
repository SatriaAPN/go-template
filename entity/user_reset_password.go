package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserResetPassword struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserId    int            `gorm:"not null"`
	User      User
	Token     string    `gorm:"not null"`
	ExpiredAt time.Time `gorm:"not null"`
}
