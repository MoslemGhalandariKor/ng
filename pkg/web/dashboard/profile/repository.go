package profile

import (
	"nextgen/internals/db/pg"
	"nextgen/pkg/auth"
)

func GetUser() (*[]auth.User, error) {
	var users []auth.User

	if err := pg.GDB.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
