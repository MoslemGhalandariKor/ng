package team

import (
	"nextgen/internals/db/pg"
)

func CreateEmployee(Employee *Employee) (*Employee, error) {
	if err := pg.GDB.Create(Employee).Error; err != nil {
		return nil, err
	}
	return Employee, nil
}

func DeleteEmployeeByID(id int) (*Employee, error) {
	var employee Employee

	// Fetch the employee first
	if err := pg.GDB.Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err // Return if employee not found
	}

	// Delete the employee
	if err := pg.GDB.Delete(&employee).Error; err != nil {
		return nil, err // Return if delete fails
	}

	return &employee, nil // Return the deleted employee
}

func GetAllEmployees() (*[]Employee, error) {
	var posts []Employee

	if err := pg.GDB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

func CreateTask(Task *Task) (*Task, error) {
	if err := pg.GDB.Create(Task).Error; err != nil {
		return nil, err
	}
	return Task, nil
}

func GetAllTasks() (*[]Task, error) {
	var posts []Task

	if err := pg.GDB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

func DeleteTaskByID(id int) (*Task, error) {
	var task Task

	// Fetch the task first
	if err := pg.GDB.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err // Return if task not found
	}

	// Delete the task
	if err := pg.GDB.Delete(&task).Error; err != nil {
		return nil, err // Return if delete fails
	}

	return &task, nil // Return the deleted task
}
