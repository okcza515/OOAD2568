package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	chandler "ModEd/curriculum/cli/curriculum/handler"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type ClassMenuState struct {
	manager          *cli.CLIMenuStateManager
	handlerContext   *handler.HandlerContext
	controller       controller.ClassControllerInterface
	courseController controller.CourseControllerInterface
}

func NewClassMenuState(
	manager *cli.CLIMenuStateManager,
	controller controller.ClassControllerInterface,
	courseController controller.CourseControllerInterface,
) *ClassMenuState {
	handlerContext := handler.NewHandlerContext()

	state := &ClassMenuState{
		manager:          manager,
		handlerContext:   handlerContext,
		controller:       controller,
		courseController: courseController,
	}

	handlerContext.AddHandler("1", "Create New Class", chandler.NewCreateClassHandler(controller, courseController))
	handlerContext.AddHandler("2", "Create Seed Class", chandler.NewCreateSeedClassHandler(controller))
	handlerContext.AddHandler("3", "List all Classes", chandler.NewListClassesHandler(controller))
	handlerContext.AddHandler("4", "Get Class by Id", chandler.NewGetClassByIdHandler(controller))
	handlerContext.AddHandler("5", "Update Class by Id", chandler.NewUpdateClassByIdHandler(controller, courseController))
	handlerContext.AddHandler("6", "Delete Class by Id", chandler.NewDeleteClassByIdHandler(controller))

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_MAIN)))
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return state
}

func (menu *ClassMenuState) Render() {
	fmt.Println()
	fmt.Println(":/curriculum/class")
	fmt.Println()
	fmt.Println("Class Management")
	fmt.Println("Your options are:")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program")
	fmt.Println()
}

func (menu *ClassMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	if input != "0" {
		utils.GetUserInput("Press Enter to continue...")
	}

	return nil
}
