package aboutus

import (
	"nextgen/internals/db/pg"
)

func GetTeamInfo() (*[]TeamInfo, error) {
	var team []TeamInfo
	
	if err := pg.GDB.Find(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

