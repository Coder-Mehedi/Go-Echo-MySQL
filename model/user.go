package model

type User struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
}