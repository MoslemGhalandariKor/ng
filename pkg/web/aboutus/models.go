package aboutus

import "gorm.io/gorm"

type TeamInfo struct {
	gorm.Model
	Name       string
	Role       string
	Task       string
	Picture    string
	PictureAlt string
	Linkdin    string
	Github     string
	Facebook   string
	X          string
}

func (TeamInfo) TableName() string {
	return "tbl_team_info"
}
