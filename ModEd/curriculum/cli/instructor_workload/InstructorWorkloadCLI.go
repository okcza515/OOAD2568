// MEP-1008
package instructorworkload

import (
	"ModEd/curriculum/cli/instructor_workload/handler"
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"ModEd/utils/deserializer"
	"fmt"

	"gorm.io/gorm"
)

func RunInstructorWorkloadModuleCLI(
	db *gorm.DB,
	courseController controller.CourseControllerInterface,
	classController controller.ClassControllerInterface,
	curriculumController controller.CurriculumControllerInterface,
) {
	// coursePlanController := controller.CreateCoursePlanController(db)
	// classWorkloadController := controller.CreateClassWorkloadController(db)
	seniorProjectWorkloadController := controller.CreateProjectController(db)
	studentWorkloadController := controller.CreateStudentWorkloadController(db)
	administrativeWorkloadController := controller.CreateMeetingController(db)

	// menuManager := cli.NewCLIMenuManager()
	// proxy := controller.NewInstructorWorkloadModuleWrapper(
	// 	db,
	// 	courseController,
	// 	classController,
	// 	curriculumController,
	// )
	// insturtorWorkloadModuleState := handler.NewInstructorWorkloadModuleMenuStateHandler(menuManager, proxy)
	// menuManager.SetState(insturtorWorkloadModuleState)

	input := ""
	for input != "exit" {
		displayOptions()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)
		switch choice {
		case "1": // Load CSV Seed Data
			fmt.Println("Loading CSV Seed Data...")
			migrationController := controller.NewMigrationController(db)
			migrationController.DropAllTables() // Drop
			migrationController.MigrateToDB()   // Migrate
			seedData := map[string]interface{}{
				"Meeting": &[]model.Meeting{},
				// "CoursePlan":        &[]model.CoursePlan{},
				// "ClassLecture":      &[]model.ClassLecture{},
				// "ClassMaterial":     &[]model.ClassMaterial{},
				// "StudentRequest":    &[]model.StudentRequest{},
				// "StudentAdvisor":    &[]model.StudentAdvisor{},
				// "ProjectAdvisor":    &[]model.ProjectAdvisor{},
				// "ProjectCommittee":  &[]model.ProjectCommittee{},
				// "ProjectEvaluation": &[]model.ProjectEvaluation{},
			}
			for filename, model := range seedData {
				fmt.Println("Loading file:", filename)
				fileDeserializer, err := deserializer.NewFileDeserializer("data/instructor-workload/" + filename + ".csv")
				if err != nil {
					fmt.Println("Error creating file deserializer:", filename, err)
					continue
				}

				if err := fileDeserializer.Deserialize(model); err != nil {
					fmt.Println("Error deserializing file:", filename, err)
					continue
				}
				result := db.Create(model)
				if result.Error != nil {
					fmt.Println("Error creating records for file:", filename, result.Error)
					continue
				}
			}

			fmt.Println("CSV Seed Data Loaded Successfully.")
		// case "2": // Teaching Responsibility
		// 	handler.RunAcademicWorkloadHandler(coursePlanController, classWorkloadController)
		case "3": // StudentAdvisor Workload
			handler.RunStudentAdvisorWorkloadHandler(studentWorkloadController)
		case "4": // Administrative Task
			handler.RunAdministrativeWorkloadHandler(administrativeWorkloadController)
		case "5": // Senior Projects
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
	fmt.Println("1. Load CSV Seed Data")
	fmt.Println("2. Academic")
	fmt.Println("3. Student Advisor Workload")
	fmt.Println("4. Administrative Task")
	fmt.Println("5. Senior Project")
	fmt.Println("Type 'exit' to quit")
}
