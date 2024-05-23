package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"ID,omitempty" gorm:"primary_key"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
