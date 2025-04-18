package loyalty

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/loyalty"
	"sort"
)




func CampaignPage(c *gin.Context) {

	campaignPageProps := loyalty.CampaignPageProps{}

	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	campaignPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)


	loyaltyPageHeaderProps := make([]loyalty.LoyaltyPageHeaderProp, 0)
	for key, value := range LoyaltyPageHeaderProps {
		if key == "CampaignPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		loyaltyPageHeaderProps = append(loyaltyPageHeaderProps, *value)
	}
	sort.Slice(loyaltyPageHeaderProps, func(i, j int) bool {
		return loyaltyPageHeaderProps[i].PositionNumber < loyaltyPageHeaderProps[j].PositionNumber
	})
	campaignPageProps.LoyaltyPageHeaderProps = loyaltyPageHeaderProps


	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, loyalty.CampaignPage(campaignPageProps))
	c.Render(http.StatusOK, r)
}