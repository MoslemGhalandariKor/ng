package middleware

import (
	"nextgen/internals/helpers"
	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LayoutPropMiddelware(c *gin.Context) {
	layoutProp := dashboardcomponents.LayoutProp{}

	userInfoProps := dashboardcomponents.UserInfoProps{}
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
	userInfoProps.CompanyLogo = "../static/images/zh-video-poster-paul (1).png"
	userInfoProps.CompanyName = "Your Company Name"
	layoutProp.UserInfoProps = userInfoProps

	alerts, err := helpers.GetAlerts(c)
	if err != nil {
		alerts = []components.AlertProps{}
	}
	layoutProp.Alerts = alerts

	c.Set("LayoutProp", layoutProp)

}
