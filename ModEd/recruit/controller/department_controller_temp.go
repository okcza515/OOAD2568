// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"encoding/csv"
	"fmt"
	"os"

	"gorm.io/gorm"
)

type DepartmentController struct {
	DB *gorm.DB
}

func NewDepartmentController(db *gorm.DB) *DepartmentController {
	return &DepartmentController{DB: db}
}

func (dc *DepartmentController) GetDepartmentsByFacultyID(facultyID uint) ([]model.Department, error) {
	var departments []model.Department
	if err := dc.DB.Where("faculty_id = ?", facultyID).Find(&departments).Error; err != nil {
		return nil, err
	}
	return departments, nil
}

func (dc *DepartmentController) GetDepartmentByID(id uint) (*model.Department, error) {
	var department model.Department
	if err := dc.DB.First(&department, id).Error; err != nil {
		return nil, err
	}
	return &department, nil
}

func (dc *DepartmentController) CreateDepartment(department *model.Department) error {
	return dc.DB.Create(department).Error
}

func (ctrl *DepartmentController) ReadDepartmentFromCSV(filePath string) error {
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

	err = ctrl.DB.Exec("DELETE FROM departments").Error
	if err != nil {
		return fmt.Errorf("failed to clear departments table: %w", err)
	}

	for _, row := range rows {
		if len(row) < 2 {
			continue
		}

		departmentName := row[0]
		facultyName := row[1]

		var faculty model.Faculty
		err := ctrl.DB.Where("name = ?", facultyName).First(&faculty).Error
		if err != nil {
			return fmt.Errorf("failed to find faculty '%s': %w", facultyName, err)
		}

		newDepartment := model.Department{
			Name:      departmentName,
			FacultyID: faculty.FacultyID,
		}

		if err := ctrl.DB.Create(&newDepartment).Error; err != nil {
			return fmt.Errorf("failed to insert department: %w", err)
		}

	}

	return nil
}
