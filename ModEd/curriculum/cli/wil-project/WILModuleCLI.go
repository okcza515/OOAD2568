// MEP-1010 Work Integrated Learning (WIL)
package wilproject

import (
	"ModEd/core/cli"
	"ModEd/curriculum/cli/wil-project/handler"
	curriculumController "ModEd/curriculum/controller"
	"gorm.io/gorm"
)

func RunWILModuleCLI(
	db *gorm.DB,
	courseController curriculumController.CourseControllerInterface,
	classController curriculumController.ClassControllerInterface,
) {

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
