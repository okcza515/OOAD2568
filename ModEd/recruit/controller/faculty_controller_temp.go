// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"ModEd/recruit/util"

	"gorm.io/gorm"
)

type FacultyController struct {
	DB *gorm.DB
}

func NewFacultyController(db *gorm.DB) *FacultyController {
	return &FacultyController{DB: db}
}

func (fc *FacultyController) GetAllFaculties() ([]model.Faculty, error) {
	var faculties []model.Faculty
	result := fc.DB.Find(&faculties)
	return faculties, result.Error
}

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

func (fc *FacultyController) ReadFacultyFromCSV(filePath string) error {
	importer := util.CSVImporter{
		DB:        fc.DB,
		TableName: "faculty",
	}

	return importer.ReadFromCSV(filePath)
}
