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
	menuManager := cli.NewCLIMenuManager()
	wrapper := curriculumController.NewWILModuleWrapper(db, courseController, classController)
	wilmoduleState := handler.NewWILModuleMenuStateHandler(menuManager, wrapper)
	menuManager.SetState(wilmoduleState)

	for {
		menuManager.Render()
		menuManager.UserInput = utils.GetUserChoice()

		if menuManager.UserInput == "exit" {
			break
		}

		err := menuManager.HandleUserInput()
		if err != nil {
			return
		}
	}
}
