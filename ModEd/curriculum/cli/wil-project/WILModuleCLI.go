package wilproject

import (
	"ModEd/curriculum/cli/wil-project/handler"
	curriculumController "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/gorm"
)

func RunWILModuleCLI(
	db *gorm.DB,
	courseController curriculumController.CourseControllerInterface,
	classController curriculumController.ClassControllerInterface,
) {

	facade := curriculumController.NewWILModuleFacade(db, courseController, classController)

	for {
		printWILModuleMenu()
		choice := utils.GetUserChoice()
		switch choice {
		case "1":
			handler.RunWILProjectCurriculumHandler(facade.WILProjectCurriculumController)
		case "2":
			handler.RunWILProjectApplicationHandler(facade.WILProjectApplicationController)
		case "3":
			handler.RunWILProjectHandler(facade.WILProjectController)
		case "4":
			handler.RunIndependentStudyHandler(facade.IndependentStudyController)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func printWILModuleMenu() {
	fmt.Println("\nWIL Module Menu:")
	fmt.Println("1. WIL Project Curriculum")
	fmt.Println("2. WIL Project Application")
	fmt.Println("3. WIL Project")
	fmt.Println("4. Independent Study")
	fmt.Println("0. Exit WIL Module")
}
