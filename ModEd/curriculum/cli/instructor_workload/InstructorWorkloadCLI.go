package instructorworkload

import (
	"ModEd/curriculum/cli/instructor_workload/handler"
	controller "ModEd/curriculum/controller"
	curriculumController "ModEd/curriculum/controller/curriculum"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/gorm"
)

func RunInstructorWorkloadModuleCLI(
	db *gorm.DB,
	courseController *curriculumController.CourseController,
	classController *curriculumController.ClassController,
	curriculumController *curriculumController.CurriculumController,
) {

	seniorProjectWorkloadController := controller.NewProjectController(db)

	input := ""
	for input != "exit" {
		displayOptions()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)
		switch choice {
		case "1": // Teaching Responsibility
			handler.RunAcademicWorkloadHandler()
		case "2": // Administrative Tasks
			handler.RunAdminstrativeWorkloadHandler()
		case "4": // Senior Projects
			handler.RunSeniorProjectWorkloadHandler(seniorProjectWorkloadController.(*controller.ProjectController))
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func displayOptions() {
	fmt.Println("\nInstructor Workload Module Menu:")
	fmt.Println("1. Academic")
	fmt.Println("2. Administrative Task")
	fmt.Println("3. Senior Project")
	fmt.Println("Type 'exit' to quit")
}
