package team

import (
	"fmt"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/team"
    "strconv"
	"sort"

	"github.com/gin-gonic/gin"
)

var TeamPageHeaderProps = map[string]*team.TeamPageHeaderProp{
	"EmployeesPage": &team.TeamPageHeaderProp{
		Label: "Employees",
		Url:   "/dashboard/employee/employees",
		Class: ""},
	"AddEmployeePage": &team.TeamPageHeaderProp{
		Label: "Add Employee",
		Url:   "/dashboard/employee/add-employee",
		Class: ""},
	"TasksPage": &team.TeamPageHeaderProp{
		Label: "Tasks",
		Url:   "/dashboard/employee/tasks",
		Class: ""},
	"AddTasksPage": &team.TeamPageHeaderProp{
		Label: "Add Tasks",
		Url:   "/dashboard/employee/add-tasks",
		Class: ""},
}

var headerElementsClass = "inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
var currentHeaderElementClass = "inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg active dark:text-gray-200 dark:border-gray-200"

func EmployeesPage(c *gin.Context) {
	employeePageProp := team.EmployeePageProp{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	employeePageProp.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	teamPageHeaderProps := make([]team.TeamPageHeaderProp, 0)
	for key, value := range TeamPageHeaderProps {
		if key == "EmployeesPaage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		teamPageHeaderProps = append(teamPageHeaderProps, *value)
	}
	sort.Slice(teamPageHeaderProps, func(i, j int) bool {
		return teamPageHeaderProps[i].Url > teamPageHeaderProps[j].Url
	})

	employeePageProp.TeamPageHeaderProps = teamPageHeaderProps

	posts, err := GetAllEmployees()
	if err != nil {
		fmt.Println(err)
	}

	employeesInfo := []team.EmployeeInfoProps{}

	DeleteEmployeeUrl := "/dashboard/employee/delete-employee/"
	for _, post := range *posts {
		employeesInfo = append(employeesInfo, team.EmployeeInfoProps{
			FirstName:   post.FirstName,
			LastName:    post.LastName,
			Address:     post.Address,
			PhoneNumber: post.PhoneNumber,
			NationalId:  post.NationalId,
			Email:       post.Email,
			Salary:      post.Salary,
			Birthday:    post.Birthday,
			Username:    post.Username,
			Role:        post.Role,
			Photo:       post.Photo,
			DeleteUrl:   DeleteEmployeeUrl + strconv.FormatUint(uint64(post.ID), 10),
		})
	}

	employeePageProp.EmployeesInfoProps = employeesInfo

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, team.EmployeesPage(employeePageProp))
	c.Render(http.StatusOK, r)

}


func AddEmployeePage(c *gin.Context) {
	addEmployeePageProp := team.AddEmployeePageProp{}
	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	addEmployeePageProp.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	var addEmployeePageContents team.AddEmployeePageContentsProps

	teamPageHeaderProps := make([]team.TeamPageHeaderProp, 0)
	for key, value := range TeamPageHeaderProps {
		if key == "AddEmployeePage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		teamPageHeaderProps = append(teamPageHeaderProps, *value)
	}
	sort.Slice(teamPageHeaderProps, func(i, j int) bool {
		return teamPageHeaderProps[i].Url > teamPageHeaderProps[j].Url
	})

	addEmployeePageContents.TeamPageHeaderProps = teamPageHeaderProps

	addEmployeePageFormProp := team.AddEmployeePageFormProps{}

	var roleSelectInputProp components.InputSelectProp
	var roleOptions []string
	roleOptions = append(roleOptions, "seller")
	roleOptions = append(roleOptions, "admin")
	roleSelectInputProp.Id = "role"
	roleSelectInputProp.Label = "Role"
	roleSelectInputProp.Options = roleOptions
	selectInputsProps := []components.InputSelectProp{roleSelectInputProp}
	addEmployeePageFormProp.InputSelectProps = selectInputsProps

	employeeUsernameInput := components.InputFloatingProps{Type: "text", Name: "username", Required: true, Label: "UserName"}
	employeeNameInput := components.InputFloatingProps{Type: "text", Name: "first_name", Required: true, Label: "Name"}
	employeeLastNameInput := components.InputFloatingProps{Type: "text", Name: "last_name", Required: true, Label: "Last Name"}
	employeeAddressInput := components.InputFloatingProps{Type: "text", Name: "address", Required: true, Label: "Address"}
	employeePhoneNumberInput := components.InputFloatingProps{Type: "text", Name: "phone_number", Required: true, Label: "Phone Number"}
	employeeEmailInput := components.InputFloatingProps{Type: "email", Name: "email", Required: true, Label: "Email"}
	employeeNationalIdInput := components.InputFloatingProps{Type: "text", Name: "national_id", Required: true, Label: "National Id"}
	employeeSalaryInput := components.InputFloatingProps{Type: "number", Name: "salary", Required: true, Label: "Salary"}
	inputProps := []components.InputFloatingProps{
		employeeNameInput,
		employeeLastNameInput,
		employeeAddressInput,
		employeePhoneNumberInput,
		employeeEmailInput,
		employeeNationalIdInput,
		employeeSalaryInput,
		employeeUsernameInput,
	}
	addEmployeePageFormProp.InputsProps = inputProps

	employeeBirthdayInput := components.DatepickerForUserProps{Placeholder: "Birthday", Name: "birthday", Label: "Birthday"}
	datePickerProps := []components.DatepickerForUserProps{employeeBirthdayInput}
	addEmployeePageFormProp.DatepickerForUserProps = datePickerProps

	addEmployeePageContents.AddEmployeePageFormProp = addEmployeePageFormProp

	addEmployeeFormProp := components.FormLayoutSimpleProp{Action: "/accounts/addemployee", Method: "POST"}
	addEmployeePageContents.FormLayoutSimpleProp = addEmployeeFormProp

	inputEmployeePhoto := components.InputFileForUserProp{Name: "employeephoto", Label: "Upload Employee Photo"}
	inputFileProp := []components.InputFileForUserProp{inputEmployeePhoto}
	addEmployeePageFormProp.InputFileForUserProp = inputFileProp
	addEmployeePageContents.AddEmployeePageFormProp = addEmployeePageFormProp

	addEmployeePageProp.AddEmployeePageContentsProps = addEmployeePageContents
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, team.AddEmployeePage(addEmployeePageProp))
	c.Render(http.StatusOK, r)
}
