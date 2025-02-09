package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"default:user"`
	ProfileImage string `gorm:"default:default.jpg"`
	Bio          string `gorm:"default:null"`
}
