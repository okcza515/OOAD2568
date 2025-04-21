package wilproject

import (
	"ModEd/curriculum/cli/wil-project/handler"
	"ModEd/curriculum/controller"
	curriculumController "ModEd/curriculum/controller/curriculum"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/gorm"
)

func RunWILModuleCLI(
	db *gorm.DB,
	courseController *curriculumController.CourseController,
	classController *curriculumController.ClassController,
) {

	wilprojectController := controller.CreateWILProjectController(db)
	wilprojectApplicationController := controller.CreateWILProjectApplicationController(db)
	wilprojectCurriculumController := controller.CreateWILProjectCurriculumController(db, courseController, classController)
	independentStudyController := controller.CreateIndependentStudyController(db)

	for {
		printWILModuleMenu()
		choice := utils.GetUserChoice()
		switch choice {
		case "1":
			handler.RunWILProjectCurriculumHandler(wilprojectCurriculumController)
		case "2":
			handler.RunWILProjectApplicationHandler(wilprojectApplicationController)
		case "3":
			handler.RunWILProjectHandler(wilprojectController)
		case "4":
			handler.RunIndependentStudyHandler(independentStudyController)
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
