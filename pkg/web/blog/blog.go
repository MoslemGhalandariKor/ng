package blog

import (
	"fmt"
	"nextgen/internals/helpers"
	"nextgen/templates/app/appcomponents"
	"nextgen/templates/blog"
	"nextgen/templates/components"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BlogPage(c *gin.Context) {

	posts, err := GetAllPosts()
	if err != nil {
		fmt.Println(err)
	}

	postsProps := []blog.BlogPagePostProps{}

	for _, post := range *posts {
		postsProps = append(postsProps, blog.BlogPagePostProps{
			Category:      post.Category,
			Duration_time: post.Duration_time,
			Title:         post.Title,
			Summary:       post.Summary,
			Writer:        post.Writer,
			WriterRole:    post.WriterRole,
			Image:   	   post.Image,
			ImageAlt:      post.ImageAlt,
			WriterPicture: post.WriterPicture,
			PictureAlt:    post.PictureAlt,
			Link:          post.Link,
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
	blog.BlogPage(userInfoProps, postsProps, alerts).Render(c.Request.Context(), c.Writer)
}
