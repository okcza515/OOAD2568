package handler

// MEP-1007

import (
	"ModEd/eval/controller"
	"ModEd/eval/util"
	"fmt"
)


type RetrieveExamHandler struct {
	ExamCtrl *controller.ExamController
	ExamSectionCtrl *controller.ExamSectionController
}

func NewRetrieveExamHandler(examCtrl *controller.ExamController, examSectionCtrl *controller.ExamSectionController) *RetrieveExamHandler {
	return &RetrieveExamHandler{
		ExamCtrl: examCtrl,
		ExamSectionCtrl: examSectionCtrl,
	}
}

func (c RetrieveExamHandler) Execute() {
	examID, err := util.PromptUint("Enter Exam ID to Retrieve: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	exam, err := c.ExamCtrl.RetrieveByID(uint(examID))
	if err != nil {
		fmt.Println("Error retrieving exam:", err)
		return
	}
	examSection , err := c.ExamSectionCtrl.RetrieveByExamID(exam.ID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Retrieved Exam:")
	fmt.Printf("Exam Name         : %s\n", exam.ExamName)
	fmt.Printf("Exam Description  : %s\n", exam.Description)
	fmt.Printf("Exam Start        : %s\n", exam.StartDate)
	fmt.Printf("Exam End          : %s\n", exam.EndDate)
	fmt.Printf("Exam Status       : %s\n", exam.ExamStatus)
	fmt.Println("Exam Sections:")
	for _, section := range examSection {
		fmt.Printf("Section No        : %d\n", section.SectionNo)
		fmt.Printf("Section Description: %s\n", section.Description)
		fmt.Printf("Number of Questions: %d\n", section.NumQuestions)
		fmt.Printf("Score             : %.2f\n", section.Score)
	}
}