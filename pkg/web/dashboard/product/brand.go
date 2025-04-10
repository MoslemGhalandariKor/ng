package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/product"
	"sort"
	"nextgen/templates/components"
	"nextgen/pkg/product_management"
)

func BrandPage(c *gin.Context) {
	brandPageProps := product.BrandPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	brandPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	productPageHeaderProps := make([]product.ProductPageHeaderProp, 0)
	for key, value := range ProductPageHeaderProps {
		if key == "BrandPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		productPageHeaderProps = append(productPageHeaderProps, *value)
	}
	sort.Slice(productPageHeaderProps, func(i, j int) bool {
		return productPageHeaderProps[i].PositionNumber < productPageHeaderProps[j].PositionNumber
	})
	brandPageProps.ProductPageHeaderProps = productPageHeaderProps

	brands, _ := product_management.GetAllBrandsService()
	brandPageProps.Brands = brands


	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.BrandPage(brandPageProps))
	c.Render(http.StatusOK, r)

}

func AddBrandPage(c *gin.Context) {
	addBrandPageProps := product.AddBrandPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	addBrandPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	productPageHeaderProps := make([]product.ProductPageHeaderProp, 0)
	for key, value := range ProductPageHeaderProps {
		if key == "AddBrandPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		productPageHeaderProps = append(productPageHeaderProps, *value)
	}
	sort.Slice(productPageHeaderProps, func(i, j int) bool {
		return productPageHeaderProps[i].PositionNumber < productPageHeaderProps[j].PositionNumber
	})
	addBrandPageProps.ProductPageHeaderProps = productPageHeaderProps

	addCategoryFormProp := components.FormLayoutSimpleProp{Action: "/dashboard/add-brand", Method: "POST"}
	addBrandPageProps.FormLayoutSimpleProp = addCategoryFormProp

	
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.AddBrandPage(addBrandPageProps))
	c.Render(http.StatusOK, r)

}
