package product

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	// product_management "nextgen/pkg/product-management"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/product"
	"sort"
)

var ProductPageHeaderProps = map[string]*product.ProductPageHeaderProp{
	"ProductsPage": &product.ProductPageHeaderProp{
		Label: "Products",
		Url:   "/dashboard/products",
		Class: ""},
	"AddProductPage": &product.ProductPageHeaderProp{
		Label: "Add Product",
		Url:   "/dashboard/add-product",
		Class: ""},
}
var headerElementsClass = "inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
var currentHeaderElementClass = "inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg active dark:text-gray-200 dark:border-gray-200"

func ProductsPage(c *gin.Context) {

	productPageProps := product.ProductPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	productPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

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

	// productsList := []product.ProductListProps{}
	// products := []product_management.Product{}

	// products, err := product_management.GetAllProducts()

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, prod := range products {
	// 		productsList = append(productsList, product.ProductListProps{
	// 			ProductName: prod.Name,
	// 			Description: prod.Description,
	// 		})
	// 	}
	// }

	productPageProps.ProductPageHeaderProps = productPageHeaderProps
	// productPageProps.ProductListProps = productsList
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.ProductPage(productPageProps))
	c.Render(http.StatusOK, r)
}

func AddProductPage(c *gin.Context) {

	addProductPageProps := product.AddProductPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	addProductPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	productPageHeaderProps := make([]product.ProductPageHeaderProp, 0)
	for key, value := range ProductPageHeaderProps {
		if key == "AddProductPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		productPageHeaderProps = append(productPageHeaderProps, *value)
	}
	sort.Slice(productPageHeaderProps, func(i, j int) bool {
		return productPageHeaderProps[i].Url > productPageHeaderProps[j].Url
	})

	addProductPageProps.AddProductPageContentsProp.ProductPageHeaderProps = productPageHeaderProps

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.AddProductPage(addProductPageProps))
	c.Render(http.StatusOK, r)

}
