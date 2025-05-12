package handler

// MEP-1007

import (
	"ModEd/eval/controller"
	"ModEd/eval/util"
	"fmt"
)

type DeleteExamHandler struct {
	ExamCtrl        *controller.ExamController
	ExamSectionCtrl *controller.ExamSectionController
}

func NewDeleteExamHandler(examCtrl *controller.ExamController, examSectionCtrl *controller.ExamSectionController) *DeleteExamHandler {
	return &DeleteExamHandler{
		ExamCtrl:        examCtrl,
		ExamSectionCtrl: examSectionCtrl,
	}
}

func (c *DeleteExamHandler) Execute(){
	examID, err := util.PromptUint("Enter Exam ID to delete: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sections, err := c.ExamSectionCtrl.RetrieveByExamID(uint(examID))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	for _, section := range sections {
		err = c.ExamSectionCtrl.DeleteByID(section.ID)
		if err != nil {
			fmt.Println("Error deleting section:", err)
			return
		}
	}

	err = c.ExamCtrl.DeleteByID(uint(examID))
	if err != nil {
		fmt.Println("Error deleting exam:", err)
		return
	}
	fmt.Println("Exam deleted successfully.")
}
