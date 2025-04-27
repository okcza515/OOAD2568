// MEP-1010 Work Integrated Learning (WIL)
package wilproject

import (
	"ModEd/core"
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

	db, err := migration.GetInstance().
		MigrateModule(core.MODULE_COMMON).
		MigrateModule(core.MODULE_CURRICULUM).
		MigrateModule(core.MODULE_WILPROJECT).
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
