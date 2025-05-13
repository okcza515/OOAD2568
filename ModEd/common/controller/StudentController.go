package controller

import (
	"ModEd/common/model"
	"ModEd/core"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type StudentController struct {
	*core.BaseController[model.Student]
}

func newStudentController(db *gorm.DB) *StudentController {
	db.AutoMigrate(&model.Student{})
	return &StudentController{
		BaseController: core.NewBaseController[model.Student](db),
	}
}

func (c *StudentController) GetAll() ([]model.Student, error) {
	return c.List(nil)
}

func (c *StudentController) GetBy(field string, value interface{}) ([]model.Student, error) {
	return c.List(map[string]interface{}{field: value})
}

func (c *StudentController) Update(code string, updatedData map[string]any) error {
	return c.UpdateByCondition(map[string]interface{}{
		"student_code": code,
	}, model.Student{})
}

func (c *StudentController) UpdateByField(field string, value interface{}, updatedData map[string]any) error {
	return c.UpdateByCondition(map[string]interface{}{field: value}, model.Student{})
}

func (c *StudentController) DeleteByCode(code string) error {
	return c.DeleteByCondition(map[string]interface{}{
		"student_code": code,
	})
}

func (c *StudentController) Register(students []model.Student) error {
	return c.InsertMany(students)
}

func (c *StudentController) Delete(field string, value interface{}) error {
	return c.DeleteByCondition(map[string]interface{}{field: value})
}

func (c *StudentController) Truncate(db *gorm.DB) error {
	return model.TruncateModel(db, "students")
}

func (c *StudentController) ManualAddStudent() error {
	fmt.Print("Enter student code: ")
	var studentCode string
	fmt.Scan(&studentCode)
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
	fmt.Print("Enter BirthDate (DD-MM-YYYY): ")
	var birthDate string
	fmt.Scan(&birthDate)
	fmt.Print("Enter Program: ")
	var program model.ProgramType
	fmt.Scan(&program)
	fmt.Print("Enter Department: ")
	var department string
	fmt.Scan(&department)
	fmt.Print("Enter Status: ")
	var status model.StudentStatus
	fmt.Scan(&status)

	var parsedStartDate time.Time
	var parsedBirthDate time.Time
	var err error

	if startDate != "" {
		parsedStartDate, err = time.Parse("02-01-2006", startDate)
		if err != nil {
			return fmt.Errorf("invalid start date format: %w", err)
		}
	}

	if birthDate != "" {
		parsedBirthDate, err = time.Parse("02-01-2006", birthDate)
		if err != nil {
			return fmt.Errorf("invalid birth date format: %w", err)
		}
	}

	student := model.Student{
		StudentCode: studentCode,
		FirstName:   firstname,
		LastName:    lastname,
		Email:       email,
		StartDate:   parsedStartDate,
		BirthDate:   parsedBirthDate,
		Program:     program,
		Department:  department,
		Status:      &status,
	}
	return c.Insert(student)
}

func (c *StudentController) GetByStudentCode(code string) (model.Student, error) {
	return c.RetrieveByCondition(map[string]interface{}{
		"student_code": code,
	})
}

func (c *StudentController) UpdateByStudentCode(code string, student model.Student) error {
	return c.UpdateByCondition(map[string]interface{}{
		"student_code": code,
	}, student)
}

func (c *StudentController) DeleteByStudentCode(code string) error {
	return c.DeleteByCondition(map[string]interface{}{
		"student_code": code,
	})
}
