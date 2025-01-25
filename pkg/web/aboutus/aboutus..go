package aboutus

import (
	"fmt"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/aboutus"
	"nextgen/templates/app/appcomponents"

	"github.com/gin-gonic/gin"
)

func TeamPage(c *gin.Context) {

	teamPageProps := aboutus.TeamPageProps{}
	layoutProps, exists := c.Get("LayoutProp")

	if !exists {
		layoutProps = appcomponents.LayoutProps{}
	}

	teamPageProps.LayoutProps = layoutProps.(appcomponents.LayoutProps)

	team, err := GetTeamInfo()
	if err != nil {
		fmt.Println(err)
	}

	teamInfo := []aboutus.TeamInfoProps{}

	for _, team := range *team {
		teamInfo = append(teamInfo, aboutus.TeamInfoProps{
			Name:       team.Name,
			Role:       team.Role,
			Task:       team.Task,
			Picture:    team.Picture,
			PictureAlt: team.PictureAlt,
			Linkdin:    team.Linkdin,
			Github:     team.Github,
			Facebook:   team.Facebook,
			X:          team.X,
		})
	}

	teamPageProps.TeamInfoProps = teamInfo

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, aboutus.TeamPage(teamPageProps))
	c.Render(http.StatusOK, r)
}
