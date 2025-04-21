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
	coursePlanController := controller.CreateCoursePlanController(db)
	classWorkloadController := controller.CreateClassWorkloadController(db)
	seniorProjectWorkloadController := controller.CreateProjectController(db)
	studentWorkloadController := controller.CreateStudentWorkloadController(db)
	administrativeWorkloadController := controller.CreateMeetingController(db)

	input := ""
	for input != "exit" {
		displayOptions()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)
		switch choice {
		case "1": // Teaching Responsibility
			handler.RunAcademicWorkloadHandler(coursePlanController, classWorkloadController)
		case "2": // StudentAdvisor Workload
			handler.RunStudentAdvisorWorkloadHandler(studentWorkloadController)
		case "3": // Administrative Task
			handler.RunAdministrativeWorkloadHandler(administrativeWorkloadController)
		case "4": // Senior Projects
			handler.RunSeniorProjectWorkloadHandler(seniorProjectWorkloadController)
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
	fmt.Println("2. Student Advisor Workload")
	fmt.Println("3. Administrative Task")
	fmt.Println("4. Senior Project")
	fmt.Println("Type 'exit' to quit")
}
