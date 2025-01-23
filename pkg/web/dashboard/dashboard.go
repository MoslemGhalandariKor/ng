package dashboard

import (
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"github.com/gin-gonic/gin"
)

func DashboardPage(c *gin.Context) {

	layoutProp, exists := c.Get("LayoutProp")
	
	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.DashboardPage(layoutProp.(dashboardcomponents.LayoutProp)))
	c.Render(http.StatusOK, r)

}

func ProductPage(c *gin.Context) {
	layoutProp, exists := c.Get("LayoutProp")
	
	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.ProductPage(layoutProp.(dashboardcomponents.LayoutProp)))
	c.Render(http.StatusOK, r)
}

func CalendarPage(c *gin.Context) {
	layoutProp, exists := c.Get("LayoutProp")
	
	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.CalendarPage(layoutProp.(dashboardcomponents.LayoutProp)))
	c.Render(http.StatusOK, r)
}

// func BusinessInfoPage(c *gin.Context) {

// 	formLayoutSimpleProp := components.FormLayoutSimpleProp{Action: "/accounts/add-company", Method: "Post"}
// 	CompanyNameInput := components.InputFloatingProps{Type: "text", Name: "company_name", Required: true, Label: "Company Name"}
// 	CompanyAddressInput := components.InputFloatingProps{Type: "text", Name: "company_address", Required: true, Label: "Company Address"}
// 	CompanyPhoneInput := components.InputFloatingProps{Type: "text", Name: "company_phone_number", Required: true, Label: "Company Phone Number"}
// 	CompanyEmailInput := components.InputFloatingProps{Type: "email", Name: "company_email", Required: true, Label: "Company Email"}

// 	layoutProp := dashboardcomponents.LayoutProp{}

// 	var userInfoProps dashboardcomponents.UserInfoProps
// 	session := sessions.Default(c)
// 	email := session.Get("Email")
// 	userName := session.Get("UserName")
// 	userInfoProps.Email = email.(string)
// 	if userName == nil {
// 		userInfoProps.Username = email.(string)
// 	} else {
// 		userInfoProps.Username = userName.(string)
// 	}

// 	alerts, err := helpers.GetAlerts(c)
// 	if err != nil {
// 		alerts = []components.AlertProps{}
// 	}
// 	layoutProp.Alerts = alerts

// 	inputFormProps := []components.InputFloatingProps{CompanyNameInput, CompanyAddressInput, CompanyEmailInput, CompanyPhoneInput}
	
// 	props := pages.DashboardProps{}
// 	props.Email = email.(string)
// 	if userName == nil {
// 		props.Name = email.(string)
// 	}

// 	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, pages.BusinessInfoPage(layoutProp, props, formLayoutSimpleProp, inputFormProps))
// 	c.Render(http.StatusOK, r)
// }

// func TeamPage(c *gin.Context) {

// 	var userInfoProps dashboardcomponents.UserInfoProps

// 	alerts, err := helpers.GetAlerts(c)
// 	if err != nil {
// 		alerts = []components.AlertProps{}
// 	}

// 	// TeamPageAddEmployee

// 	teamPageAddEmployeeProps := pages.TeamPageAddEmployeeProp{}
// 	var roleOptions []string
// 	roleOptions = append(roleOptions, "seller")
// 	roleOptions = append(roleOptions, "admin")
// 	var roleSelectInputProp components.InputSelectProp
// 	roleSelectInputProp.Options = roleOptions
// 	roleSelectInputProp.Id = "role"
// 	roleSelectInputProp.Label = "role"
// 	addEmployeeFormProp := components.FormLayoutSimpleProp{Action: "/accounts/employees", Method: "Post"}
// 	employeeUsernameInput := components.InputFloatingProps{Type: "text", Name: "username", Required: true, Label: "UserName"}
// 	employeeNameInput := components.InputFloatingProps{Type: "text", Name: "name", Required: true, Label: "Name"}
// 	employeeLastNameInput := components.InputFloatingProps{Type: "text", Name: "last_name", Required: true, Label: "Last Name"}
// 	employeeAddressInput := components.InputFloatingProps{Type: "text", Name: "address", Required: true, Label: "Address"}
// 	employeePhoneNumberInput := components.InputFloatingProps{Type: "text", Name: "phone_number", Required: true, Label: "Phone Number"}
// 	employeeEmailInput := components.InputFloatingProps{Type: "email", Name: "email", Required: true, Label: "Email"}
// 	employeeNationalIdInput := components.InputFloatingProps{Type: "text", Name: "national_id", Label: "National Id"}
// 	employeeBirthdayInput := components.InputFloatingProps{Type: "text", Name: "birthday", Label: "Birthday"}
// 	employeeSalaryInput := components.InputFloatingProps{Type: "number", Name: "salary", Label: "Salary"}
// 	inputProps := []components.InputFloatingProps{
// 		employeeNameInput,
// 		employeeLastNameInput,
// 		employeeAddressInput,
// 		employeePhoneNumberInput,
// 		employeeEmailInput,
// 		employeeNationalIdInput,
// 		employeeBirthdayInput,
// 		employeeSalaryInput,
// 		employeeUsernameInput,
// 	}
// 	teamPageProps := pages.TeamPageProp{}
// 	teamPageAddEmployeeProps.InputsProps = inputProps
// 	teamPageAddEmployeeProps.FormProps = addEmployeeFormProp
// 	selectInputsProps := []components.InputSelectProp{roleSelectInputProp}
// 	teamPageAddEmployeeProps.InputSelectProps = selectInputsProps

// 	// AddTAsk

// 	teamPageAddTaskProps := pages.TeamPageAddTaskProp{}
// 	addTaskFormProp := components.FormLayoutSimpleProp{Action: "/auth/add-task", Method: "Post"}
// 	taskNameInput := components.InputFloatingProps{Type: "text", Name: "name", Required: true, Label: "Name"}
// 	taskRoleInput := components.InputFloatingProps{Type: "text", Name: "role", Required: true, Label: "Role"}
// 	taskInput := components.InputFloatingProps{Type: "text", Name: "task", Required: true, Label: "Task"}
// 	deadline := components.InputFloatingProps{Type: "text", Name: "deadlin", Required: true, Label: "Deadline"}
// 	inputProps = []components.InputFloatingProps{
// 		taskNameInput,
// 		taskRoleInput,
// 		taskInput,
// 		deadline,
// 	}
// 	teamPageProps = pages.TeamPageProp{}
// 	teamPageAddTaskProps.InputsProps = inputProps
// 	teamPageAddTaskProps.FormProps = addTaskFormProp

// 	session := sessions.Default(c)
// 	userName := session.Get("UserName")
// 	var teamPageHeaderProps []pages.TeamPageHeaderProp
// 	url := strings.SplitAfter(c.Request.URL.Path, "/")
// 	currentPage := url[len(url)-1]
// 	headerElementsClass := "inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
// 	currentHeaderElementClass := "inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg active dark:text-gray-200 dark:border-gray-200"

// 	teamPageHeaderProps = append(teamPageHeaderProps, pages.TeamPageHeaderProp{Label: "Employees", Url: "/dashboard/employee/employees", Class: ""})
// 	teamPageHeaderProps = append(teamPageHeaderProps, pages.TeamPageHeaderProp{Label: "Add Employee", Url: "/dashboard/employee/add-employee", Class: ""})
// 	teamPageHeaderProps = append(teamPageHeaderProps, pages.TeamPageHeaderProp{Label: "Tasks", Url: "/dashboard/employee/tasks", Class: ""})
// 	teamPageHeaderProps = append(teamPageHeaderProps, pages.TeamPageHeaderProp{Label: "Add Tasks", Url: "/dashboard/employee/add-tasks", Class: ""})

// 	for i := 0; i < len(teamPageHeaderProps); i++ {
// 		elementPath := strings.SplitAfter(teamPageHeaderProps[i].Url, "/")
// 		elementHead := elementPath[len(elementPath)-1]
// 		if elementHead == currentPage {
// 			teamPageHeaderProps[i].Class = currentHeaderElementClass
// 		} else {
// 			teamPageHeaderProps[i].Class = headerElementsClass
// 		}

// 	}
// 	email := session.Get("Email")
// 	if userName == nil {
// 		teamPageProps.Username = email.(string)
// 	} else {
// 		teamPageProps.Username = userName.(string)
// 	}

// 	teamPageProps.TeamPageHeaderProps = teamPageHeaderProps
// 	teamPageProps.Page = currentPage

// 	pages.TeamPage(userInfoProps, teamPageProps, teamPageAddEmployeeProps, teamPageAddTaskProps, alerts).Render(c.Request.Context(), c.Writer)
// }
