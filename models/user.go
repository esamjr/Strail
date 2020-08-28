package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"model"`
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
}
