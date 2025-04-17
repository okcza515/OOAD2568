// filepath: /Users/kanitbunyinkgool/Desktop/OOAD2568/ModEd/hr/service/StudentHRService.go
package service

import (
	"errors"
	"fmt"
	"os"

	"ModEd/hr/model"
	"ModEd/hr/util"

	"gorm.io/gorm"
)

// HRService manages student operations.
type HRService struct {
	db *gorm.DB
}

// NewHRService initializes HRService.
func NewHRService(dbPath string) *HRService {
	db := util.OpenDatabase(dbPath)
	return &HRService{db: db}
}

// ListStudents fetches all students.
func (s *HRService) ListStudents() ([]model.StudentInfo, error) {
	var students []model.StudentInfo
	if err := s.db.Find(&students).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch students: %w", err)
	}
	return students, nil
}

// GetStudent retrieves a student by ID.
func (s *HRService) GetStudent(studentID string) (*model.StudentInfo, error) {
	var student model.StudentInfo
	if err := s.db.Where("student_code = ?", studentID).First(&student).Error; err != nil {
		return nil, fmt.Errorf("student not found: %w", err)
	}
	return &student, nil
}

// CreateStudent adds a new student.
func (s *HRService) CreateStudent(student model.StudentInfo) error {
	if student.StudentCode == "" || student.FirstName == "" || student.LastName == "" {
		return errors.New("missing required student fields")
	}
	return s.db.Create(&student).Error
}

// UpdateStudent updates an existing student.
func (s *HRService) UpdateStudent(studentID string, updatedStudent model.StudentInfo) error {
	var student model.StudentInfo
	if err := s.db.Where("student_code = ?", studentID).First(&student).Error; err != nil {
		return fmt.Errorf("student not found: %w", err)
	}

	// Apply updates only if values are provided
	if updatedStudent.FirstName != "" {
		student.FirstName = updatedStudent.FirstName
	}
	if updatedStudent.LastName != "" {
		student.LastName = updatedStudent.LastName
	}
	if updatedStudent.Gender != "" {
		student.Gender = updatedStudent.Gender
	}
	if updatedStudent.CitizenID != "" {
		student.CitizenID = updatedStudent.CitizenID
	}
	if updatedStudent.PhoneNumber != "" {
		student.PhoneNumber = updatedStudent.PhoneNumber
	}
	if updatedStudent.Email != "" {
		student.Email = updatedStudent.Email
	}

	return s.db.Save(&student).Error
}

// DeleteStudent removes a student by ID.
func (s *HRService) DeleteStudent(studentID string) error {
	if err := s.db.Where("student_code = ?", studentID).Delete(&model.StudentInfo{}).Error; err != nil {
		return fmt.Errorf("failed to delete student: %w", err)
	}
	return nil
}

// ImportStudents processes student data from a file.
func (s *HRService) ImportStudents(filePath string) error {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("file %s does not exist", filePath)
	}

	hrMapper, err := util.CreateMapper[model.StudentInfo](filePath)
	if err != nil {
		return fmt.Errorf("failed to create HR mapper: %w", err)
	}

	students := hrMapper.Map()
	for _, student := range students {
		// Check if student exists
		var existingStudent model.StudentInfo
		err := s.db.Where("student_code = ?", student.StudentCode).First(&existingStudent).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Insert new student
			if err := s.db.Create(&student).Error; err != nil {
				fmt.Printf("Failed to insert student %s: %v\n", student.StudentCode, err)
			}
		} else {
			// Update existing student
			student.ID = existingStudent.ID
			if err := s.db.Save(&student).Error; err != nil {
				fmt.Printf("Failed to update student %s: %v\n", student.StudentCode, err)
			}
		}
	}
	return nil
}

// SynchronizeStudents synchronizes students with external sources.
func (s *HRService) SynchronizeStudents() error {
	// Implement logic to sync students
	fmt.Println("Synchronizing students... (placeholder)")
	return nil
}
