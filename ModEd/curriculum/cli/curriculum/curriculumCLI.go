// MEP-1002
package curriculum

import (
	"ModEd/curriculum/cli/curriculum/handler"
	curriculumController "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/gorm"
)

func RunCurriculumModuleCLI(
	db *gorm.DB,
	courseController curriculumController.CourseControllerInterface,
	classController curriculumController.ClassControllerInterface,
	curriculumController curriculumController.CurriculumControllerInterface,
) {

	input := ""
	for input != "exit" {
		displayOptions()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)
		switch choice {
		case "1":
			handler.RunCurriculumCLIHandler(curriculumController)
		case "2":
			handler.RunCourseCLIHandler(courseController)
		case "3":
			handler.RunClassCLIHandler(classController)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func displayOptions() {
	fmt.Println("\nCurriculum Module Menu:")
	fmt.Println("1. Curriculum")
	fmt.Println("2. Course")
	fmt.Println("3. Class")
	fmt.Println("0. Exit")
}
