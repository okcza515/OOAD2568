package menu

import (
	controller "ModEd/common/controller"
	"errors"
	"fmt"
	"os"

	"gorm.io/gorm"
)

type ReadFileHandler struct {
	path string
}

func (h *ReadFileHandler) ExecuteItem(parameters []string) {
	if _, err := os.Stat(h.path); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("*** Error: %s does not exist.\n", h.path)
		return
	} else {
		fmt.Printf("*** File %s is readable\n", h.path)
	}
}

type RegisterHandler struct {
	db   *gorm.DB
	path string
}

func (h *RegisterHandler) ExecuteItem(parameters []string) {
	if _, err := os.Stat(h.path); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("*** Error: %s does not exist.\n", h.path)
		return
	}

	submenu := NewMenuHandler()
	submenu.AppendItem("back", "Back", &BackHandler{})
	submenu.AppendItem("student", "Student", &RegisterModelHandler{db: h.db, modelType: 1, path: h.path})
	submenu.AppendItem("instructor", "Instuctor", &RegisterModelHandler{db: h.db, modelType: 2, path: h.path})
	submenu.AppendItem("department", "Department", &RegisterModelHandler{db: h.db, modelType: 3, path: h.path})
	submenu.AppendItem("faculty", "Faculty", &RegisterModelHandler{db: h.db, modelType: 4, path: h.path})

	submenu.DisplayMenu()
	choice := submenu.GetMenuChoice()
	submenu.Execute(choice, parameters)
}

type RegisterModelHandler struct {
	db        *gorm.DB
	modelType int
	path      string
}

func (h *RegisterModelHandler) ExecuteItem(parameters []string) {
	GenericRegister(h.modelType, h.db, h.path)
}

type RetrieveHandler struct {
	db *gorm.DB
}

func (h *RetrieveHandler) ExecuteItem(parameters []string) {
	submenu := NewMenuHandler()
	submenu.AppendItem("back", "Back", &BackHandler{})
	submenu.AppendItem("student", "Student", &RetrieveModelHandler{db: h.db, modelType: 1})
	submenu.AppendItem("instructor", "Instuctor", &RetrieveModelHandler{db: h.db, modelType: 2})
	submenu.AppendItem("department", "Department", &RetrieveModelHandler{db: h.db, modelType: 3})
	submenu.AppendItem("faculty", "Faculty", &RetrieveModelHandler{db: h.db, modelType: 4})

	submenu.DisplayMenu()
	choice := submenu.GetMenuChoice()
	submenu.Execute(choice, parameters)
}

type RetrieveModelHandler struct {
	db        *gorm.DB
	modelType int
}

func (h *RetrieveModelHandler) ExecuteItem(parameters []string) {
	GenericRetrieve(h.modelType, h.db)
}

type DeleteHandler struct {
	db *gorm.DB
}

func (h *DeleteHandler) ExecuteItem(parameters []string) {
	submenu := NewMenuHandler()
	submenu.AppendItem("back", "Back", &BackHandler{})
	submenu.AppendItem("student", "Student", &DeleteModelHandler{db: h.db, modelType: 1})
	submenu.AppendItem("instructor", "Instuctor", &DeleteModelHandler{db: h.db, modelType: 2})
	submenu.AppendItem("department", "Department", &DeleteModelHandler{db: h.db, modelType: 3})
	submenu.AppendItem("faculty", "Faculty", &DeleteModelHandler{db: h.db, modelType: 4})

	submenu.DisplayMenu()
	choice := submenu.GetMenuChoice()
	submenu.Execute(choice, parameters)
}

type DeleteModelHandler struct {
	db        *gorm.DB
	modelType int
}

func (h *DeleteModelHandler) ExecuteItem(parameters []string) {
	GenericDelete(h.modelType, h.db)
}

type ClearDBHandler struct {
	db *gorm.DB
}

func (h *ClearDBHandler) ExecuteItem(parameters []string) {
	if confirmAction("Are you sure you want to clear all tables? This action cannot be undone (y/n): ") {
		studentController := controller.NewStudentController(h.db)
		instructorController := controller.NewInstructorController(h.db)
		departmentController := controller.NewDepartmentController(h.db)
		facultyController := controller.NewFacultyController(h.db)

		studentController.Truncate()
		instructorController.Truncate()
		departmentController.Truncate()
		facultyController.Truncate()

		fmt.Println("All tables have been cleared.")
	} else {
		fmt.Println("Operation cancelled.")
	}
}

type ExitHandler struct{}

func (h *ExitHandler) ExecuteItem(parameters []string) {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

type BackHandler struct{}

func (h *BackHandler) ExecuteItem(parameters []string) {
}

type TestHandler struct {
	db *gorm.DB
}

func (h *TestHandler) ExecuteItem(parameters []string) {
	studentController := controller.NewStudentController(h.db)
	studentController.UpdateByField("student_code", "64070501092", map[string]any{"first_name": "John Doe"})
}

type DefaultHandler struct{}

func (h *DefaultHandler) ExecuteItem(parameters []string) {
	fmt.Println("Invalid choice. Please select a valid option.")
}
