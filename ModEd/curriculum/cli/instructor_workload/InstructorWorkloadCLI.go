// MEP-1008
package instructorworkload

import (
	"fmt"

	"ModEd/curriculum/cli/instructor_workload/handler"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"

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
	menu.Add(string(handler.MENU_LOAD_SEED_DATA), LoadSeedData{db: db})
	menu.Add(string(handler.MENU_ACADEMIC), handler.NewAcademicWorkloadHandler(db))
	menu.Add(string(handler.MENU_ADMINISTRATIVE), handler.NewAdminstrativeWorkloadHandler(db))
	menu.Add(string(handler.MENU_SENIOR_PROJECT), handler.NewSeniorProjectWorkloadHandler(db))
	menu.Add(string(handler.MENU_STUDENT_ADVISOR), handler.NewStudentAdvisorWorkloadHandler(db))
	menu.Add(string(handler.MENU_WORKLOAD_REPORT), handler.NewWorkloadReportHandler(db))
	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

type LoadSeedData struct {
	db *gorm.DB
}

func (l LoadSeedData) Execute() {
	fmt.Println("Loading seed data...")
	SeedCsvData := map[string]interface{}{
		"Meeting":         &[]model.Meeting{},
		"OnlineMeeting":   &[]model.OnlineMeeting{},
		"ExternalMeeting": &[]model.ExternalMeeting{},
		"ClassMaterial":   &[]model.ClassMaterial{},
		"CoursePlan":      &[]model.CoursePlan{},
	}
	for filename, model := range SeedCsvData {
		fileDeserializer, err := deserializer.NewFileDeserializer("../../data/instructor-workload/" + filename + ".csv")
		if err != nil {
			fmt.Println("Error creating file deserializer:", filename, err)
			continue
		}

		if err := fileDeserializer.Deserialize(model); err != nil {
			fmt.Println("Error deserializing file:", filename, err)
			continue
		}
		result := l.db.Create(model)
		if result.Error != nil {
			fmt.Println("Error creating records for file:", filename, result.Error)
			continue
		}
	}

	SeedJsonData := map[string]interface{}{
		"course":     &[]model.Course{},
		"class":      &[]model.Class{},
		"curriculum": &[]model.Curriculum{},
	}
	for filename, model := range SeedJsonData {
		fileDeserializer, err := deserializer.NewFileDeserializer("../../data/curriculum/" + filename + ".json")
		if err != nil {
			fmt.Println("Error creating file deserializer:", filename, err)
			continue
		}

		if err := fileDeserializer.Deserialize(model); err != nil {
			fmt.Println("Error deserializing file:", filename, err)
			continue
		}
		result := l.db.Create(model)
		if result.Error != nil {
			fmt.Println("Error creating records for file:", filename, result.Error)
			continue
		}
	}
}
