package controller

import (
	"ModEd/common/model"
	"ModEd/core"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type InstructorController struct {
	*core.BaseController[model.Instructor]
}

func NewInstructorController(db *gorm.DB) *InstructorController {
	db.AutoMigrate(&model.Instructor{})
	return &InstructorController{
		BaseController: core.NewBaseController[model.Instructor](db),
	}
}

func (c *InstructorController) GetAll() ([]model.Instructor, error) {
	return c.List(nil)
}

func (c *InstructorController) GetBy(field string, value interface{}) ([]model.Instructor, error) {
	return c.List(map[string]interface{}{field: value})
}

func (c *InstructorController) Update(code string, updatedData map[string]any) error {
	return c.UpdateByCondition(map[string]interface{}{
		"instructor_code": code,
	}, model.Instructor{})
}

func (c *InstructorController) UpdateByField(field string, value interface{}, updatedData map[string]any) error {
	return c.UpdateByCondition(map[string]interface{}{field: value}, model.Instructor{})
}

func (c *InstructorController) DeleteByCode(code string) error {
	return c.DeleteByCondition(map[string]interface{}{
		"instructor_code": code,
	})
}

func (c *InstructorController) Register(instructors []model.Instructor) error {
	return c.InsertMany(instructors)
}

func (c *InstructorController) Delete(field string, value interface{}) error {
	return c.DeleteByCondition(map[string]interface{}{field: value})
}

func (c *InstructorController) Truncate() error {
	return c.DeleteByCondition(map[string]interface{}{})
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
	fmt.Print("Enter StartDate (DD-MM-YYYY): ")
	var startDate string
	fmt.Scan(&startDate)
	fmt.Print("Enter Department: ")
	var department string
	fmt.Scan(&department)

	var parsedStartDate *time.Time
	if startDate != "" {
		parsed, err := time.Parse("02-01-2006", startDate)
		if err == nil {
			parsedStartDate = &parsed
		}
	}

	var departmentPtr *string
	if department != "" {
		departmentPtr = &department
	}

	instructor := model.Instructor{
		InstructorCode: instructorCode,
		FirstName:      firstname,
		LastName:       lastname,
		Email:          email,
		StartDate:      parsedStartDate,
		Department:     departmentPtr,
	}
	return c.Insert(instructor)
}

// Additional instructor-specific methods
func (c *InstructorController) GetByInstructorCode(code string) (model.Instructor, error) {
	return c.RetrieveByCondition(map[string]interface{}{
		"instructor_code": code,
	})
}

func (c *InstructorController) UpdateByInstructorCode(code string, instructor model.Instructor) error {
	return c.UpdateByCondition(map[string]interface{}{
		"instructor_code": code,
	}, instructor)
}

func (c *InstructorController) DeleteByInstructorCode(code string) error {
	return c.DeleteByCondition(map[string]interface{}{
		"instructor_code": code,
	})
}
