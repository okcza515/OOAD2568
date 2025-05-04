package controller

import (
	"ModEd/common/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

func CreateStudentController(db *gorm.DB) *StudentController {
	db.AutoMigrate(&model.Student{})
	return &StudentController{DB: db}
}

func (c *StudentController) GetAll() ([]*model.Student, error) {
	return model.CommonModelGetAll[model.Student](c.DB)
}

func (c *StudentController) GetBy(field string, value interface{}) ([]*model.Student, error) {
	return model.GetRecordByField[model.Student](c.DB, field, value)
}

func (c *StudentController) Update(code string, updatedData map[string]any) error {
	return model.UpdateStudentByCode(c.DB, code, updatedData)
}

func (c *StudentController) DeleteByCode(code string) error {
	return model.DeleteStudentByCode(c.DB, code)
}

func (c *StudentController) Register(students []*model.Student) error {
	return model.CommonRegister(c.DB, students)
}

func (c *StudentController) Delete(field string, value interface{}) error {
	return model.DeleteRecordByField[model.Student](c.DB, field, value)
}

func (c *StudentController) Truncate() error {
	return model.TruncateModel(c.DB, "students")
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
	fmt.Print("Enter StartDate: ")
	var startDate string
	fmt.Scan(&startDate)
	fmt.Print("Enter BirthDate: ")
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
	
	parseStartDate, _ := time.Parse("2006-01-02", startDate)
	parseBirthDate, _ := time.Parse("2006-01-02", birthDate)

	student := &model.Student{
		StudentCode: studentCode, 
		FirstName: firstname, 
		LastName: lastname, 
		Email: email, 
		StartDate: parseStartDate, 
		BirthDate: parseBirthDate, 
		Program: program, 
		Department: department, 
		Status: &status,
	}	
	return model.ManualAddStudent(c.DB, student)
}