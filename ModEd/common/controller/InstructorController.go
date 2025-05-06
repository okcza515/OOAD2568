package controller

import (
	"ModEd/common/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type InstructorController struct {
	DB *gorm.DB
}

func NewInstructorController(db *gorm.DB) *InstructorController {
	db.AutoMigrate(&model.Instructor{})
	return &InstructorController{DB: db}
}

func (c *InstructorController) GetAll() ([]*model.Instructor, error) {
	return model.CommonModelGetAll[model.Instructor](c.DB)
}

func (c *InstructorController) GetBy(field string, value interface{}) ([]*model.Instructor, error) {
	return model.GetRecordByField[model.Instructor](c.DB, field, value)
}

func (c *InstructorController) Update(code string, updatedData map[string]any) error {
	return model.UpdateInstructorByCode(c.DB, code, updatedData)
}

func (c *InstructorController) UpdateByField(field string, value interface{}, updatedData map[string]any) error {
	return model.UpdateRecordByField[model.Instructor](c.DB, field, value, updatedData, model.Instructor{})
}

func (c *InstructorController) DeleteByCode(code string) error {
	return model.DeleteInstructorByCode(c.DB, code)
}

func (c *InstructorController) Register(instructors []*model.Instructor) error {
	return model.CommonRegister(c.DB, instructors)
}

func (c *InstructorController) Truncate() error {
	return model.TruncateModel(c.DB, "instructors")
}

func (c *InstructorController) ManualAddInstructor() error {
	fmt.Print("Enter instructor code: ")
	var instructorCode string
	fmt.Scan(&instructorCode)
	fmt.Print("Enter FirstName: ")
	var firstname string
	fmt.Scan(&firstname)
	fmt.Print("Enter LastName: ")
	var lastname string
	fmt.Scan(&lastname)
	fmt.Print("Enter Email: ")
	var email string
	fmt.Scan(&email)
	fmt.Print("Enter StartDate: ")
	var startDate string
	fmt.Scan(&startDate)
	fmt.Print("Enter Department: ")
	var department string
	fmt.Scan(&department)

	parseStartDate, err := time.Parse("02-01-2006", startDate)
	if err != nil {
		return fmt.Errorf("invalid date format: %w", err)
	}

	instructor := &model.Instructor{
		InstructorCode: instructorCode,
		FirstName:      firstname,
		LastName:       lastname,
		Email:          email,
		StartDate:      &parseStartDate,
		Department:     &department,
	}
	return model.ManualAddInstructor(c.DB, instructor)
}
