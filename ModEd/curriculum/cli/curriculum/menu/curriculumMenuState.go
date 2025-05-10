package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	chandler "ModEd/curriculum/cli/curriculum/handler"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type CurriculumMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewCurriculumMenuState(
	manager *cli.CLIMenuStateManager,
	curriculumController controller.CurriculumControllerInterface,
) *CurriculumMenuState {
	handlerContext := handler.NewHandlerContext()

	curriculumMenu := &CurriculumMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	// Add menu options with corresponding handlers
	handlerContext.AddHandler("1", "Create New Curriculum", chandler.NewCreateCurriculumHandler(curriculumController))
	handlerContext.AddHandler("2", "Create Seed Curriculum", chandler.NewCreateSeedCurriculumHandler(curriculumController))
	handlerContext.AddHandler("3", "List all Curriculums", chandler.NewListCurriculumsHandler(curriculumController))
	handlerContext.AddHandler("4", "Get Curriculum by Id", chandler.NewGetCurriculumByIdHandler(curriculumController))
	handlerContext.AddHandler("5", "Update Curriculum by Id", chandler.NewUpdateCurriculumByIdHandler(curriculumController))
	handlerContext.AddHandler("6", "Delete Curriculum by Id", chandler.NewDeleteCurriculumByIdHandler(curriculumController))

	// Add back option to main menu
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_MAIN)))
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return curriculumMenu
}

func (menu *CurriculumMenuState) Render() {
	fmt.Println()
	fmt.Println(":/curriculum/curriculum")
	fmt.Println()
	fmt.Println("Curriculum Management")
	fmt.Println("Your options are:")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program")
	fmt.Println()
}

func (menu *CurriculumMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	if input != "0" {
		utils.GetUserInput("Press Enter to continue...")
	}

	return nil
}
