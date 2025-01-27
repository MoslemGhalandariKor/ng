package product

import (
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/product"

	"github.com/gin-gonic/gin"
)

func ShopPage(c *gin.Context) {

	shopPageProps := product.ShopPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	shopPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.ShopPage(shopPageProps))
	c.Render(http.StatusOK, r)
}
