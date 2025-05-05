package handler

import (
	controller "ModEd/curriculum/controller"
	model "ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"

	"time"
)

type AcademicWorkloadHandler struct {
	db *gorm.DB
}

func NewAcademicWorkloadHandler(db *gorm.DB) AcademicWorkloadHandler {
	return AcademicWorkloadHandler{db: db}
}

type Back struct{}

func (b Back) Execute() {
	return
}

type UnknownCommand struct{}

func (u UnknownCommand) Execute() {
	fmt.Println("Unknown command, please try again.")
}

func (c AcademicWorkloadHandler) Execute() {
	academicMenu := NewMenuHandler("Academic Workload Menu", true)
	academicMenu.Add("Curriculum", nil)
	academicMenu.Add("Course", nil)
	academicMenu.Add("Class", nil)
	academicMenu.Add("Class Material", NewClassMaterialHandler(c.db))
	academicMenu.Add("Course Plan", NewCoursePlanHandler(c.db))
	academicMenu.SetBackHandler(Back{})
	academicMenu.SetDefaultHandler(UnknownCommand{})
	academicMenu.Execute()
}

// ClassMaterial Menu
func (h *ClassMaterialHandler) Execute() {
	menu := NewMenuHandler("Class Material Menu", true)

	menu.Add("Insert", CreateClassMaterial{db: h.db})
	menu.Add("Retrieve", nil)
	menu.Add("Update", nil)
	menu.Add("Delete", nil)

	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

type ClassMaterialHandler struct {
	db *gorm.DB
}

func NewClassMaterialHandler(db *gorm.DB) *ClassMaterialHandler {
	return &ClassMaterialHandler{}
}

type CreateClassMaterial struct {
	db *gorm.DB
}

func (c CreateClassMaterial) Execute() {
	ClassMaterialController := controller.NewClassMaterialController(c.db)
	mockClassMaterial := &model.ClassMaterial{
		ClassId: 1,
		Class: model.Class{
			ClassId:  1,
			CourseId: 1,
			Section:  1,
			Schedule: time.Now(),
		},
		FileName: "example.txt",
		FilePath: "/path/to/example.txt",
	}
	if err := ClassMaterialController.Insert(mockClassMaterial); err != nil {
		fmt.Println("Error creating ClassMaterial:", err)
		return
	}

	fmt.Println("ClassMaterial created successfully!")
}

type CoursePlanHandler struct {
	db *gorm.DB
}

func NewCoursePlanHandler(db *gorm.DB) CoursePlanHandler {
	return CoursePlanHandler{db: db}
}

func (c CoursePlanHandler) Execute() {
	coursePlanMenu := NewMenuHandler("Course Plan Menu", true)
	coursePlanMenu.Add("Create Course Plan", CreateCoursePlan{db: c.db})
	coursePlanMenu.Add("List All Course Plans", ListAllCoursePlans{db: c.db}) // Add ListAllCoursePlans option
	coursePlanMenu.SetBackHandler(Back{})
	coursePlanMenu.SetDefaultHandler(UnknownCommand{})
	coursePlanMenu.Execute()
}

type CreateCoursePlan struct {
	db *gorm.DB
}

func (c CreateCoursePlan) Execute() {
	coursePlanController := controller.NewCoursePlanController(c.db)

	mockCoursePlan := &model.CoursePlan{
		CourseId:     1,
		Week:         2,
		Date:         time.Now().AddDate(0, 0, 7),
		InstructorId: 101,
		Topic:        "Introduction of Course",
		Description:  "Detailing the course objectives and syllabus",
	}

	if err := mockCoursePlan.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	id, err := coursePlanController.CreateCoursePlan(mockCoursePlan)
	if err != nil {
		fmt.Println("Error creating course plan:", err)
		return
	}

	fmt.Println("Course Plan created successfully with ID:", id)
}

type ListAllCoursePlans struct {
	db *gorm.DB
}

func (l ListAllCoursePlans) Execute() {
	coursePlanController := controller.NewCoursePlanController(l.db)

	coursePlans, err := coursePlanController.ListAllCoursePlans()
	if err != nil {
		fmt.Println("Error listing course plans:", err)
		return
	}

	if len(coursePlans) == 0 {
		fmt.Println("No course plans found.")
		return
	}

	fmt.Println("List of All Course Plans:")
	for _, plan := range coursePlans {
		fmt.Printf("ID: %d, Course ID: %d, Week: %d, Date: %s, Instructor ID: %d, Topic: %s, Description: %s\n",
			plan.ID, plan.CourseId, plan.Week, plan.Date.Format("2006-01-02"), plan.InstructorId, plan.Topic, plan.Description)
	}
}
