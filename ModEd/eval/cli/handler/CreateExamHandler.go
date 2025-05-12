package handler

// MEP-1007

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"ModEd/eval/util"
	"fmt"
)

type CreateExamHandler struct {
	ExamCtrl *controller.ExamController
	ExamSectionCtrl *controller.ExamSectionController
}

func NewCreateExamHandler(examCtrl *controller.ExamController, examSectionCtrl *controller.ExamSectionController) *CreateExamHandler {
	return &CreateExamHandler{
		ExamCtrl: examCtrl,
		ExamSectionCtrl: examSectionCtrl,
	}
}

func (c CreateExamHandler) Execute() {
	examName := util.PromptString("Enter Exam Name: ")
	instructorID, err := util.PromptUint("Enter Instructor ID: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	classID, err := util.PromptUint("Enter class ID: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	description := util.PromptString("Enter Exam Description: ")
	startDate, err := util.PromptDate("Enter Exam Start Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	endDate, err := util.PromptDate("Enter Exam Start Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	numSections, err := util.PromptUint("Enter number of sections: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sectionDescriptions := make([]string, numSections)
	for i := 0; i < int(numSections); i++ {
		sectionDescriptions[i] = util.PromptString(fmt.Sprintf("Enter description for section %d: ", i+1))
	}

	sectionNumQuestions := make([]int, numSections)
	for i := 0; i < int(numSections); i++ {
		numQuestions, err := util.PromptUint(fmt.Sprintf("Enter number of questions for section %d: ", i+1))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		sectionNumQuestions[i] = int(numQuestions)
	}

	sectionScores := make([]float64, numSections)
	for i := 0; i < int(numSections); i++ {
		score, err := util.PromptFloat(fmt.Sprintf("Enter score for section %d: ", i+1))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		sectionScores[i] = score
	}

	newExam := &model.Exam{
		ExamName: examName,
		InstructorID: uint(instructorID),
		ClassID: uint(classID),
		Description: description,
		ExamStatus: model.ExamStatus("Draft"),
		Attempt: 1,
		StartDate: startDate,
		EndDate: endDate,
	}

	if newExam.StartDate.After(newExam.EndDate) {
		fmt.Println("Start date must be before end date.")
		return
	}

	newExamSection := make([]*model.ExamSection, numSections)
	for i := 0; i < int(numSections); i++ {
		newExamSection[i] = &model.ExamSection{
			ExamID: newExam.ID,
			SectionNo: uint(i + 1),
			Description: sectionDescriptions[i],
			NumQuestions: sectionNumQuestions[i],
			Score: sectionScores[i],
		}
	}

	err = c.ExamCtrl.Insert(newExam)
	if err != nil {
		println("Error creating exam:", err.Error())
		return
	}
	err = c.ExamSectionCtrl.InsertMany(newExamSection)
	if err != nil {
		println("Error creating section:", err.Error())
		return
	}

	println("Exam created successfully.")
}