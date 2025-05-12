package menu

// MEP-1007

import (
	newMenuHanler "ModEd/curriculum/cli/instructor_workload/handler"
	"ModEd/eval/cli/examination/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"

	"gorm.io/gorm"
)

type ExamMenu struct {
	ExamCtrl *controller.ExamController
	ExamSectionCtrl *controller.ExamSectionController
}

func NewExamMenu(db *gorm.DB) ExamMenu {
	return ExamMenu{
		ExamCtrl: controller.NewExamController(db),
		ExamSectionCtrl: controller.NewExamSectionController(db),
	}
}

type Back struct{}

func (b Back) Execute() {
	return
}

type UnknownCommand struct{}

func (u UnknownCommand) Execute() {
	fmt.Println("Unknown command, please try again.")
}

func (e ExamMenu) Execute() {
	menu := newMenuHanler.NewMenuHandler("Exam Menu", true)
	menu.Add("List Published Exams", handler.NewListExamByStatusHandler(e.ExamCtrl, model.ExamStatus("Publish")))
	menu.Add("List Draft Exams", handler.NewListExamByStatusHandler(e.ExamCtrl, model.ExamStatus("Draft")))
	menu.Add("List Hidden Exams", handler.NewListExamByStatusHandler(e.ExamCtrl, model.ExamStatus("Hidden")))
	menu.Add("Publish Exam", handler.NewPublishExamHandler(e.ExamCtrl))
	menu.Add("Hidden Exam", handler.NewHiddenExamHandler(e.ExamCtrl))
	menu.Add("Retrieve Exam", handler.NewRetrieveExamHandler(e.ExamCtrl, e.ExamSectionCtrl))
	menu.Add("Create Exam", handler.NewCreateExamHandler(e.ExamCtrl, e.ExamSectionCtrl))
	menu.Add("Update Exam", handler.NewUpdateExamHandler(e.ExamCtrl))
	menu.Add("Update Exam Section", handler.NewUpdateExamSectionHandler(e.ExamSectionCtrl))
	menu.Add("Delete Exam", handler.NewDeleteExamHandler(e.ExamCtrl, e.ExamSectionCtrl))

	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}