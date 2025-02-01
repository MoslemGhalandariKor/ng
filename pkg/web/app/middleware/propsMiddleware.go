package middleware

import (
	"nextgen/internals/helpers"
	"nextgen/templates/app/appcomponents"
	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LayoutPropMiddelware(c *gin.Context) {
	layoutProp := appcomponents.LayoutProps{}

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
	layoutProp.UserInfoProps = userInfoProps

	alerts, err := helpers.GetAlerts(c)
	if err != nil {
		alerts = []components.AlertProps{}
	}
	layoutProp.Alerts = alerts

	c.Set("LayoutProp", layoutProp)

}
