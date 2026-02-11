package sqlite

import "gorm.io/gorm"

// Declaring User Model
type UserModel struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}
