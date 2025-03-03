package product

import (
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/pkg/web/dashboard/product"
	"nextgen/templates/dashboard/dashboardcomponents"
	"sort"

	"github.com/gin-gonic/gin"
)

func Category(c *gin.Context) {
	categoryPagePros := product.CategoryPagePros{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	categoryPagePros.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	productPageHeaderProps := make([]product.ProductPageHeaderProp, 0)
	for key, value := range ProductPageHeaderProps {
		if key == "ProductsPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		productPageHeaderProps = append(productPageHeaderProps, *value)
	}
	sort.Slice(productPageHeaderProps, func(i, j int) bool {
		return productPageHeaderProps[i].Url > productPageHeaderProps[j].Url
	})

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.CategoryPage(categoryPagePros))
	c.Render(http.StatusOK, r)

}
