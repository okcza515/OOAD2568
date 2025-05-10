// MEP-1008
package handler

import (
	controller "ModEd/curriculum/controller"
	model "ModEd/curriculum/model"
	commonModel "ModEd/common/model"
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
	academicMenu.Add("Curriculum", NewCurriculumHandler(c.db))
	academicMenu.Add("Course", NewCourseHandler(c.db))
	academicMenu.Add("Class", nil)
	academicMenu.Add("Class Material", NewClassMaterialHandler(c.db))
	academicMenu.Add("Course Plan", NewCoursePlanHandler(c.db))
	academicMenu.SetBackHandler(Back{})
	academicMenu.SetDefaultHandler(UnknownCommand{})
	academicMenu.Execute()
}

// Curriculum Menu
func (h *CurriculumHandler) Execute() {
	menu := NewMenuHandler("Curriculum Menu", true)

	menu.Add("Insert", CreateCurriculum{db: h.db})
	menu.Add("Retrieve", RetrieveCurriculum{db: h.db})
	menu.Add("Update", UpdateCurriculum{db: h.db})
	menu.Add("Delete", DeleteCurriculum{db: h.db})
	menu.Add("List All", ListCurriculums{db: h.db})

	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

type CurriculumHandler struct {
	db *gorm.DB
}

func NewCurriculumHandler(db *gorm.DB) *CurriculumHandler {
	return &CurriculumHandler{db: db}
}

type CreateCurriculum struct {
	db *gorm.DB
}
type RetrieveCurriculum struct {
	db *gorm.DB
}
type UpdateCurriculum struct {
	db *gorm.DB
}
type DeleteCurriculum struct {
	db *gorm.
		DB
}
type ListCurriculums struct {
	db *gorm.DB
}
func (c CreateCurriculum) Execute() {
	CurriculumController := controller.NewCurriculumController(c.db)
	mockCurriculum := &model.Curriculum{
		CurriculumId:  1,
		Name:"Example Curriculum",
		StartYear:     2023,
		EndYear:       2025,
		DepartmentId:  1,
		ProgramType:  commonModel.REGULAR,
	}
	if _, err := CurriculumController.CreateCurriculum(mockCurriculum); err != nil {
		fmt.Println("Error creating Curriculum:", err)
		return
	}

	fmt.Println("Curriculum created successfully!")
}
func (r RetrieveCurriculum) Execute() {
	id := utils.GetUserInputUint("Enter ID to retrieve: ")
	CurriculumController := controller.NewCurriculumController(r.db)
	curriculum, err := CurriculumController.GetCurriculum(id)
	if err != nil {
		fmt.Println("Error retrieving curriculum:", err)
		return
	}

	fmt.Println("Retrieved Curriculum:")
	fmt.Printf("ID: %d\n", curriculum.CurriculumId)
	fmt.Printf("Name: %s\n", curriculum.Name)
	fmt.Printf("Startyear: %s\n", curriculum.StartYear)
	fmt.Printf("Endyear: %s\n", curriculum.EndYear)
	fmt.Printf("DepartmentId: %s\n", curriculum.DepartmentId)
	fmt.Printf("ProgramType: %s\n", curriculum.ProgramType)
}

func (u UpdateCurriculum) Execute() {
	id := utils.GetUserInputUint("Enter ID to Update: ")
	CurriculumController := controller.NewCurriculumController(u.db)
	mockCurriculum := &model.Curriculum{
		CurriculumId:  id,
		Name:"Updated Curriculum",
		StartYear:     2020,
		EndYear:       2029,
		DepartmentId:  2,
		ProgramType:  commonModel.INTERNATIONAL,
	}
	_, err := CurriculumController.UpdateCurriculum(mockCurriculum)
	if err != nil {
		fmt.Println("Error updating Curriculum:", err)
		return
	}

	fmt.Println("Curriculum updated successfully!")
}
func (d DeleteCurriculum) Execute() {
	id := utils.GetUserInputUint("Enter ID to Delete: ")
	CurriculumController := controller.NewCurriculumController(d.db)
	_, err := CurriculumController.DeleteCurriculum(id)
	if err != nil {
		fmt.Println("Error updating Curriculum:", err)
		return
	}

	fmt.Println("Curriculum deleted successfully!")
}
func (l ListCurriculums) Execute() {	
	CurriculumController := controller.NewCurriculumController(l.db)
	curriculums, err := CurriculumController.GetCurriculums()
	if err != nil {
		fmt.Println("Error listing Curriculums:", err)
		return
	}

	if len(curriculums) == 0 {
		fmt.Println("No Curriculums found.")
		return
	}

	fmt.Println("List of All Curriculums:")
	for _, material := range curriculums {
		fmt.Printf("ID: %d, Name: %s, Start Year: %d, End Year: %d, Department ID: %d, Program Type: %s\n",
			material.CurriculumId, material.Name, material.StartYear, material.EndYear, material.DepartmentId, material.ProgramType)
	}
}

// Course Menu
func (h *CourseHandler) Execute() {
	menu := NewMenuHandler("Course Menu", true)

	menu.Add("Insert", CreateCourse{db: h.db})
	menu.Add("Retrieve", RetrieveCourse{db: h.db})
	menu.Add("Update", UpdateCourse{db: h.db})
	menu.Add("Delete", DeleteCourse{db: h.db})
	menu.Add("List All", ListCourses{db: h.db})

	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}
type CourseHandler struct {
	db *gorm.DB
}
func NewCourseHandler(db *gorm.DB) *CourseHandler {
	return &CourseHandler{db: db}
}
type CreateCourse struct {
	db *gorm.DB
}
type RetrieveCourse struct {
	db *gorm.DB
}
type UpdateCourse struct {
	db *gorm.DB
}
type DeleteCourse struct {
	db *gorm.DB
}
type ListCourses struct {
	db *gorm.DB
}
func (c CreateCourse) Execute() {
	CourseController := controller.NewCourseController(c.db)
	mockCourse := &model.Course{
		CourseId:  1,
		Name:      "Example Course",
		Description: "This is an example course",
		CurriculumId: 1,
		Optional:  false,
		CourseStatus: model.ACTIVE,
	}
	if _, err := CourseController.CreateCourse(mockCourse); err != nil {
		fmt.Println("Error creating Course:", err)
		return
	}

	fmt.Println("Course created successfully!")
}
func (r RetrieveCourse) Execute() {
	id := utils.GetUserInputUint("Enter ID to retrieve: ")
	CourseController := controller.NewCourseController(r.db)
	course, err := CourseController.GetCourse(id)
	if err != nil {
		fmt.Println("Error retrieving Course:", err)
		return
	}

	fmt.Println("Retrieved Course:")
	fmt.Printf("ID: %d\n", course.CourseId)
	fmt.Printf("Name: %s\n", course.Name)
	fmt.Printf("Description: %s\n", course.Description)
	fmt.Printf("CurriculumId: %d\n", course.CurriculumId)
	fmt.Printf("Optional: %t\n", course.Optional)
	fmt.Printf("Status: %s\n", model.CourseStatusLabel[course.CourseStatus])
}
func (u UpdateCourse) Execute() {
	id := utils.GetUserInputUint("Enter ID to Update: ")
	CourseController := controller.NewCourseController(u.db)
	mockCourse := &model.Course{
		CourseId:  id,
		Name:      "Updated Course",
		Description: "This is an updated example course",
		CurriculumId: 2,
		Optional:  true,
		CourseStatus: model.INACTIVE,
	}
	if _, err := CourseController.UpdateCourse(mockCourse); err != nil {
		fmt.Println("Error updating Course:", err)
		return
	}

	fmt.Println("Course updated successfully!")
}
func (d DeleteCourse) Execute() {
	id := utils.GetUserInputUint("Enter ID to Delete: ")
	CourseController := controller.NewCourseController(d.db)
	_, err := CourseController.DeleteCourse(id)
	if err != nil {
		fmt.Println("Error deleting Course:", err)
		return
	}

	fmt.Println("Course deleted successfully!")
}
func (l ListCourses) Execute() {
	CourseController := controller.NewCourseController(l.db)
	courses, err := CourseController.GetCourses()
	if err != nil {
		fmt.Println("Error listing Courses:", err)
		return
	}

	if len(courses) == 0 {
		fmt.Println("No Courses found.")
		return
	}

	fmt.Println("List of All Courses:")
	for _, course := range courses {
		fmt.Printf("ID: %d, Name: %s, Description: %s, Curriculum ID: %d, Optional: %t, Status: %s\n",
			course.CourseId, course.Name, course.Description, course.CurriculumId, course.Optional, model.CourseStatusLabel[course.CourseStatus])
	}
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

// CoursePlan Menu
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
