package accounts

import (
	"nextgen/internals/db/pg"
)

func CreateMembership(membership *Memberships) (*Memberships, error) {
	if err := pg.GDB.Create(membership).Error; err != nil {
		return nil, err
	}
	return membership, nil
}


