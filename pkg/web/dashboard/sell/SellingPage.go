package sell

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/pkg/product_management"
	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/selling"
)

func SellingPage(c *gin.Context) {

	sellingPageProps := selling.SellingPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	sellingPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	sellingPageFormProp := components.FormLayoutSimpleProp{Action: "/dashboard/selling", Method: "POST"}
	sellingPageProps.FormLayoutSimpleProp = sellingPageFormProp

	products, _ := product_management.GetAllProductsService()
	sellingPageProps.Products = products

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, selling.SellingPage(sellingPageProps))
	c.Render(http.StatusOK, r)
}
