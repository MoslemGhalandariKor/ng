package blog

import (
	
	"nextgen/internals/helpers"
	"nextgen/templates/app/appcomponents"
	"nextgen/templates/blog"
	"nextgen/templates/components"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostPage(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(400, "Invalid ID")
		return
	}

	post, err := GetPostById(idInt)
	if err != nil {
		c.Redirect(404, "Post not found")
		return
	}
	postProps := blog.BlogPagePostProps{
		Category:      post.Category,
		Duration_time: post.Duration_time,
		Title:         post.Title,
		Summary:       post.Summary,
		Content:       post.Content,
		Writer:        post.Writer,
		WriterRole:    post.WriterRole,
		WriterPicture: post.WriterPicture,
		PictureAlt:    post.PictureAlt,
		Image:         post.Image,
		ImageAlt:      post.ImageAlt,
		Link:          post.Link,
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
	blog.PostPage(userInfoProps, postProps, alerts).Render(c.Request.Context(), c.Writer)
}
