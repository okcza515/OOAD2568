// MEP-1008
package instructorworkload

import (
	"ModEd/core/cli"
	"ModEd/curriculum/cli/instructor_workload/handler"
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"

	"gorm.io/gorm"
)

func RunInstructorWorkloadModuleCLI(
	db *gorm.DB,
	courseController controller.CourseControllerInterface,
	classController controller.ClassControllerInterface,
	curriculumController controller.CurriculumControllerInterface,
) {

	menuManager := cli.NewCLIMenuManager()
	wrapper := controller.NewInstructorWorkloadModuleWrapper(
		db,
		courseController,
		classController,
		curriculumController,
	)
	instructorWorkloadModuleState := handler.NewInstructorWorkloadModuleMenuStateHandler(menuManager, wrapper)
	menuManager.SetState(instructorWorkloadModuleState)

	for {
		menuManager.Render()
		menuManager.UserInput = utils.GetUserChoice()
		err := menuManager.HandleUserInput()
		if err != nil {
			panic(err)
		}
	}
}
