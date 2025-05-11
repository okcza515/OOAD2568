// MEP-1002
package curriculum

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/curriculum/cli/curriculum/menu"
	"ModEd/curriculum/controller"
	"fmt"
)

type CurriculumCLIParams struct {
	CurriculumController controller.CurriculumControllerInterface
	CourseController     controller.CourseControllerInterface
	ClassController      controller.ClassControllerInterface
}

func RunCurriculumModuleCLI(params *CurriculumCLIParams) {
	manager := cli.NewCLIMenuManager()

	mainMenuParams := &menu.MainMenuParams{
		CurriculumController: params.CurriculumController,
		CourseController:     params.CourseController,
		ClassController:      params.ClassController,
	}
	mainMenuState := menu.NewMainMenuState(manager, mainMenuParams)

	// Set the initial state to the main menu
	manager.SetState(mainMenuState)

	// Run the menu loop
	for {
		// Display the current menu
		manager.Render()

		// Get user input and set in manager
		manager.UserInput = util.GetCommandInput()

		if manager.UserInput == "exit" {
			break
		}

		// Handle the input
		err := manager.HandleUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			util.PressEnterToContinue()
		}
	}
}
