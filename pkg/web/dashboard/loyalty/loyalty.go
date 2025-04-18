package loyalty

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/components"
	"nextgen/templates/dashboard/pages/loyalty"
	"sort"
)

var LoyaltyPageHeaderProps = map[string]*loyalty.LoyaltyPageHeaderProp{
	"LoyaltyPage": &loyalty.LoyaltyPageHeaderProp{
		PositionNumber: 1,
		Label:          "Loyalty Form",
		Url:            "/dashboard/loyalty",
		Class:          ""},
		"CampaignPage": &loyalty.LoyaltyPageHeaderProp{
		PositionNumber: 1,
		Label:          "Campaign Page",
		Url:            "/dashboard/campaign",
		Class:          ""},
}
var headerElementsClass = "inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
var currentHeaderElementClass = "inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg active dark:text-gray-200 dark:border-gray-200"

func LoyaltyPage(c *gin.Context) {

	loyaltyPageProps := loyalty.LoyaltyPageProps{}

	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	loyaltyPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	loyaltyPageHeaderProps := make([]loyalty.LoyaltyPageHeaderProp, 0)
	for key, value := range LoyaltyPageHeaderProps {
		if key == "LoyaltyPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		loyaltyPageHeaderProps = append(loyaltyPageHeaderProps, *value)
	}
	sort.Slice(loyaltyPageHeaderProps, func(i, j int) bool {
		return loyaltyPageHeaderProps[i].PositionNumber < loyaltyPageHeaderProps[j].PositionNumber
	})
	loyaltyPageProps.LoyaltyPageHeaderProps = loyaltyPageHeaderProps

	loyaltyFormProp := components.FormLayoutSimpleProp{Action: "/dashboard/add-brand", Method: "POST"}
	loyaltyPageProps.FormLayoutSimpleProp = loyaltyFormProp

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, loyalty.LoyaltyPage(loyaltyPageProps))
	c.Render(http.StatusOK, r)
}
