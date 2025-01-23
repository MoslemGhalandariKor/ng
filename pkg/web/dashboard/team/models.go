package team

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string `gorm:"unique"`
	NationalId  string `gorm:"unique"`
	Email       string `gorm:"unique"`
	Salary      string
	Birthday    string
	Username    string `gorm:"unique"`
	Role        string
	Photo       string
}

func (Employee) TableName() string {
	return "tbl_employees"
}

type Task struct {
	gorm.Model
	FirstName    string
	LastName     string
	EmployeeTask string
	Role         string
	Deadline     string
}

func (Task) TableName() string {
	return "tbl_tasks"
}
