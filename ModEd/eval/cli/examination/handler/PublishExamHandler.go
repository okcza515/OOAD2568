package handler

// MEP-1007

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"ModEd/eval/util"
	"fmt"
)

type PublishExamHandler struct {
	ExamCtrl *controller.ExamController
}

func NewPublishExamHandler(examCtrl *controller.ExamController) *PublishExamHandler {
	return &PublishExamHandler{
		ExamCtrl: examCtrl,
	}
}

func (c *PublishExamHandler) Execute() {
	examID, err := util.PromptUint("Enter Exam ID to publish: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	exam, err := c.ExamCtrl.RetrieveByID(uint(examID))
	if err != nil {
		fmt.Println("Error retrieving exam:", err)
		return
	}
	
	if exam.ExamStatus != model.ExamStatus("Draft") {
		fmt.Println("Exam is not in Draft status, cannot publish.")
		return
	}

	exam.ExamStatus = model.ExamStatus("Publish")

	err = c.ExamCtrl.UpdateByID(exam)
	if err != nil {
		fmt.Println("Error publishing exam:", err)
		return
	}
	fmt.Println("Exam published successfully.")
}