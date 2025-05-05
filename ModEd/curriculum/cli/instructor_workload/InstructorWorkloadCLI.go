// MEP-1008
package instructorworkload

import (
	"fmt"

	"ModEd/curriculum/cli/instructor_workload/handler"

	controller "ModEd/curriculum/controller"

	"gorm.io/gorm"
)

type Back struct{}

func (b Back) Execute() {
	return
}

type UnknownCommand struct{}

func (u UnknownCommand) Execute() {
	fmt.Println("Unknown command, please try again.")
}

func RunInstructorWorkloadModuleCLI(
	db *gorm.DB,
	courseController controller.CourseControllerInterface,
	classController controller.ClassControllerInterface,
	curriculumController controller.CurriculumControllerInterface,
) {
	menu := handler.NewMenuHandler("Instructor Workload Menu", true)
	menu.Add("Academic", handler.AcademicWorkload{})
	menu.Add("Adminstrative", handler.AdminstrativeWorkload{})
	menu.Add("Senior Project", handler.SeniorProjectWorkload{})
	menu.Add("Student Advisor", handler.StudentAdvisorWokrload{})
	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}
