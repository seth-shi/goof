package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`

	// add other fields
	gorm.Model

	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt *time.Time
}

func (User) TableName() string {

	return "users"
}
