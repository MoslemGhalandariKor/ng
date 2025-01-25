package app

import (
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/app/appcomponents"
	"nextgen/templates/app/pages"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	homePageProps := pages.HomePageProps{}
	layoutProps, exists := c.Get("LayoutProp")

	if !exists {
		layoutProps = appcomponents.LayoutProps{}
	}

	homePageProps.LayoutProps = layoutProps.(appcomponents.LayoutProps)

	pages.HomePage(homePageProps).Render(c.Request.Context(), c.Writer)
}

func LoginPage(c *gin.Context) {
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = appcomponents.LayoutProps{}
	}
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.LoginPage(layoutProp.(appcomponents.LayoutProps)))
	c.Render(http.StatusOK, r)
}

func RegisterPage(c *gin.Context) {
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = appcomponents.LayoutProps{}
	}
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.RegisterPage(layoutProp.(appcomponents.LayoutProps)))
	c.Render(http.StatusOK, r)
}

func PageNotFound(c *gin.Context) {
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.PageNotFound())
	c.Render(http.StatusOK, r)
}
