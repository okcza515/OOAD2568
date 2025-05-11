package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	chandler "ModEd/curriculum/cli/curriculum/handler"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type CourseMenuState struct {
	manager          *cli.CLIMenuStateManager
	handlerContext   *handler.HandlerContext
	courseController controller.CourseControllerInterface
	currController   controller.CurriculumControllerInterface
}

func NewCourseMenuState(
	manager *cli.CLIMenuStateManager,
	courseController controller.CourseControllerInterface,
	currController controller.CurriculumControllerInterface,
) *CourseMenuState {
	handlerContext := handler.NewHandlerContext()

	state := &CourseMenuState{
		manager:          manager,
		handlerContext:   handlerContext,
		courseController: courseController,
		currController:   currController,
	}

	// Add menu options with corresponding handlers
	handlerContext.AddHandler("1", "Create Seed Course", chandler.NewCreateSeedCourseHandler(courseController))
	handlerContext.AddHandler("2", "Create New Course", chandler.NewCreateCourseHandler(courseController, currController))
	handlerContext.AddHandler("3", "List all Courses", chandler.NewListCoursesHandler(courseController))
	handlerContext.AddHandler("4", "Get Course by Id", chandler.NewGetCourseByIdHandler(courseController))
	handlerContext.AddHandler("5", "Update Course by Id", chandler.NewUpdateCourseByIdHandler(courseController, currController))
	handlerContext.AddHandler("6", "Delete Course by Id", chandler.NewDeleteCourseByIdHandler(courseController))

	// Add back option
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_MAIN)))
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return state
}

func (menu *CourseMenuState) Render() {
	fmt.Println()
	fmt.Println(":/curriculum/course")
	fmt.Println()
	fmt.Println("Course Management")
	fmt.Println("Your options are:")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program")
	fmt.Println()
}

func (menu *CourseMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	if input != "0" {
		utils.GetUserInput("Press Enter to continue...")
	}

	return nil
}
