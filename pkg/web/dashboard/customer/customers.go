package customer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/customer"
	"nextgen/templates/components"
	"sort"
)

var CustomerPageHeaderProps = map[string]*customer.CustomerPageHeaderProp{
	"CustomerPage": &customer.CustomerPageHeaderProp{
		PositionNumber: 1,
		Label:          "Customers",
		Url:            "/dashboard/customers",
		Class:          ""},
	"AddCustomerPage": &customer.CustomerPageHeaderProp{
		PositionNumber: 2,
		Label:          "Add Customer",
		Url:            "/dashboard/add-customer",
		Class:          ""},
}
var headerElementsClass = "inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
var currentHeaderElementClass = "inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg active dark:text-gray-200 dark:border-gray-200"

func CustomerPage(c *gin.Context) {
	customerPageProps := customer.CustomerPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	customerPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	customerPageHeaderProps := make([]customer.CustomerPageHeaderProp, 0)
	for key, value := range CustomerPageHeaderProps {
		if key == "CustomerPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		customerPageHeaderProps = append(customerPageHeaderProps, *value)
	}
	sort.Slice(customerPageHeaderProps, func(i, j int) bool {
		return customerPageHeaderProps[i].PositionNumber < customerPageHeaderProps[j].PositionNumber
	})
	customerPageProps.CustomerPageHeaderProp = customerPageHeaderProps

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, customer.CustomerPage(customerPageProps))
	c.Render(http.StatusOK, r)

}

func AddCustomerPage(c *gin.Context) {
	addCustomerPageProps := customer.AddCustomerPageProps{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	addCustomerPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	customerPageHeaderProps := make([]customer.CustomerPageHeaderProp, 0)
	for key, value := range CustomerPageHeaderProps {
		if key == "AddCustomerPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		customerPageHeaderProps = append(customerPageHeaderProps, *value)
	}
	sort.Slice(customerPageHeaderProps, func(i, j int) bool {
		return customerPageHeaderProps[i].PositionNumber < customerPageHeaderProps[j].PositionNumber
	})
	addCustomerPageProps.CustomerPageHeaderProps = customerPageHeaderProps

	addCustomerFormProp := components.FormLayoutSimpleProp{Action: "/dashboard/add-customer", Method: "POST"}
	addCustomerPageProps.FormLayoutSimpleProp = addCustomerFormProp

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, customer.AddCustomerPage(addCustomerPageProps))
	c.Render(http.StatusOK, r)

}
