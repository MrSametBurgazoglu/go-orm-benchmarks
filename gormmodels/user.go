package gormmodels

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(100);not null;unique"`
	Posts []Post
}
