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

type UpdateExamSectionHandler struct {
	ExamSectionCtrl *controller.ExamSectionController
}

func NewUpdateExamSectionHandler(examSectionCtrl *controller.ExamSectionController) *UpdateExamSectionHandler {
	return &UpdateExamSectionHandler{
		ExamSectionCtrl: examSectionCtrl,
	}
}

func (c *UpdateExamSectionHandler) Execute() {
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
	fmt.Println("Exam Sections of Exam ID [%d]: \n", examID)
	for _, section := range sections {
		fmt.Printf("Section ID: %d, Section No: %d, Description: %s\n", section.ID, section.SectionNo, section.Description)
		fmt.Printf("Number of Questions: %d, Score: %.2f\n", section.NumQuestions, section.Score)
		fmt.Println("--------------------------------------------------")
	}

	sectionID, err := util.PromptUint("Enter Section ID to update: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, section := range sections {
		if section.ID != uint(sectionID) {
			fmt.Println("Section ID not found.")
			return
		}
	}
	newDescription := util.PromptString("Enter new description: ")
	newNumQuestions, err := util.PromptUint("Enter new number of questions: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newScore, err := util.PromptFloat("Enter new score: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	updateSection := &model.ExamSection{
		BaseModel: core.BaseModel{
			Model: gorm.Model{ID: uint(sectionID)},
		},
		ExamID: uint(examID),
		SectionNo: uint(sectionID),
		Description: newDescription,
		NumQuestions: int(newNumQuestions),
		Score: newScore,
	}
	err = c.ExamSectionCtrl.UpdateByID(updateSection)
	if err != nil {
		fmt.Println("Error updating section:", err)
		return
	}
	fmt.Println("Section updated successfully.")
}