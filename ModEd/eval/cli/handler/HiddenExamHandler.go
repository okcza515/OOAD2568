package handler

// MEP-1007

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"ModEd/eval/util"
	"fmt"
)

type HiddenExamHandler struct {
	ExamCtrl *controller.ExamController
}

func NewHiddenExamHandler(examCtrl *controller.ExamController) *HiddenExamHandler {
	return &HiddenExamHandler{
		ExamCtrl: examCtrl,
	}
}

func (c *HiddenExamHandler) Execute() {
	examID, err := util.PromptUint("Enter Exam ID to hide: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	exam, err := c.ExamCtrl.RetrieveByID(uint(examID))
	if err != nil {
		fmt.Println("Error retrieving exam:", err)
		return
	}

	exam.ExamStatus = model.ExamStatus("Hidden")

	err = c.ExamCtrl.UpdateByID(exam)
	if err != nil {
		fmt.Println("Error hiding exam:", err)
		return
	}
	fmt.Println("Exam hidden successfully.")
}