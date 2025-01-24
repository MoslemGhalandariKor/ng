package app

import (
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/internals/helpers"
	"nextgen/templates/app/appcomponents"
	"nextgen/templates/app/pages"
	"nextgen/templates/components"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {

	posts11 := pages.IntroLandingprops{
		Title:   "Your life's work,powered by our life's work",
		Text:    "A unique and powerful software suite to transform the way you work. Designed for businesses of all sizes, built by a company that values your privacy.",
		Link:    "#",
		Picture: "../static/images/zh-homev2-banner.webp",
	}

	posts21 := pages.FeaturedAppsprops{
		Title: "CRM",
		Intro: "Comprehensive CRM platform for customer-facing teams.",
		Link:  "#",
	}
	posts22 := pages.FeaturedAppsprops{
		Title: "CRM",
		Intro: "Comprehensive CRM platform for customer-facing teams.",
		Link:  "#",
	}
	posts23 := pages.FeaturedAppsprops{
		Title: "CRM",
		Intro: "Comprehensive CRM platform for customer-facing teams.",
		Link:  "#",
	}
	posts24 := pages.FeaturedAppsprops{
		Title: "CRM",
		Intro: "Comprehensive CRM platform for customer-facing teams.",
		Link:  "#",
	}
	posts25 := pages.FeaturedAppsprops{
		Title: "CRM",
		Intro: "Comprehensive CRM platform for customer-facing teams.",
		Link:  "#",
	}

	posts1 := []pages.IntroLandingprops{posts11}
	posts2 := []pages.FeaturedAppsprops{posts21, posts22, posts23, posts24, posts25}
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
	pages.HomePage(userInfoProps, posts1, posts2, alerts).Render(c.Request.Context(), c.Writer)
}

func LoginPage(c *gin.Context) {
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
	pages.LoginPage(userInfoProps, alerts).Render(c.Request.Context(), c.Writer)
}

func RegisterPage(c *gin.Context) {
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
	pages.RegisterPage(userInfoProps, alerts).Render(c.Request.Context(), c.Writer)
}

func PageNotFound(c *gin.Context) {
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.PageNotFound())
	c.Render(http.StatusOK, r)
}
