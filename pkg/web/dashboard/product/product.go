package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/pkg/product_management"
	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/product"
	"nextgen/templates/dashboard/pages/product/productcomponents"
	"nextgen/templates/dashboard/pages/selling"
	"sort"
)

var ProductPageHeaderProps = map[string]*product.ProductPageHeaderProp{
	"ProductsPage": &product.ProductPageHeaderProp{
		PositionNumber: 1,
		Label:          "Products",
		Url:            "/dashboard/products",
		Class:          ""},
	"AddProductPage": &product.ProductPageHeaderProp{
		PositionNumber: 2,
		Label:          "Add Product",
		Url:            "/dashboard/add-product",
		Class:          ""},
	"CategoryPage": &product.ProductPageHeaderProp{
		PositionNumber: 3,
		Label:          "Categories",
		Url:            "/dashboard/categories",
		Class:          ""},
	"AddCategoryPage": &product.ProductPageHeaderProp{
		PositionNumber: 4,
		Label:          "Add Category",
		Url:            "/dashboard/add-category",
		Class:          ""},
	"BrandPage": &product.ProductPageHeaderProp{
		PositionNumber: 5,
		Label:          "Brands",
		Url:            "/dashboard/brands",
		Class:          ""},
	"AddBrandPage": &product.ProductPageHeaderProp{
		PositionNumber: 6,
		Label:          "Add Brand",
		Url:            "/dashboard/add-brand",
		Class:          ""},
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
		return productPageHeaderProps[i].PositionNumber < productPageHeaderProps[j].PositionNumber
	})
	productPageProps.ProductPageHeaderProps = productPageHeaderProps

	products, _ := product_management.GetAllProductsService()
	productPageProps.Products = products
	for _, prod := range productPageProps.Products {
		fmt.Println(prod)
	}

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
		return productPageHeaderProps[i].PositionNumber < productPageHeaderProps[j].PositionNumber
	})

	addProductPageProps.AddProductPageContentsProps.ProductPageHeaderProps = productPageHeaderProps

	addProductFormProp := components.FormLayoutSimpleProp{Action: "/dashboard/add-product", Method: "POST"}
	addProductPageProps.AddProductPageContentsProps.FormLayoutSimpleProp = addProductFormProp

	categories, _ := product_management.GetAllCategoriesService()
	addProductPageProps.AddProductPageContentsProps.CategoryInfo = categories

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, product.AddProductPage(addProductPageProps))
	c.Render(http.StatusOK, r)

}

func ProductCategorySearch(c *gin.Context) {
	q := c.Query("input-category-pattern-search")
	fmt.Println(q)
	var (
		categories []product_management.CategoryView
		err        error
	)

	if len(q) == 0 {
		categories, err = product_management.GetAllCategories()
	} else {
		categories, err = product_management.GetCategoriesByPatternService(q)
	}
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, productcomponents.ProductCategories(categories))
	c.Render(http.StatusOK, r)
}
func GetProductByNameHandler(c *gin.Context) {
	productName := c.Query("product-name")
	fmt.Println(productName)
	var (
		products []product_management.ProductView
	)
	products, _ = product_management.GetProductByNameService(productName)
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, selling.ProductRow(products))
	c.Render(http.StatusOK, r)

}
