// MEP-1010 Work Integrated Learning (WIL)
package wilproject

import (
	"ModEd/core/cli"
	"ModEd/core/migration"
	"ModEd/curriculum/cli/wil-project/handler"
	curriculumController "ModEd/curriculum/controller"
	"fmt"
	"gorm.io/gorm"
)

func RunWILModuleCLI(
	db *gorm.DB,
	courseController curriculumController.CourseControllerInterface,
	classController curriculumController.ClassControllerInterface,
) {

	manager := migration.MigrationManager{}
	db, err := manager.
		MigrateModule(migration.MODULE_COMMON).
		MigrateModule(migration.MODULE_CURRICULUM).
		MigrateModule(migration.MODULE_WILPROJECT).
		BuildDB()

	if err != nil {
		fmt.Println("error! cannot initialize db")
	}

	menuManager := cli.NewCLIMenuManager()
	proxy := curriculumController.NewWILModuleProxy(db, courseController, classController)
	wilmoduleState := handler.NewWILModuleMenuStateHandler(menuManager, proxy)
	menuManager.SetState(wilmoduleState)

	for {
		menuManager.Render()
		err := menuManager.HandleUserInput()
		if err != nil {
			panic(err)
		}
	}
}
