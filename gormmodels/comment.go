package gormmodels

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey"`
	Text   string `gorm:"type:varchar(255);not null"`
	PostID uint   `gorm:"not null"` // Foreign key
	Post   Post   `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
