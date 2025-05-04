// MEP-1002
package curriculum

import (
	"ModEd/curriculum/cli/curriculum/handler"
	"ModEd/curriculum/controller"
	"fmt"
)

type CurriculumCLIParams struct {
	CurriculumController controller.CurriculumControllerInterface
	CourseController     controller.CourseControllerInterface
	ClassController      controller.ClassControllerInterface
}

func RunCurriculumModuleCLI(params *CurriculumCLIParams) {
	handlerParams := &handler.CurriculumCLIParams{
		CurriculumController: params.CurriculumController,
		CourseController:     params.CourseController,
		ClassController:      params.ClassController,
	}

	mainState := handler.NewMainMenuState(handlerParams)

	stateManager := handler.NewMenuStateManager(mainState)

	// Run menu state manager
	err := stateManager.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func printCurriculumMenu() {
	fmt.Println("\nCurriculum Module Menu:")
	fmt.Println("1. Curriculum")
	fmt.Println("2. Course")
	fmt.Println("3. Class")
	fmt.Println("0. Exit")
}

func newCurriculumCLI(params *CurriculumCLIParams) *CurriculumCLIParams {
	return &CurriculumCLIParams{
		CurriculumController: params.CurriculumController,
		CourseController:     params.CourseController,
		ClassController:      params.ClassController,
	}
}
