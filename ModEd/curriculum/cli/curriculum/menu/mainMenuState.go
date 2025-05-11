package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/curriculum/controller"
	"fmt"
)

type MainMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

type MainMenuParams struct {
	CurriculumController controller.CurriculumControllerInterface
	CourseController     controller.CourseControllerInterface
	ClassController      controller.ClassControllerInterface
}

func NewMainMenuState(manager *cli.CLIMenuStateManager, params *MainMenuParams) *MainMenuState {
	// Create handler context
	handlerContext := handler.NewHandlerContext()

	// Create the main menu state
	mainMenu := &MainMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	manager.AddMenu(string(MENU_MAIN), mainMenu)

	manager.AddMenu(string(MENU_CURRICULUM), NewCurriculumMenuState(manager, params.CurriculumController))
	manager.AddMenu(string(MENU_COURSE), NewCourseMenuState(manager, params.CourseController, params.CurriculumController))
	manager.AddMenu(string(MENU_CLASS), NewClassMenuState(manager, params.ClassController, params.CourseController))

	curriculumHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_CURRICULUM)))
	courseHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_COURSE)))
	classHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_CLASS)))

	handlerContext.AddHandler("1", "Curriculum Management", curriculumHandler)
	handlerContext.AddHandler("2", "Course Management", courseHandler)
	handlerContext.AddHandler("3", "Class Management", classHandler)

	return mainMenu
}

func (menu *MainMenuState) Render() {
	fmt.Println()
	fmt.Println(":/curriculum")
	fmt.Println()
	fmt.Println("Welcome to ModEd Curriculum Module CLI!")
	fmt.Println("Here is the list of options:")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program")
	fmt.Println()
}

func (menu *MainMenuState) HandleUserInput(input string) error {
	if input == "exit" {
		return nil
	}

	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("err: Invalid input, menu '" + input + "' doesn't exist")
	}

	return nil
}
