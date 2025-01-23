package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string 
	LastName    string 
	Username    string
	Email       string `gorm:"index, unique"`
	PhoneNumber string `gorm:"index, unique"`
	Password    string
}

func (User) TableName() string {
	return "tbl_users"
}



