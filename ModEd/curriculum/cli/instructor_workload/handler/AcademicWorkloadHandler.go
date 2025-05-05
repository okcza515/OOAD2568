package handler

import (
	"fmt"
	"gorm.io/gorm"
	controller "ModEd/curriculum/controller"
	model "ModEd/curriculum/model"

	"time"
	
)

type AcademicWorkloadHandler struct{
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
	academicMenu.Add("Class Material",NewClassMaterialHandler(c.db))
	academicMenu.Add("Course Plan", nil)
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

type ClassMaterialHandler struct{
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
			ClassId: 1,
			CourseId: 1,
			Section: 1,
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
