package product

import (
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/product"
	"nextgen/templates/components"
	"github.com/gin-gonic/gin"
	"nextgen/pkg/product_management"
)

func SellingPage(c *gin.Context) {

	sellingPageProps := product.SellingPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	sellingPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	sellingPageFormProp := components.FormLayoutSimpleProp{Action: "/dashboard/selling", Method: "POST"}
	sellingPageProps.FormLayoutSimpleProp = sellingPageFormProp

	products, _ := product_management.GetAllProductsService()
	sellingPageProps.Products = products

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.SellingPage(sellingPageProps))
	c.Render(http.StatusOK, r)
}
