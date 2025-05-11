// MEP-1007
package submission

import (
	"ModEd/eval/cli/submission/handler"
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

func RunSubmissionModuleCLI(
	db *gorm.DB,
	submissionCtrl *controller.SubmissionController,
) {
	menu := handler.NewMenuHandler("Submission Menu", true)
	menu.Add("Submission", handler.NewSubmissionHandler(db))
	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

type LoadSeedData struct {
	db *gorm.DB
}

func (l LoadSeedData) Execute() {
	SeedJsonData := map[string]interface{}{
		"submission": &[]model.AnswerSubmission{},
	}
	for filename, model := range SeedJsonData {
		fileDeserializer, err := deserializer.NewFileDeserializer("../../data/submission/" + filename + ".json")
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