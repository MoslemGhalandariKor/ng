package dashboard

import (
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages"

	"github.com/gin-gonic/gin"
)

func DashboardPage(c *gin.Context) {

	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.DashboardPage(layoutProp.(dashboardcomponents.LayoutProp)))
	c.Render(http.StatusOK, r)

}

func ProductPage(c *gin.Context) {
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.ProductPage(layoutProp.(dashboardcomponents.LayoutProp)))
	c.Render(http.StatusOK, r)
}

func CalendarPage(c *gin.Context) {
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.CalendarPage(layoutProp.(dashboardcomponents.LayoutProp)))
	c.Render(http.StatusOK, r)
}

func LoyaltyPage(c *gin.Context) {
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.LoyaltyPage(layoutProp.(dashboardcomponents.LayoutProp)))
	c.Render(http.StatusOK, r)
}

func CampaignPage(c *gin.Context) {
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.CampaignPage(layoutProp.(dashboardcomponents.LayoutProp)))
	c.Render(http.StatusOK, r)
}

func SettingsPage(c *gin.Context) {
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.SettingsPage(layoutProp.(dashboardcomponents.LayoutProp)))
	c.Render(http.StatusOK, r)
}

