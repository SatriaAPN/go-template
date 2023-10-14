package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Password  string         `json:"-" gorm:"not null"`
}
