package team

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddEmployeeHandler(c *gin.Context) {
	Employee := &Employee{}
	Employee.FirstName = c.PostForm("first_name")
	Employee.LastName = c.PostForm("last_name")
	Employee.PhoneNumber = c.PostForm("phone_number")
	Employee.Address = c.PostForm("address")
	Employee.Email = c.PostForm("email")
	Employee.Salary = c.PostForm("salary")
	Employee.Birthday = c.PostForm("birthday")
	Employee.Username = c.PostForm("username")
	Employee.NationalId = c.PostForm("national_id")
	Employee.Role = c.PostForm("role")
	Employee.Photo = c.PostForm("employeephoto")

	_, err := CreateEmployee(Employee)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.Redirect(http.StatusFound, "/dashboard/employee/employees")
}

func DeleteEmployeeHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	_, _ = DeleteEmployeeByID(id)
	c.Redirect(http.StatusFound, "/dashboard/employee/employees")
	
}

func AddTaskHandler(c *gin.Context) {
	Task := &Task{}
	Task.FirstName = c.PostForm("first_name")
	Task.LastName = c.PostForm("last_name")
	Task.EmployeeTask = c.PostForm("employeetask")
	Task.Deadline = c.PostForm("deadline")
	Task.Role = c.PostForm("role")
	_, err := CreateTask(Task)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.Redirect(http.StatusFound, "/dashboard/employee/tasks")
}


func DeleteTaskHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	_, _ = DeleteTaskByID(id)
	c.Redirect(http.StatusFound, "/dashboard/employee/tasks")
}