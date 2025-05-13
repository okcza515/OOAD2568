package menu

import (
	"ModEd/common/controller"
	"ModEd/common/model"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"

	"gorm.io/gorm"
)

type ModelType int

const (
	ModelTypeStudent ModelType = iota
	ModelTypeInstructor
	ModelTypeDepartment
	ModelTypeFaculty
)

type ActionHandlerContext struct {
	*handler.HandlerContext
	db          *gorm.DB
	modelType   ModelType
	menuManager *cli.CLIMenuStateManager
}

func NewActionHandlerContext(db *gorm.DB, modelType ModelType, menuManager *cli.CLIMenuStateManager) *ActionHandlerContext {
	context := &ActionHandlerContext{
		HandlerContext: handler.NewHandlerContext(),
		db:             db,
		modelType:      modelType,
		menuManager:    menuManager,
	}
	context.initializeHandlers()
	return context
}

func (c *ActionHandlerContext) Render() {
	c.ShowMenu()
}

func (c *ActionHandlerContext) HandleUserInput(input string) error {
	return c.HandleInput(input)
}

func (c *ActionHandlerContext) getModelTitle() string {
	switch c.modelType {
	case ModelTypeStudent:
		return "Student"
	case ModelTypeInstructor:
		return "Instructor"
	case ModelTypeDepartment:
		return "Department"
	case ModelTypeFaculty:
		return "Faculty"
	default:
		return "Invalid Model"
	}
}

func (c *ActionHandlerContext) initializeHandlers() {
	switch c.modelType {
	case ModelTypeStudent:
		c.initializeStudentHandlers()
	case ModelTypeInstructor:
		c.initializeInstructorHandlers()
	case ModelTypeDepartment:
		c.initializeDepartmentHandlers()
	case ModelTypeFaculty:
		c.initializeFacultyHandlers()
	}

	// Back handler
	c.AddHandler("back", "Back to model selection", &handler.FuncStrategy{
		Action: func() error {
			c.menuManager.GoToMenu("main")
			return nil
		},
	})

	c.SetMenuTitle(fmt.Sprintf("%s Management - Select Action", c.getModelTitle()))
}

func (c *ActionHandlerContext) initializeStudentHandlers() {
	controller := controller.GetStudentController(c.db)
	c.AddHandler("1", "List Students", handler.NewListHandlerStrategy[model.Student](controller))
	c.AddHandler("2", "Insert Students", handler.NewInsertHandlerStrategy[model.Student](controller))
	c.AddHandler("3", "Delete Student", handler.NewDeleteHandlerStrategy[model.Student](controller))
}

func (c *ActionHandlerContext) initializeInstructorHandlers() {
	controller := controller.GetInstructorController(c.db)
	c.AddHandler("1", "List Instructors", handler.NewListHandlerStrategy[model.Instructor](controller))
	c.AddHandler("2", "Insert Instructors", handler.NewInsertHandlerStrategy[model.Instructor](controller))
	c.AddHandler("3", "Delete Instructor", handler.NewDeleteHandlerStrategy[model.Instructor](controller))
}

func (c *ActionHandlerContext) initializeDepartmentHandlers() {
	controller := controller.GetDepartmentController(c.db)
	c.AddHandler("1", "List Departments", handler.NewListHandlerStrategy[model.Department](controller))
	c.AddHandler("2", "Insert Departments", handler.NewInsertHandlerStrategy[model.Department](controller))
	c.AddHandler("3", "Delete Department", handler.NewDeleteHandlerStrategy[model.Department](controller))
}

func (c *ActionHandlerContext) initializeFacultyHandlers() {
	controller := controller.GetFacultyController(c.db)
	c.AddHandler("1", "List Faculties", handler.NewListHandlerStrategy[model.Faculty](controller))
	c.AddHandler("2", "Insert Faculties", handler.NewInsertHandlerStrategy[model.Faculty](controller))
	c.AddHandler("3", "Delete Faculty", handler.NewDeleteHandlerStrategy[model.Faculty](controller))
}
