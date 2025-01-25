package blog

import (
	"fmt"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/app/appcomponents"
	"nextgen/templates/blog"

	"github.com/gin-gonic/gin"
)

func BlogPage(c *gin.Context) {

	blogPageProps := blog.BlogPageProps{}
	layoutProps, exists := c.Get("LayoutProp")

	if !exists {
		layoutProps = appcomponents.LayoutProps{}
	}

	blogPageProps.LayoutProps = layoutProps.(appcomponents.LayoutProps)

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
			Image:         post.Image,
			ImageAlt:      post.ImageAlt,
			WriterPicture: post.WriterPicture,
			PictureAlt:    post.PictureAlt,
			Link:          post.Link,
		})
	}

	blogPageProps.BlogPagePostProps = postsProps

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, blog.BlogPage(blogPageProps))
	c.Render(http.StatusOK, r)

}
