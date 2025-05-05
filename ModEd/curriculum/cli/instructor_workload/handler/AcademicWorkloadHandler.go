// MEP-1008
package handler

import (
	controller "ModEd/curriculum/controller"
	model "ModEd/curriculum/model"
	utils "ModEd/curriculum/utils"
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
	menu.Add("Retrieve", RetrieveClassMaterial{db: h.db})
	menu.Add("Update", UpdateClassMaterial{db: h.db})
	menu.Add("Delete", DeleteClassMaterial{db: h.db})
	menu.Add("List All", ListClassMaterials{db: h.db})

	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

type ClassMaterialHandler struct {
	db *gorm.DB
}

func NewClassMaterialHandler(db *gorm.DB) *ClassMaterialHandler {
	return &ClassMaterialHandler{db: db}
}

type CreateClassMaterial struct {
	db *gorm.DB
}
type RetrieveClassMaterial struct {
	db *gorm.DB
}
type UpdateClassMaterial struct {
	db *gorm.DB
}
type DeleteClassMaterial struct {
	db *gorm.
		DB
}
type ListClassMaterials struct {
	db *gorm.DB
}

func (c CreateClassMaterial) Execute() {
	ClassMaterialController := controller.NewClassMaterialController(c.db)
	mockClassMaterial := &model.ClassMaterial{
		ClassId:  1,
		FileName: "example.txt",
		FilePath: "/path/to/example.txt",
	}
	if err := ClassMaterialController.Insert(mockClassMaterial); err != nil {
		fmt.Println("Error creating ClassMaterial:", err)
		return
	}

	fmt.Println("ClassMaterial created successfully!")
}
func (r RetrieveClassMaterial) Execute() {
	id := utils.GetUserInputUint("Enter ID to retrieve: ")
	ClassMaterialController := controller.NewClassMaterialController(r.db)
	classMaterial, err := ClassMaterialController.RetrieveByID(id)
	if err != nil {
		fmt.Println("Error retrieving ClassMaterial:", err)
		return
	}

	fmt.Println("Retrieved ClassMaterial:")
	fmt.Printf("ID: %d\n", classMaterial.ID)
	fmt.Printf("ClassID: %d\n", classMaterial.ClassId)
	fmt.Printf("FileName: %s\n", classMaterial.FileName)
	fmt.Printf("FilePath: %s\n", classMaterial.FilePath)
}
func (u UpdateClassMaterial) Execute() {
	id := utils.GetUserInputUint("Enter ID to Update: ")
	ClassMaterialController := controller.NewClassMaterialController(u.db)
	mockClassMaterial := &model.ClassMaterial{
		ClassId:  id,
		FileName: "bla_example.txt",
		FilePath: "/path/to/bla_example.txt",
	}
	if err := ClassMaterialController.UpdateByID(mockClassMaterial); err != nil {
		fmt.Println("Error updating ClassMaterial:", err)
		return
	}

	fmt.Println("ClassMaterial updated successfully!")
}
func (d DeleteClassMaterial) Execute() {
	id := utils.GetUserInputUint("Enter ID to Delete: ")
	ClassMaterialController := controller.NewClassMaterialController(d.db)
	if err := ClassMaterialController.DeleteByID(id); err != nil {
		fmt.Println("Error deleting ClassMaterial:", err)
		return
	}

	fmt.Println("ClassMaterial deleted successfully!")
}
func (l ListClassMaterials) Execute() {
	ClassMaterialController := controller.NewClassMaterialController(l.db)

	classMaterials, err := ClassMaterialController.List(nil)
	if err != nil {
		fmt.Println("Error listing ClassMaterials:", err)
		return
	}

	if len(classMaterials) == 0 {
		fmt.Println("No ClassMaterials found.")
		return
	}

	fmt.Println("List of All Class Materials:")
	for _, material := range classMaterials {
		fmt.Printf("ID: %d, Class ID: %d, File Name: %s, File Path: %s\n",
			material.ID, material.ClassId, material.FileName, material.FilePath)
	}
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
	coursePlanMenu.Add("Retrieve Course Plan", RetrieveCoursePlan{db: c.db})
	coursePlanMenu.Add("Update Course Plan", UpdateCoursePlan{db: c.db})
	coursePlanMenu.Add("Delete Course Plan", DeleteCoursePlan{db: c.db})
	coursePlanMenu.Add("List All Course Plans", ListAllCoursePlans{db: c.db})
	coursePlanMenu.Add("List Upcoming Course Plans", ListAllCoursePlans{db: c.db})
	coursePlanMenu.SetBackHandler(Back{})
	coursePlanMenu.SetDefaultHandler(UnknownCommand{})
	coursePlanMenu.Execute()
}

type CreateCoursePlan struct {
	db *gorm.DB
}
type RetrieveCoursePlan struct {
	db *gorm.DB
}
type UpdateCoursePlan struct {
	db *gorm.DB
}
type DeleteCoursePlan struct {
	db *gorm.DB
}
type ListAllCoursePlans struct {
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

func (r RetrieveCoursePlan) Execute() {
	id := utils.GetUserInputUint("Enter ID to retrieve: ")
	coursePlanController := controller.NewCoursePlanController(r.db)

	coursePlan, err := coursePlanController.RetrieveByID(id)
	if err != nil {
		fmt.Println("Error retrieving course plan:", err)
		return
	}

	fmt.Println("Retrieved Course Plan:")
	fmt.Printf("ID: %d\n", coursePlan.ID)
	fmt.Printf("Course ID: %d\n", coursePlan.CourseId)
	fmt.Printf("Week: %d\n", coursePlan.Week)
	fmt.Printf("Date: %s\n", coursePlan.Date.Format("2006-01-02"))
	fmt.Printf("Instructor ID: %d\n", coursePlan.InstructorId)
	fmt.Printf("Topic: %s\n", coursePlan.Topic)
	fmt.Printf("Description: %s\n", coursePlan.Description)
}

func (u UpdateCoursePlan) Execute() {
	id := utils.GetUserInputUint("Enter ID to Update: ")
	coursePlanController := controller.NewCoursePlanController(u.db)
	mockCoursePlan := &model.CoursePlan{
		CourseId:     id,
		Week:         2,
		Date:         time.Now().AddDate(0, 0, 7),
		InstructorId: 101,
		Topic:        "Updated Introduction of Course",
		Description:  "Updated detailing the course objectives and syllabus",
	}

	if err := coursePlanController.UpdateByID(mockCoursePlan); err != nil {
		fmt.Println("Error updating course plan:", err)
		return
	}

	fmt.Println("Course Plan updated successfully!")
}
func (d DeleteCoursePlan) Execute() {
	id := utils.GetUserInputUint("Enter ID to Delete: ")
	coursePlanController := controller.NewCoursePlanController(d.db)

	if err := coursePlanController.DeleteByID(id); err != nil {
		fmt.Println("Error deleting course plan:", err)
		return
	}

	fmt.Println("Course Plan deleted successfully!")
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
