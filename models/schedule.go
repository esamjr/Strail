package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// type JSONTime time.Time

// func (t JSONTime) MarshalJSON() ([]byte, error) {
// 	//do your serializing here
// 	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
// 	return []byte(stamp), nil
// }

// Schedule Model
type Schedule struct {
	gorm.Model `json:"model"`
	// Username   string    `json:"username" gorm:"foreignKey:username"`
	Title string    `json:"title" validate:"required"`
	Time  time.Time `json:"time" validate:"required"`
}
