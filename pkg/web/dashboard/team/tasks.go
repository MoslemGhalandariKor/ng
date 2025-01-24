package team

import (
	"fmt"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"strconv"

	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/team"
	"sort"

	"github.com/gin-gonic/gin"
)

func TasksPage(c *gin.Context) {

	taskPageProps := team.TasksPageProp{}

	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	taskPageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	teamPageHeaderProps := make([]team.TeamPageHeaderProp, 0)
	for key, value := range TeamPageHeaderProps {
		if key == "TasksPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		teamPageHeaderProps = append(teamPageHeaderProps, *value)
	}
	sort.Slice(teamPageHeaderProps, func(i, j int) bool {
		return teamPageHeaderProps[i].Url > teamPageHeaderProps[j].Url
	})

	taskPageProps.TeamPageHeaderProps = teamPageHeaderProps

	posts, err := GetAllTasks()
	if err != nil {
		fmt.Println(err)
	}

	taskinfo := []team.TaskInfoProps{}
	DeleteTaskUrl := "/dashboard/employee/delete-task/"

	for _, post := range *posts {
		taskinfo = append(taskinfo, team.TaskInfoProps{
			FirstName: post.FirstName,
			LastName:  post.LastName,
			Task:      post.EmployeeTask,
			Role:      post.Role,
			Deadline:  post.Deadline,
			DeleteUrl: DeleteTaskUrl + strconv.FormatUint(uint64(post.ID), 10),
		})
	}

	taskPageProps.TaskInfoProps = taskinfo

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, team.TasksPage(taskPageProps))
	c.Render(http.StatusOK, r)

}

func AddTaskPage(c *gin.Context) {

	addTaskPageProp := team.AddTaskPageProp{}

	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	addTaskPageProp.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	var addTaskPageContents team.AddTaskPageContentsProps

	teamPageHeaderProps := make([]team.TeamPageHeaderProp, 0)
	for key, value := range TeamPageHeaderProps {
		if key == "AddTasksPage" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		teamPageHeaderProps = append(teamPageHeaderProps, *value)
	}
	sort.Slice(teamPageHeaderProps, func(i, j int) bool {
		return teamPageHeaderProps[i].Url > teamPageHeaderProps[j].Url
	})

	addTaskPageContents.TeamPageHeaderProps = teamPageHeaderProps

	addTaskPageFormProp := team.AddTaskPageFormProps{}

	var roleSelectInputProp components.InputSelectProp
	var roleOptions []string
	roleOptions = append(roleOptions, "seller")
	roleOptions = append(roleOptions, "admin")
	roleSelectInputProp.Id = "role"
	roleSelectInputProp.Label = "role"
	roleSelectInputProp.Options = roleOptions
	selectInputsProps := []components.InputSelectProp{roleSelectInputProp}
	addTaskPageFormProp.InputSelectProps = selectInputsProps

	employeeNameInput := components.InputFloatingProps{Type: "text", Name: "first_name", Required: true, Label: "Name"}
	employeeLastNameInput := components.InputFloatingProps{Type: "text", Name: "last_name", Required: true, Label: "Last Name"}
	employeeTaskInput := components.InputFloatingProps{Type: "text", Name: "employeetask", Required: true, Label: "Task"}
	inputProps := []components.InputFloatingProps{
		employeeNameInput,
		employeeLastNameInput,
		employeeTaskInput,
	}
	addTaskPageFormProp.InputsProps = inputProps

	taskDeadlineInput := components.DatepickerForUserProps{Placeholder: "Deadline", Name: "deadline", Label: "Deadline"}
	datePickerProps := []components.DatepickerForUserProps{taskDeadlineInput}
	addTaskPageFormProp.DatePickerProps = datePickerProps

	addTaskPageContents.AddTaskPageFormProp = addTaskPageFormProp

	addEmployeeFormProp := components.FormLayoutSimpleProp{Action: "/accounts/addtask", Method: "POST"}
	addTaskPageContents.FormLayoutSimpleProp = addEmployeeFormProp

	addTaskPageProp.AddTaskPageContentsProps = addTaskPageContents

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, team.AddTaskPage(addTaskPageProp))
	c.Render(http.StatusOK, r)
}
