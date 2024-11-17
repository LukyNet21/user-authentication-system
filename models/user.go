package models

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `json:"username" gorm:"not null;unique; min:3; max:30"`
	Password  string `json:"password" gorm:"not null; min:8"`
	Email     string `json:"email" gorm:"not null;unique; email"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
}
