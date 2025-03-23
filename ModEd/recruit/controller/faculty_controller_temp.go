// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"encoding/csv"
	"fmt"
	"os"

	
	"gorm.io/gorm"
)

// FacultyController handles faculty operations
type FacultyController struct {
	DB *gorm.DB
}

// NewFacultyController creates a new instance of FacultyController
func NewFacultyController(db *gorm.DB) *FacultyController {
	return &FacultyController{DB: db}
}

// GetAllFaculties retrieves all faculties
func (fc *FacultyController) GetAllFaculties() ([]model.Faculty, error) {
	var faculties []model.Faculty
	result := fc.DB.Find(&faculties)
	return faculties, result.Error
}

// GetFacultyByID retrieves a faculty by ID
func (fc *FacultyController) GetFacultyByID(id uint) (*model.Faculty, error) {
	var faculty model.Faculty
	if err := fc.DB.Preload("Departments").First(&faculty, id).Error; err != nil {
		return nil, err
	}
	return &faculty, nil
}

func (fc *FacultyController) CreateFaculty(faculty *model.Faculty) error {
	return fc.DB.Create(faculty).Error
}

func (ctrl *FacultyController) ReadFacultyFromCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Failed to read CSV: %w", err)
	}

	for _, row := range rows {
		if len(row) < 1 {
			fmt.Println("Skipping row due to insufficient data:", row)
			continue
		}

		facultyName := row[0]

		// ตรวจสอบว่ามี Faculty ชื่อเดียวกันหรือไม่
		var existingFaculty model.Faculty
		err := ctrl.DB.Where("name = ?", facultyName).First(&existingFaculty).Error
		if err == nil {
			continue
		}

		// สร้าง Faculty ใหม่
		newFaculty := model.Faculty{
			FacultyID: 0,
			Name:      facultyName,
		}

		if err := ctrl.DB.Create(&newFaculty).Error; err != nil {
			return fmt.Errorf("Failed to insert faculty: %w", err)
		}
	}

	return nil
}
