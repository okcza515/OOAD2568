package examination

// MEP-1007

import (
	newMenuHandler "ModEd/curriculum/cli/instructor_workload/handler"
	examMenu "ModEd/eval/cli/examination/menu"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"ModEd/utils/deserializer"
	"fmt"

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

func RunExamModuleCLI(
	db *gorm.DB,
	examCtrl *controller.ExamController,
	examSectionCtrl *controller.ExamSectionController,
) {
	menu := newMenuHandler.NewMenuHandler("Exam Module", true)
	menu.Add("Load Seed Data", LoadSeedData{db: db})
	menu.Add("Manage Exams", examMenu.ExamMenu{ExamCtrl: examCtrl, ExamSectionCtrl: examSectionCtrl})
	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

type LoadSeedData struct {
	db *gorm.DB
}

func (l LoadSeedData) Execute() {
	SeedJsonData := map[string]interface{}{
		"exam":     &[]model.Exam{},
	}
	for filename, model := range SeedJsonData {
		fileDeserializer, err := deserializer.NewFileDeserializer("../../data/exam/" + filename + ".json")
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