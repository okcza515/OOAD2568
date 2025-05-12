package handler

// MEP-1007

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"
)



type ListExamByStatusHandler struct {
	ExamCtrl *controller.ExamController
	status   model.ExamStatus
}

func NewListExamByStatusHandler(examCtrl *controller.ExamController, status model.ExamStatus) *ListExamByStatusHandler {
	return &ListExamByStatusHandler{
		ExamCtrl: examCtrl,
		status:   status,
	}
}

func (c *ListExamByStatusHandler) Execute() {
	exams, err := c.ExamCtrl.List(map[string]interface{}{"exam_status": c.status})
	if err != nil {
		fmt.Println("Error retrieving exams:", err)
		return
	}
	for _, exam := range exams {
		fmt.Printf("ID: %d, Name: %s, Status: %s\n", exam.ID, exam.ExamName, exam.ExamStatus)
	}
}