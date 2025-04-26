// MEP-1010 Work Integrated Learning (WIL)
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

	proxy := curriculumController.NewWILModuleProxy(db, courseController, classController)

	for {
		printWILModuleMenu()
		choice := utils.GetUserChoice()
		switch choice {
		case "1":
			handler.RunWILProjectCurriculumHandler(proxy)
		case "2":
			handler.RunWILProjectApplicationHandler(proxy)
		case "3":
			handler.RunWILProjectHandler(proxy)
		case "4":
			handler.RunIndependentStudyHandler(proxy)
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
