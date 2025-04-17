package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/service"
	"fmt"
)

// HRController handles student operations for the CLI.
type HRController struct {
	service *service.HRService
}

// NewHRController initializes a new HRController.
func NewHRController(dbPath string) *HRController {
	hrService := service.NewHRService(dbPath)
	return &HRController{service: hrService}
}

// ListStudents fetches all students.
func (c *HRController) ListStudents() ([]model.StudentInfo, error) {
	return c.service.ListStudents()
}

// GetStudent fetches a student by ID.
func (c *HRController) GetStudent(studentID string) (*model.StudentInfo, error) {
	return c.service.GetStudent(studentID)
}

// CreateStudent adds a new student.
func (c *HRController) CreateStudent(student model.StudentInfo) error {
	if err := c.service.CreateStudent(student); err != nil {
		return fmt.Errorf("failed to create student: %w", err)
	}
	return nil
}

// UpdateStudent modifies an existing student.
func (c *HRController) UpdateStudent(studentID string, updatedStudent model.StudentInfo) error {
	return c.service.UpdateStudent(studentID, updatedStudent)
}

// DeleteStudent removes a student.
func (c *HRController) DeleteStudent(studentID string) error {
	return c.service.DeleteStudent(studentID)
}

// func (c *HRController) UpdateStudentStatus(studentID string, newStatus string) error {

// }

// ImportStudents processes student data from a file.
func (c *HRController) ImportStudents(filePath string) error {
	return c.service.ImportStudents(filePath)
}

// SynchronizeStudents syncs students from external sources.
func (c *HRController) SynchronizeStudents() error {
	return c.service.SynchronizeStudents()
}
