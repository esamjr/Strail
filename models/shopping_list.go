package models

import "github.com/jinzhu/gorm"

// Shopping List Models
type ShoppingList struct {
	gorm.Model `json:"model"`
	Name       string `json:"name" validate:"required"`
	Total      int    `json:"total" validate:"required"`
}
