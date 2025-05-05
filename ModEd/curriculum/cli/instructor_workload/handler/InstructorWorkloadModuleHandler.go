package handler

import (
	"ModEd/core/cli"
	"ModEd/core/migration"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"
	"reflect"
)

type InstructorWorkloadModuleMenuStateHandler struct {
	menuManager *cli.CLIMenuStateManager
	wrapper     *controller.InstructorWorkloadModuleWrapper

	AdministrativeTaskMenuStateHandler    *AdministrativeTaskMenuStateHandler
	SeniorProjectWorkloadMenuStateHandler *SeniorProjectWorkloadMenuStateHandler
}

func NewInstructorWorkloadModuleMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InstructorWorkloadModuleWrapper) *InstructorWorkloadModuleMenuStateHandler {
	instructorWorkloadModuleHandler := &InstructorWorkloadModuleMenuStateHandler{
		menuManager: manager,
		wrapper:     wrapper,
	}

	instructorWorkloadModuleHandler.SeniorProjectWorkloadMenuStateHandler = NewSeniorProjectModuleStateHandler(manager, wrapper, instructorWorkloadModuleHandler)

	return instructorWorkloadModuleHandler
}

func (handler *InstructorWorkloadModuleMenuStateHandler) Render() {
	fmt.Println("\nInstructor Workload Module Menu:")
	fmt.Println("1. Load CSV Seed Data")
	fmt.Println("2. Today Workload Report")
	fmt.Println("3. Wokrload Report With Filter")
	fmt.Println("4. Student Advisor Workload")
	fmt.Println("5. Administrative Task")
	fmt.Println("6. Senior Project")
	fmt.Println("Type 'exit' to quit")
}

func LoadCSVData(pairs map[string]interface{}) error {
	for filename, model := range pairs {
		fd, err := deserializer.NewFileDeserializer("data/instructor-workload/" + filename + ".csv")
		if err != nil {
			return err
		}
		fmt.Println("Deserializing file:", filename)

		err = fd.Deserialize(model)
		if err != nil {
			return err
		}

		v := reflect.ValueOf(model)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		// ถ้า slice ว่างก็ข้าม
		if v.Kind() == reflect.Slice && v.Len() > 0 {
			result := migration.GetInstance().DB.Create(v.Interface())
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

func LoadJsonData(pairs map[string]interface{}) error {
	for filename, model := range pairs {
		fd, err := deserializer.NewFileDeserializer("data/curriculum/" + filename + ".json")
		if err != nil {
			return err
		}

		err = fd.Deserialize(model)
		if err != nil {
			return err
		}

		result := migration.GetInstance().DB.Create(model)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (handler *InstructorWorkloadModuleMenuStateHandler) HandleUserInput(input string) error {
	facadeWorkload := &controller.WorkloadReportFacade{
		ClassController:           controller.NewClassController(migration.GetInstance().DB),
		MeetingController:         controller.NewMeetingController(migration.GetInstance().DB),
		StudentWorkloadController: controller.NewStudentWorkloadController(migration.GetInstance().DB),
	}
	switch input {
	case "1":
		fmt.Println("Loading CSV data...")
		LoadCSVData(
			map[string]interface{}{
				"Meeting":           &[]model.Meeting{},
				"OnlineMeeting":     &[]model.OnlineMeeting{},
				"ExternalMeeting":   &[]model.ExternalMeeting{},
				"ProjectEvaluation": &[]model.ProjectEvaluation{},
				"StudentRequest":    &[]model.StudentRequest{},
			},
		)
		fmt.Println("CSV data loaded successfully")
		LoadJsonData(
			map[string]interface{}{
				"Class":  &[]model.Class{},
				"Course": &[]model.Course{},
			},
		)
		fmt.Println("Seed data loaded successfully")

	case "2":
		facadeWorkload.GenerateDailyWorkloadReport(1)
	case "3":
		workloadReportBuilder := controller.NewWorkloadReportBuilder(facadeWorkload)
		workloadReportBuilder.IncludeClasses().IncludeMeetings().SetDateRange("2023-01-01", "2023-12-31").Generate(1)
	case "4":
		handler.menuManager.SetState(handler.AdministrativeTaskMenuStateHandler)
	case "5":
		handler.menuManager.SetState(handler.SeniorProjectWorkloadMenuStateHandler)
	default:
		fmt.Println("Invalid input")
	}

	return nil
}
