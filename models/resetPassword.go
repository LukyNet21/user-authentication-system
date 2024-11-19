package models

import "time"

type ResetPassword struct {
	ID      uint      `gorm:"primary_key"`
	UserId  uint      `json:"user_id" gorm:"not null"`
	Token   string    `json:"token" gorm:"not null"`
	ValidTo time.Time `json:"valid_to" gorm:"not null"`
}
