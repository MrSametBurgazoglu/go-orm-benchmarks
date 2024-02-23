package gormmodels

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey"`
	Title    string    `gorm:"type:varchar(255);not null"`
	Content  string    `gorm:"type:text;not null"`
	UserID   uint      `gorm:"not null"` // Foreign key
	User     User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments []Comment // One-to-many relationship
}
