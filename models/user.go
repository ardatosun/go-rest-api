package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
