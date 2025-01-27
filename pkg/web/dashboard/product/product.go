package product

import (
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/product"

	"github.com/gin-gonic/gin"
)

func ProductPage(c *gin.Context) {

	productPageprops := product.ProductPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	productPageprops.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.ProductPage(productPageprops))
	c.Render(http.StatusOK, r)
}
