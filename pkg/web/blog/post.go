package blog

import (
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/app/appcomponents"
	"nextgen/templates/blog"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostPage(c *gin.Context) {

	postPageProps := blog.PostPageProps{}

	layoutProps, exists := c.Get("LayoutProp")

	if !exists {
		layoutProps = appcomponents.LayoutProps{}
	}

	postPageProps.LayoutProps = layoutProps.(appcomponents.LayoutProps)

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
	postPageProps.BlogPagePostProps = []blog.BlogPagePostProps{postProps}

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, blog.PostPage(postPageProps))
	c.Render(http.StatusOK, r)

}
