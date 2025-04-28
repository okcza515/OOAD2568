// MEP-1010 Work Integrated Learning (WIL)
package wilproject

import (
	"ModEd/core/cli"
	"ModEd/curriculum/cli/wil-project/handler"
	curriculumController "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"

	"gorm.io/gorm"
)

func RunWILModuleCLI(
	db *gorm.DB,
	courseController curriculumController.CourseControllerInterface,
	classController curriculumController.ClassControllerInterface,
) {

	// db, err := migration.GetInstance().
	// 	MigrateModule(core.MODULE_CURRICULUM).
	// 	MigrateModule(core.MODULE_WILPROJECT).
	// 	BuildDB()

	// if err != nil {
	// 	fmt.Println("error! cannot initialize db")
	// }

	menuManager := cli.NewCLIMenuManager()
	wrapper := curriculumController.NewWILModuleWrapper(db, courseController, classController)
	wilmoduleState := handler.NewWILModuleMenuStateHandler(menuManager, wrapper)
	menuManager.SetState(wilmoduleState)

	for {
		menuManager.Render()
		menuManager.UserInput = utils.GetUserChoice()
		err := menuManager.HandleUserInput()
		if err != nil {
			panic(err)
		}
	}
}
