package menu

import (
	"ModEd/common/controller"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
	"os"

	"gorm.io/gorm"
)

type ModelHandlerContext struct {
	*handler.HandlerContext
	db          *gorm.DB
	menuManager *cli.CLIMenuStateManager
}

func NewModelHandlerContext(db *gorm.DB, menuManager *cli.CLIMenuStateManager) *ModelHandlerContext {
	context := &ModelHandlerContext{
		HandlerContext: handler.NewHandlerContext(),
		db:             db,
		menuManager:    menuManager,
	}
	context.initializeHandlers()
	return context
}

func (c *ModelHandlerContext) Render() {
	c.ShowMenu()
}

func (c *ModelHandlerContext) HandleUserInput(input string) error {
	return c.HandleInput(input)
}

func (c *ModelHandlerContext) initializeHandlers() {
	c.AddHandler("1", "Student Management", &handler.FuncStrategy{
		Action: func() error {
			c.menuManager.GoToMenu("student")
			return nil
		},
	})
	c.AddHandler("2", "Instructor Management", &handler.FuncStrategy{
		Action: func() error {
			c.menuManager.GoToMenu("instructor")
			return nil
		},
	})
	c.AddHandler("3", "Department Management", &handler.FuncStrategy{
		Action: func() error {
			c.menuManager.GoToMenu("department")
			return nil
		},
	})
	c.AddHandler("4", "Faculty Management", &handler.FuncStrategy{
		Action: func() error {
			c.menuManager.GoToMenu("faculty")
			return nil
		},
	})

	c.AddHandler("clear", "Clear all tables", &handler.FuncStrategy{
		Action: func() error {
			if !confirmClearDBAction("Are you sure you want to clear all tables? This action cannot be undone (y/n): ") {
				fmt.Println("Operation cancelled.")
				return nil
			}

			if err := clearDatabase(c.db); err != nil {
				fmt.Printf("Failed to clear database: %v\n", err)
				return err
			}

			fmt.Println("All tables have been cleared successfully.")
			return nil
		},
	})

	c.AddHandler("exit", "Exit the program", &handler.FuncStrategy{
		Action: func() error {
			os.Exit(0)
			return nil
		},
	})

	c.SetMenuTitle("Common Data Management - Select Model")
}

func confirmClearDBAction(prompt string) bool {
	var response string
	fmt.Print(prompt)
	fmt.Scan(&response)
	return response == "y" || response == "Y"
}

func clearDatabase(db *gorm.DB) error {
	studentController := controller.GetStudentController(db)
	instructorController := controller.GetInstructorController(db)
	departmentController := controller.GetDepartmentController(db)
	facultyController := controller.GetFacultyController(db)

	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	if err := facultyController.Truncate(tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to clear faculties: %v", err)
	}

	if err := departmentController.Truncate(tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to clear departments: %v", err)
	}

	if err := instructorController.Truncate(tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to clear instructors: %v", err)
	}

	if err := studentController.Truncate(tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to clear students: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
