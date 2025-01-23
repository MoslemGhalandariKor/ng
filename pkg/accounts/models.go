package accounts

import "gorm.io/gorm"

type Memberships struct {
	gorm.Model
	Email string `gorm:"unique"`
}

func (Memberships) TableName() string {
	return "tbl_memberships"
}

