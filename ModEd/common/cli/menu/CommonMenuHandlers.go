package menu

import (
	"ModEd/common/controller"
	"ModEd/common/model"
	"ModEd/core/handler"
	"fmt"
	"os"

	"gorm.io/gorm"
)

type CommonHandlerContext struct {
	*handler.HandlerContext
	db *gorm.DB
}

func NewCommonHandlerContext(db *gorm.DB) *CommonHandlerContext {
	context := &CommonHandlerContext{
		HandlerContext: handler.NewHandlerContext(),
		db:             db,
	}
	context.initializeHandlers()
	return context
}

func confirmClearDBAction(prompt string) bool {
	var response string
	fmt.Print(prompt)
	fmt.Scan(&response)
	return response == "y" || response == "Y"
}

func (c *CommonHandlerContext) clearDatabase() error {
	studentController := controller.GetStudentController(c.db)
	instructorController := controller.GetInstructorController(c.db)
	departmentController := controller.GetDepartmentController(c.db)
	facultyController := controller.GetFacultyController(c.db)

	tx := c.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	if err := facultyController.Truncate(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to clear faculties: %v", err)
	}

	if err := departmentController.Truncate(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to clear departments: %v", err)
	}

	if err := instructorController.Truncate(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to clear instructors: %v", err)
	}

	if err := studentController.Truncate(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to clear students: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (c *CommonHandlerContext) initializeHandlers() {
	// Student handlers
	c.AddHandler("1", "List Students", handler.NewListHandlerStrategy[model.Student](
		controller.GetStudentController(c.db),
	))
	c.AddHandler("2", "Insert Students", handler.NewInsertHandlerStrategy[model.Student](
		controller.GetStudentController(c.db),
	))
	c.AddHandler("3", "Delete Student", handler.NewDeleteHandlerStrategy[model.Student](
		controller.GetStudentController(c.db),
	))

	// Instructor handlers
	c.AddHandler("4", "List Instructors", handler.NewListHandlerStrategy[model.Instructor](
		controller.GetInstructorController(c.db),
	))
	c.AddHandler("5", "Insert Instructors", handler.NewInsertHandlerStrategy[model.Instructor](
		controller.GetInstructorController(c.db),
	))
	c.AddHandler("6", "Delete Instructor", handler.NewDeleteHandlerStrategy[model.Instructor](
		controller.GetInstructorController(c.db),
	))

	// Department handlers
	c.AddHandler("7", "List Departments", handler.NewListHandlerStrategy[model.Department](
		controller.GetDepartmentController(c.db),
	))
	c.AddHandler("8", "Insert Departments", handler.NewInsertHandlerStrategy[model.Department](
		controller.GetDepartmentController(c.db),
	))
	c.AddHandler("9", "Delete Department", handler.NewDeleteHandlerStrategy[model.Department](
		controller.GetDepartmentController(c.db),
	))

	// Faculty handlers
	c.AddHandler("10", "List Faculties", handler.NewListHandlerStrategy[model.Faculty](
		controller.GetFacultyController(c.db),
	))
	c.AddHandler("11", "Insert Faculties", handler.NewInsertHandlerStrategy[model.Faculty](
		controller.GetFacultyController(c.db),
	))
	c.AddHandler("12", "Delete Faculty", handler.NewDeleteHandlerStrategy[model.Faculty](
		controller.GetFacultyController(c.db),
	))

	// Clear DB handler
	c.AddHandler("clear", "Clear all tables", &handler.FuncStrategy{
		Action: func() error {
			if !confirmClearDBAction("Are you sure you want to clear all tables? This action cannot be undone (y/n): ") {
				fmt.Println("Operation cancelled.")
				return nil
			}

			if err := c.clearDatabase(); err != nil {
				fmt.Printf("Failed to clear database: %v\n", err)
				return err
			}

			fmt.Println("All tables have been cleared successfully.")
			return nil
		},
	})

	// Exit handler
	c.AddHandler("exit", "Exit the program", &handler.FuncStrategy{
		Action: func() error {
			os.Exit(0)
			return nil
		},
	})

	c.SetMenuTitle("Common Data Management")
}
