package handler

// MEP-1007

import (
	"ModEd/core"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"ModEd/eval/util"
	"fmt"

	"gorm.io/gorm"
)


type ExamSectionHandler struct {
	ExamSectionCtrl *controller.ExamSectionController
}

func NewExamSectionHandler(db *gorm.DB) ExamSectionHandler {
	return ExamSectionHandler{
		ExamSectionCtrl: controller.NewExamSectionController(db),
	}
}


func (e ExamSectionHandler) Execute() {
	menu := NewMenuHandler("Exam Section Options", true)
	menu.Add("Create Section", CreateExamSectionCommand{ExamSectionCtrl: e.ExamSectionCtrl})
	menu.Add("List Sections", &ListExamSectionCommand{ExamSectionCtrl: e.ExamSectionCtrl})
	menu.Add("Retrieve Section", &RetrieveExamSectionCommand{ExamSectionCtrl: e.ExamSectionCtrl})
	menu.Add("Update Section", &UpdateExamSectionCommand{ExamSectionCtrl: e.ExamSectionCtrl})
	menu.Add("Delete Section", &DeleteExamSectionCommand{ExamSectionCtrl: e.ExamSectionCtrl})

	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

//Exam Section Commands

type CreateExamSectionCommand struct {
	ExamSectionCtrl *controller.ExamSectionController
}

func (c CreateExamSectionCommand) Execute() {
	newSection := &model.ExamSection{
		ExamID: 1,
		SectionNo: 1,
		Description: "2 questions 30 points",
		NumQuestions: 2,
		Score: 30.0,
	}
	err := c.ExamSectionCtrl.Insert(newSection)
	if err != nil {
		println("Error creating section:", err.Error())
		return
	}
	println("Section created successfully.")
}

type ListExamSectionCommand struct {
	ExamSectionCtrl *controller.ExamSectionController
}

func (c *ListExamSectionCommand) Execute() {
	sections, err := c.ExamSectionCtrl.List(nil)
	if err != nil {
		println("Error listing sections:", err.Error())
		return
	}
	for _,section := range sections {
		fmt.Printf("ExamID: %d, Section No: %d, Description: %s, Score: %.2f\n", section.ExamID, section.SectionNo, section.Description, section.Score)
	}
}

type RetrieveExamSectionCommand struct {
	ExamSectionCtrl *controller.ExamSectionController
}
func (c *RetrieveExamSectionCommand) Execute() {
	examID, err := util.PromptUint("Enter Exam ID to Retrieve: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	sections, err := c.ExamSectionCtrl.RetrieveByExamID(uint(examID))
	if err != nil {
		fmt.Println("Error retrieving sections:", err)
		return
	}
	for _, section := range sections {
		fmt.Printf("ExamID: %d, Section No: %d, Description: %s, Score: %.2f\n", section.ExamID, section.SectionNo, section.Description, section.Score)
	}
}

type UpdateExamSectionCommand struct {
	ExamSectionCtrl *controller.ExamSectionController
}

func (c *UpdateExamSectionCommand) Execute() {
	sectionID, err := util.PromptUint("Enter Section ID to update: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	updateSection := &model.ExamSection{
		BaseModel: core.BaseModel{
			Model: gorm.Model{ID: uint(sectionID)},
		},
		ExamID: 1,
		SectionNo: 1,
		Description: "2 questions 30 points",
		NumQuestions: 2,
		Score: 30.0,
	}
	err = c.ExamSectionCtrl.UpdateByID(updateSection)
	if err != nil {
		fmt.Println("Error updating section:", err)
		return
	}
	fmt.Println("Section updated successfully.")
}

type DeleteExamSectionCommand struct {
	ExamSectionCtrl *controller.ExamSectionController
}

func (c *DeleteExamSectionCommand) Execute() {
	sectionID, err := util.PromptUint("Enter Section ID to delete: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = c.ExamSectionCtrl.DeleteByID(uint(sectionID))
	if err != nil {
		fmt.Println("Error deleting section:", err)
		return
	}
	fmt.Println("Section deleted successfully.")
}