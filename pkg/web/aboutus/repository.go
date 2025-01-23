package aboutus

import (
	"nextgen/internals/db/pg"
)

func GetTeamInfo() (*[]TeamInfo, error) {
	var posts []TeamInfo
	
	if err := pg.GDB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

