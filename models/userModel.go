package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_name string
	Last_name  string
	Email      string
	Username   string
	Password   string
	IsAdmin    string `gorm:"default:null"`
}

type Errors struct {
	Empty        string
	Notavailable string
}
