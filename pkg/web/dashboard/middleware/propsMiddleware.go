package middleware

import (
	"github.com/gin-gonic/gin"
	"nextgen/templates/dashboard/dashboardcomponents"
	"github.com/gin-contrib/sessions"
	"nextgen/templates/components"
	"nextgen/internals/helpers"

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
	layoutProp.UserInfoProps = userInfoProps

	alerts, err := helpers.GetAlerts(c)
	if err != nil {
		alerts = []components.AlertProps{}
	}
	layoutProp.Alerts = alerts

	c.Set("LayoutProp", layoutProp)

}
