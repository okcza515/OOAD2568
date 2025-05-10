package examination

// MEP-1007

import (
	"ModEd/eval/cli/examination/handler"
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
	menu := handler.NewMenuHandler("Exam Menu", true)
	menu.Add("Manage Exams", handler.ExamHandler{ExamCtrl: examCtrl})
	menu.Add("Manage Exam Sections", handler.ExamSectionHandler{ExamSectionCtrl: examSectionCtrl})
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
		"exam_section":      &[]model.ExamSection{},
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