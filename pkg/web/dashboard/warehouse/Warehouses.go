package warehouse

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/pkg/product_management"
	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/warehouse"
	"sort"
	"fmt"
)

var WarehouseHeaderPageProps = map[string]*warehouse.WarehouseHeaderPageProps{
	"WarehousesPage": &warehouse.WarehouseHeaderPageProps{
		Label: "Warehouses",
		Url:   "/dashboard/warehouses",
		Class: ""},
	"AddWarehousePage": &warehouse.WarehouseHeaderPageProps{
		Label: "Add Warehouse",
		Url:   "/dashboard/add-warehouse",
		Class: ""},
	"InventoryPage": &warehouse.WarehouseHeaderPageProps{
		Label: "Inventory",
		Url:   "/dashboard/inventory",
		Class: ""},
	// "UpdateProductPage": &warehouse.WarehouseHeaderPageProps{
	// 	Label: "Update Product",
	// 	Url:   "/dashboard/update-product",
	// 	Class: ""},
}
var headerElementsClass = "inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
var currentHeaderElementClass = "inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg active dark:text-gray-200 dark:border-gray-200"

func WarehousesPage(c *gin.Context) {

	warehousesPageProps := warehouse.WarehousesPageProps{}

	layoutProp, exists := c.Get("LayoutProp")
	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}
	warehousesPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	warehouseHeaderPageProps := make([]warehouse.WarehouseHeaderPageProps, 0)
	for key, value := range WarehouseHeaderPageProps {
		if key == "WarehousesPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		warehouseHeaderPageProps = append(warehouseHeaderPageProps, *value)
	}
	sort.Slice(warehouseHeaderPageProps, func(i, j int) bool {
		return warehouseHeaderPageProps[i].Url > warehouseHeaderPageProps[j].Url
	})
	warehousesPageProps.WarehouseHeaderPageProps = warehouseHeaderPageProps

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, warehouse.WarehpusesPage(warehousesPageProps))
	c.Render(http.StatusOK, r)
}

func AddWarehousePage(c *gin.Context) {

	addWarehousePageProps := warehouse.AddWarehousePageProps{}

	layoutProp, exists := c.Get("LayoutProp")
	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}
	addWarehousePageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	warehouseHeaderPageProps := make([]warehouse.WarehouseHeaderPageProps, 0)
	for key, value := range WarehouseHeaderPageProps {
		if key == "AddWarehousePage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		warehouseHeaderPageProps = append(warehouseHeaderPageProps, *value)
	}
	sort.Slice(warehouseHeaderPageProps, func(i, j int) bool {
		return warehouseHeaderPageProps[i].Url > warehouseHeaderPageProps[j].Url
	})
	addWarehousePageProps.WarehouseHeaderPageProps = warehouseHeaderPageProps

	addWarehouseFormProp := components.FormLayoutSimpleProp{Action: "/dashboard/add-warehouse", Method: "POST"}
	addWarehousePageProps.FormLayoutSimpleProp = addWarehouseFormProp

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, warehouse.AddWarehousePage(addWarehousePageProps))
	c.Render(http.StatusOK, r)
}

func InventoryPage(c *gin.Context) {

	inventoryPageProps := warehouse.InventoryPageProps{}

	layoutProp, exists := c.Get("LayoutProp")
	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}
	inventoryPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	warehouseHeaderPageProps := make([]warehouse.WarehouseHeaderPageProps, 0)
	for key, value := range WarehouseHeaderPageProps {
		if key == "InventoryPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		warehouseHeaderPageProps = append(warehouseHeaderPageProps, *value)
	}
	sort.Slice(warehouseHeaderPageProps, func(i, j int) bool {
		return warehouseHeaderPageProps[i].Url > warehouseHeaderPageProps[j].Url
	})
	inventoryPageProps.WarehouseHeaderPageProps = warehouseHeaderPageProps

	inventoryFormProp := components.FormLayoutSimpleProp{Action: "/dashboard/add-warehouse", Method: "POST"}
	inventoryPageProps.FormLayoutSimpleProp = inventoryFormProp

	products, _ := product_management.GetAllProductsService()
	inventoryPageProps.Products = products

	for _, prod := range inventoryPageProps.Products {
		fmt.Println(prod)
	}

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, warehouse.InventoryPage(inventoryPageProps))
	c.Render(http.StatusOK, r)
}
