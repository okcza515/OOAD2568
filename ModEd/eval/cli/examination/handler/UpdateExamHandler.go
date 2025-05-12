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

type UpdateExamHandler struct {
	ExamCtrl *controller.ExamController
}

func NewUpdateExamHandler(examCtrl *controller.ExamController) *UpdateExamHandler {
	return &UpdateExamHandler{
		ExamCtrl: examCtrl,
	}
}

func (c *UpdateExamHandler) Execute() {
	examID, err := util.PromptUint("Enter Exam ID to update: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	exam, err := c.ExamCtrl.RetrieveByID(uint(examID))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if exam.ExamStatus != model.ExamStatus("Draft") {
		fmt.Println("Exam is not in Draft status, cannot update.")
		return
	}

	newExamName := util.PromptString("Enter Exam Name: ")
	newDescription := util.PromptString("Enter Exam Description: ")
	newExamStatus := util.PromptString("Enter Exam Status (Draft/Publish/Hidden): ")
	newStartDate, err := util.PromptDate("Enter Exam Start Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newEndDate, err := util.PromptDate("Enter Exam Start Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// mock up
	updatedExam := &model.Exam{
		BaseModel: core.BaseModel{
			Model: gorm.Model{ID: exam.ID},
		},
		ExamName:    newExamName,
		Description: newDescription,
		ExamStatus:  model.ExamStatus(newExamStatus),
		StartDate:   newStartDate,
		EndDate:     newEndDate,
	}

	if updatedExam.StartDate.After(updatedExam.EndDate) {
		fmt.Println("Start date must be before end date.")
		return
	}

	err = c.ExamCtrl.UpdateByID(updatedExam)
	if err != nil {
		fmt.Println("Update failed:", err)
		return
	}
	fmt.Println("Exam updated successfully.")
}