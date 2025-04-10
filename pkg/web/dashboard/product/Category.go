package product

import (
	"fmt"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/pkg/product_management"
	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/product"
	"sort"

	"github.com/gin-gonic/gin"
)

func CategoryPage(c *gin.Context) {
	categoryPageProps := product.CategoryPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	categoryPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	productPageHeaderProps := make([]product.ProductPageHeaderProp, 0)
	for key, value := range ProductPageHeaderProps {
		if key == "CategoryPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		productPageHeaderProps = append(productPageHeaderProps, *value)
	}
	sort.Slice(productPageHeaderProps, func(i, j int) bool {
		return productPageHeaderProps[i].PositionNumber < productPageHeaderProps[j].PositionNumber
	})
	categoryPageProps.ProductPageHeaderProps = productPageHeaderProps

	categories, err := product_management.GetAllCategoriesService()
	if err != nil {
		fmt.Println(err)
	}

	categoryPageProps.Categories = categories

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.CategoryPage(categoryPageProps))
	c.Render(http.StatusOK, r)

}

func AddCategoryPage(c *gin.Context) {
	addCategoryPageProps := product.AddCategoryPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	addCategoryPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	addCategoryPageHeaderProps := make([]product.ProductPageHeaderProp, 0)
	for key, value := range ProductPageHeaderProps {
		if key == "AddCategoryPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		addCategoryPageHeaderProps = append(addCategoryPageHeaderProps, *value)
	}
	sort.Slice(addCategoryPageHeaderProps, func(i, j int) bool {
		return addCategoryPageHeaderProps[i].PositionNumber < addCategoryPageHeaderProps[j].PositionNumber
	})
	addCategoryPageProps.ProductPageHeaderProps = addCategoryPageHeaderProps

	addCategoryFormProp := components.FormLayoutSimpleProp{Action: "/dashboard/add-category", Method: "POST",}
	addCategoryPageProps.FormLayoutSimpleProp = addCategoryFormProp

	
	categories, _ := product_management.GetAllCategoriesService()
	addCategoryPageProps.Categories = categories
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.AddCategoryPage(addCategoryPageProps))
	c.Render(http.StatusOK, r)

}
