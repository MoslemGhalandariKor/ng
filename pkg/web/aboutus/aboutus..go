package aboutus

import (
	"fmt"
	"nextgen/internals/helpers"
	"nextgen/templates/aboutus"
	"nextgen/templates/app/appcomponents"
	"nextgen/templates/components"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func TeamPage(c *gin.Context) {
	posts, err := GetTeamInfo()
	if err != nil {
		fmt.Println(err)
	}

	postsProps := []aboutus.TeamInfoProps{}

	for _, post := range *posts {
		postsProps = append(postsProps, aboutus.TeamInfoProps{
			Name:       post.Name,
			Role:       post.Role,
			Task:       post.Task,
			Picture:    post.Picture,
			PictureAlt: post.PictureAlt,
			Linkdin:    post.Linkdin,
			Github:     post.Github,
			Facebook:   post.Facebook,
			X:          post.X,
		})
	}

	var alerts []components.AlertProps
	var userInfoProps appcomponents.UserInfoProps

	session := sessions.Default(c)
	email := session.Get("Email")
	username := session.Get("Username")
	if email != nil {
		userInfoProps.IsLoggedIn = true
		userInfoProps.Email = email.(string)
		userInfoProps.Username = username.(string)
	} else {
		userInfoProps.IsLoggedIn = false
	}
	flashes, _ := helpers.GetFlash(c, "flash")

	for _, f := range flashes {
		alerts = append(alerts, components.AlertProps{Type: f.Type, Message: f.Message, Id: f.Id, DismissId: f.DismissId})
	}
	aboutus.TeamPage(userInfoProps, postsProps, alerts).Render(c.Request.Context(), c.Writer)
}
