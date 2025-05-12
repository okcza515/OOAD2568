package handler

// MEP-1007

import (
	"ModEd/core"
	"ModEd/curriculum/cli/instructor_workload/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"ModEd/eval/util"
	"fmt"

	"gorm.io/gorm"
)

type ExamHandler struct {
	ExamCtrl *controller.ExamController
	ExamSectionCtrl *controller.ExamSectionController
}

func NewExamHandler(db *gorm.DB) ExamHandler {
	return ExamHandler{
		ExamCtrl: controller.NewExamController(db),
		ExamSectionCtrl: controller.NewExamSectionController(db),
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
	menu := handler.NewMenuHandler("Exam Menu", true)
	menu.Add("List Published Exams", &ListExamByStatusCommand{ExamCtrl: e.ExamCtrl, status: model.Publish})
	menu.Add("List Draft Exams", &ListExamByStatusCommand{ExamCtrl: e.ExamCtrl, status: model.Draft})
	menu.Add("List Closed Exams", &ListExamByStatusCommand{ExamCtrl: e.ExamCtrl, status: model.Hidden})
	menu.Add("Publish Exam", &PublishExamCommand{ExamCtrl: e.ExamCtrl})
	menu.Add("Hidden Exam", &HiddenExamCommand{ExamCtrl: e.ExamCtrl})
	menu.Add("Retrieve Exam", RetrieveExamCommand{ExamCtrl: e.ExamCtrl , ExamSectionCtrl: e.ExamSectionCtrl})
	menu.Add("Create Exam", CreateExamCommand{ExamCtrl: e.ExamCtrl , ExamSectionCtrl: e.ExamSectionCtrl})
	menu.Add("Update Exam", &UpdateExamCommand{ExamCtrl: e.ExamCtrl})
	menu.Add("Update Exam Section", &UpdateExamSectionCommand{ExamSectionCtrl: e.ExamSectionCtrl})
	menu.Add("Delete Exam", &DeleteExamCommand{ExamCtrl: e.ExamCtrl})

	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

// Exam Commands

type ListExamByStatusCommand struct {
	ExamCtrl *controller.ExamController
	status   model.ExamStatus
}

func (c *ListExamByStatusCommand) Execute() {
	exams, err := c.ExamCtrl.List(map[string]interface{}{"exam_status": c.status})
	if err != nil {
		fmt.Println("Error retrieving exams:", err)
		return
	}
	for _, exam := range exams {
		fmt.Printf("ID: %d, Name: %s, Status: %s\n", exam.ID, exam.ExamName, exam.ExamStatus)
	}
}


type RetrieveExamCommand struct {
	ExamCtrl *controller.ExamController
	ExamSectionCtrl *controller.ExamSectionController
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

type CreateExamCommand struct {
	ExamCtrl *controller.ExamController
	ExamSectionCtrl *controller.ExamSectionController
}

func (c CreateExamCommand) Execute() {
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

type UpdateExamSectionCommand struct {
	ExamSectionCtrl *controller.ExamSectionController
}

func (c *UpdateExamSectionCommand) Execute() {
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

type DeleteExamCommand struct {
	ExamCtrl *controller.ExamController
	ExamSectionCtrl *controller.ExamSectionController
}

func (c *DeleteExamCommand) Execute(){
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

type PublishExamCommand struct {
	ExamCtrl *controller.ExamController
}

func (c *PublishExamCommand) Execute() {
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

type HiddenExamCommand struct {
	ExamCtrl *controller.ExamController
}

func (c *HiddenExamCommand) Execute() {
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