package handler

// MEP-1007

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"ModEd/eval/util"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ExamHandler struct {
	ExamCtrl *controller.ExamController
}

func NewExamHandler(db *gorm.DB) ExamHandler {
	return ExamHandler{
		ExamCtrl: controller.NewExamController(db),
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

func (e ExamHandler) Execute() {
	menu := NewMenuHandler("Exam Options", true)
	menu.Add("List Exams", ListExamsCommand{ExamCtrl: e.ExamCtrl})
	menu.Add("Retrieve Exam", RetrieveExamCommand{ExamCtrl: e.ExamCtrl})
	menu.Add("Create Exam", CreateExamCommand{ExamCtrl: e.ExamCtrl})
	menu.Add("Update Exam", &UpdateExamCommand{ExamCtrl: e.ExamCtrl})
	menu.Add("Delete Exam", &DeleteExamCommand{ExamCtrl: e.ExamCtrl})

	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

type ExamSectionMenuHandler struct {
	ExamSectionCtrl *controller.ExamSectionController
}

// Exam Commands

type ListExamsCommand struct {
	ExamCtrl *controller.ExamController
}

func (c ListExamsCommand) Execute() {
	exams, err := c.ExamCtrl.List(nil)
	if err != nil {
		println("Error listing exams:", err.Error())
		return
	}
	for _, exam := range exams {
		println("Exam:", exam.ExamName)
	}
}

type RetrieveExamCommand struct {
	ExamCtrl *controller.ExamController
}

func (c RetrieveExamCommand) Execute() {
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
	fmt.Println("Retrieved Exam:")
	fmt.Printf("Exam Name         : %s\n", exam.ExamName)
	fmt.Printf("Exam Description  : %s\n", exam.Description)
	fmt.Printf("Exam Start        : %s\n", exam.StartDate)
	fmt.Printf("Exam End          : %s\n", exam.EndDate)
}

type CreateExamCommand struct {
	ExamCtrl *controller.ExamController
}

func (c CreateExamCommand) Execute() {
	newExam := &model.Exam{
		ExamName: "Midterm",
		InstructorID: 1,
		ClassID: 1,
		Description: "this is midterm",
		ExamStatus: "Draft",
		Attempt: 1,
		StartDate: func() time.Time {
			t, _ := time.Parse("2006-01-02 15:04:05", "2023-10-01 10:00:00")
			return t
		}(),
		EndDate: func() time.Time {
			t, _ := time.Parse("2006-01-02 15:04:05", "2023-10-01 12:00:00")
			return t
		}(),
	}
	err := c.ExamCtrl.Insert(newExam)
	if err != nil {
		println("Error creating exam:", err.Error())
		return
	}
	println("Exam created successfully.")
}

type UpdateExamCommand struct {
	ExamCtrl *controller.ExamController
}

func (c *UpdateExamCommand) Execute() {
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

	// mock up
	exam.ExamName = "Mock Exam Updated"
	exam.Description = "This is a mocked updated exam"
	exam.Attempt = 5
	exam.ExamStatus = "Draft"

	updated, err := c.ExamCtrl.UpdateExam(exam)
	if err != nil {
		fmt.Println("Update failed:", err)
		return
	}

	fmt.Println("Mock exam updated successfully:")
	fmt.Printf("ID: %d, Name: %s, Status: %s\n", updated.ID, updated.ExamName, updated.ExamStatus)
}

type DeleteExamCommand struct {
	ExamCtrl *controller.ExamController
}

func (c *DeleteExamCommand) Execute(){
	examID, err := util.PromptUint("Enter Exam ID to delete: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = c.ExamCtrl.DeleteByID(uint(examID))
	if err != nil {
		fmt.Println("Error deleting exam:", err)
		return
	}
	fmt.Println("Exam deleted successfully.")
}