// MEP-1002
package curriculum

import (
	"ModEd/curriculum/cli/curriculum/handler"
	"ModEd/curriculum/controller"
	"errors"
	"fmt"
)

type CurriculumCLIParams struct {
	CurriculumController controller.CurriculumControllerInterface
	CourseController     controller.CourseControllerInterface
	ClassController      controller.ClassControllerInterface
}

func RunCurriculumModuleCLI(params *CurriculumCLIParams) {
	curriculumCLI := newCurriculumCLI(params)
	menuManager := handler.NewMenuManager(map[string]func() error{
		"1": curriculumCLI.RunCurriculumCLIHandler,
		"2": curriculumCLI.RunCourseCLIHandler,
		"3": curriculumCLI.RunClassCLIHandler,
		"0": func() error {
			fmt.Println("Exiting...")
			return handler.ExitCommand
		},
	})

	for {
		choice := menuManager.HandlerUserInput(printCurriculumMenu)
		_, ok := menuManager.Actions[choice]
		if !ok {
			fmt.Println("Invalid option")
			continue
		}

		err := menuManager.Execute(choice)
		if err != nil {
			if errors.Is(err, handler.ExitCommand) {
				return
			}
			fmt.Println("Error executing choice:", err)
		}
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

func (c *CurriculumCLIParams) RunCurriculumCLIHandler() error {
	handler.RunCurriculumCLIHandler(c.CurriculumController)
	return nil
}

func (c *CurriculumCLIParams) RunCourseCLIHandler() error {
	handler.RunCourseCLIHandler(c.CourseController)
	return nil
}

func (c *CurriculumCLIParams) RunClassCLIHandler() error {
	handler.RunClassCLIHandler(c.ClassController)
	return nil
}
