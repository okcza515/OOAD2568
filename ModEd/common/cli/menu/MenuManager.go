package menu

import (
	"ModEd/core/cli"

	"gorm.io/gorm"
)

var (
	menuManager *cli.CLIMenuStateManager
)

func InitMenuManager(db *gorm.DB) *cli.CLIMenuStateManager {
	if menuManager == nil {
		manager := cli.NewCLIMenuManager()

		modelMenu := NewModelHandlerContext(db, manager)
		studentMenu := NewActionHandlerContext(db, ModelTypeStudent, manager)
		instructorMenu := NewActionHandlerContext(db, ModelTypeInstructor, manager)
		departmentMenu := NewActionHandlerContext(db, ModelTypeDepartment, manager)
		facultyMenu := NewActionHandlerContext(db, ModelTypeFaculty, manager)

		manager.AddMenu("main", modelMenu)
		manager.AddMenu("student", studentMenu)
		manager.AddMenu("instructor", instructorMenu)
		manager.AddMenu("department", departmentMenu)
		manager.AddMenu("faculty", facultyMenu)

		menuManager = manager
	}
	return menuManager
}

func GetMenuManager() *cli.CLIMenuStateManager {
	return menuManager
}
