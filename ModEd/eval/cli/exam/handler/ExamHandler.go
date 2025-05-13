package handler

// MEP-1007

import (
	assetUtil "ModEd/asset/util"
	evalUtil "ModEd/eval/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"

	"gorm.io/gorm"
)

type ExamMenuState struct {
	Manager                    *cli.CLIMenuStateManager
	wrapper                    *controller.ExamModuleWrapper
	handler                    *handler.HandlerContext
	backhandler                *handler.ChangeMenuHandlerStrategy
}

func NewExamMenuState(manager *cli.CLIMenuStateManager, wrapper *controller.ExamModuleWrapper, examModuleMenuHandler cli.MenuState) *ExamMenuState {
	return &ExamMenuState{
		Manager:                    manager,
		wrapper:                    wrapper,
		handler:                    handler.NewHandlerContext(),
		backhandler:                handler.NewChangeMenuHandlerStrategy(manager, examModuleMenuHandler),
	}
}

func (menu *ExamMenuState) Render() {
	menu.handler.SetMenuTitle("Exam management:\n")
	menu.handler.AddHandler("1", "Create a new exam.", handler.FuncStrategy{Action: menu.CreateExam})
	menu.handler.AddHandler("2", "List all exams.", handler.FuncStrategy{Action: menu.ListAllExams})
	menu.handler.AddHandler("3", "Retrieve an exam by ID.", handler.FuncStrategy{Action: menu.RetrieveExamByID})
	menu.handler.AddHandler("4", "Update an exam.", handler.FuncStrategy{Action: menu.UpdateExam})
	menu.handler.AddHandler("5", "Update exam section.", handler.FuncStrategy{Action: menu.UpdateExamSection})
	menu.handler.AddHandler("6", "List all questions by exam ID.", handler.FuncStrategy{Action: menu.ListAllQuestionsByExamID})
	menu.handler.AddHandler("7", "Delete an exam.", handler.FuncStrategy{Action: menu.DeleteExam})
	menu.handler.AddHandler("8", "Publish an exam.", handler.FuncStrategy{Action: menu.PublishExam})
	menu.handler.AddHandler("9", "Hide an exam.", handler.FuncStrategy{Action: menu.HideExam})
	menu.handler.AddHandler("b", "Back to previous menu.", menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *ExamMenuState) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *ExamMenuState) PrintExamLists(exams []*model.Exam) {
	if len(exams) == 0 {
		fmt.Println("No exams found.")
		return
	}
	columns := []string{"ID", "Exam Name", "Instructor ID", "Class ID", "Description", "Status", "Attempt", "Start Date", "End Date"}
	data := [][]string{}
	for _, exam := range exams {
		row := []string{
			fmt.Sprintf("%d", exam.ID),
			exam.ExamName,
			fmt.Sprintf("%d", exam.InstructorID),
			fmt.Sprintf("%d", exam.ClassID),
			exam.Description,
			string(exam.ExamStatus),
			fmt.Sprintf("%d", exam.Attempt),
			exam.StartDate.Format("2006-01-02 15:04:05"),
			exam.EndDate.Format("2006-01-02 15:04:05"),
		}
		data = append(data, row)
	}
	core.NewMenuIO().PrintTable(columns, data)
	fmt.Println("--------------------------------------------------")
}

func (menu *ExamMenuState) CreateExam() error {
	examName := assetUtil.GetStringInput("Enter Exam Name: ")
	instructorID := assetUtil.GetUintInput("Enter Instructor ID: ")
	classID := assetUtil.GetUintInput("Enter Class ID: ")
	description := assetUtil.GetStringInput("Enter Exam Description: ")
	startDate, err := evalUtil.GetDateTimeInput("Enter Exam Start Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		fmt.Println("Invalid start date:", err)
		return nil
	}
	endDate, err := evalUtil.GetDateTimeInput("Enter Exam End Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		fmt.Println("Invalid end date:", err)
		return nil
	}

	numSections := assetUtil.GetUintInput("Enter number of sections: ")
	sectionDescriptions := make([]string, numSections)
	sectionNumQuestions := make([]uint, numSections)
	sectionScores := make([]float64, numSections)
	for i := 0; i < int(numSections); i++ {
		sectionDescriptions[i] = assetUtil.GetStringInput(fmt.Sprintf("Enter description for section %d: ", i+1))
		numQuestions := assetUtil.GetUintInput(fmt.Sprintf("Enter number of questions for section %d: ", i+1))
		sectionNumQuestions[i] = numQuestions
		score := assetUtil.GetFloatInput(fmt.Sprintf("Enter score for section %d: ", i+1))
		sectionScores[i] = score
	}

	newExam := &model.Exam{
		ExamName:     examName,
		InstructorID: instructorID,
		ClassID:      classID,
		Description:  description,
		ExamStatus:   model.Draft,
		Attempt:      1,
		StartDate:    startDate,
		EndDate:      endDate,
	}

	if newExam.StartDate.After(newExam.EndDate) {
		fmt.Println("Start date must be before end date.")
		return nil
	}

	newExamSection := make([]*model.ExamSection, numSections)
	for i := 0; i < int(numSections); i++ {
		newExamSection[i] = &model.ExamSection{
			ExamID:       newExam.ID,
			SectionNo:    uint(i + 1),
			Description:  sectionDescriptions[i],
			NumQuestions: sectionNumQuestions[i],
			Score:        sectionScores[i],
		}
	}

	err = menu.wrapper.ExamController.Insert(newExam)
	if err != nil {
		fmt.Println("Error creating exam:", err)
		return nil
	}
	err = menu.wrapper.ExamSectionController.InsertMany(newExamSection)
	if err != nil {
		fmt.Println("Error creating exam sections:", err)
		return nil
	}
	println("Exam created successfully.")
	return nil
}

func (menu *ExamMenuState) ListAllExams() error {
	exams, err := menu.wrapper.ExamController.List(nil)
	if err != nil {
		fmt.Println("Error retrieving exams:", err)
		return nil
	}
	if len(exams) == 0 {
		fmt.Println("No exams found.")
		return nil
	}
	menu.PrintExamLists(exams)
	return nil
}

func (menu *ExamMenuState) RetrieveExamByID() error {
	examID := assetUtil.GetUintInput("Enter Exam ID to retrieve: ")
	exam, err := menu.wrapper.ExamController.List(map[string]interface{}{"id": examID})
	if err != nil {
		fmt.Println("Error retrieving exam:", err)
		return nil
	}
	sections, err := menu.wrapper.ExamSectionController.List(map[string]interface{}{"exam_id": examID})
	if err != nil {
		fmt.Println("Error retrieving exam sections:", err)
		return nil
	}
	fmt.Println("Exam Details:")
	menu.PrintExamLists(exam)
	if len(sections) == 0 {
		fmt.Println("No sections found for this exam.")
		return nil
	}
	fmt.Println("Exam Sections:")
	for _, section := range sections {
		fmt.Printf("Section ID			: %d 	| Section No   : %d 	| Description: %s\n", section.ID, section.SectionNo, section.Description)
		fmt.Printf("Number of Questions : %d 	| Score		   : %.2f\n", section.NumQuestions, section.Score)
		fmt.Println("--------------------------------------------------")
	}
	return nil
}

func (menu *ExamMenuState) UpdateExam() error {
	examID :=  assetUtil.GetUintInput("Enter Exam ID to update: ")
	exam, err := menu.wrapper.ExamController.RetrieveByID(examID)
	if err != nil {
		fmt.Println("Error retrieving exam:", err)
		return nil
	}

	if exam.ExamStatus != model.Draft {
		fmt.Println("Exam is not in Draft status, cannot update")
		return nil
	}

	newExamName := assetUtil.GetStringInput("Enter Exam Name: ")
	newDescription := assetUtil.GetStringInput("Enter Exam Description: ")
	newStartDate, err := evalUtil.GetDateTimeInput("Enter Exam Start Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		fmt.Println("Invalid start date:", err)
		return nil
	}
	newEndDate, err := evalUtil.GetDateTimeInput("Enter Exam End Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		fmt.Println("Invalid end date:", err)
		return nil
	}

	updatedExam := &model.Exam{
		BaseModel: core.BaseModel{
			Model: gorm.Model{ID: exam.ID},
		},
		ExamName:    newExamName,
		Description: newDescription,
		ExamStatus:  model.Draft,
		StartDate:   newStartDate,
		EndDate:     newEndDate,
	}

	if updatedExam.StartDate.After(updatedExam.EndDate) {
		fmt.Println("Start date must be before end date.")
		return nil
	}

	err = menu.wrapper.ExamController.UpdateByID(updatedExam)
	if err != nil {
		fmt.Println("Error updating exam:", err)
		return nil
	}

	fmt.Println("Exam updated successfully.")
	return nil
}

func (menu *ExamMenuState) UpdateExamSection() error {
	examID := assetUtil.GetUintInput("Enter Exam ID to update section: ")
	examSections, err := menu.wrapper.ExamSectionController.List(map[string]interface{}{"exam_id": examID})
	if err != nil {
		fmt.Println("Error retrieving exam sections:", err)
		return nil
	}
	if len(examSections) == 0 {
		fmt.Println("No sections found for this exam.")
		return nil
	}
	fmt.Println("Exam Sections:")
	for _, section := range examSections {
		fmt.Printf("Section ID			: %d 	| Section No   : %d 	| Description: %s\n", section.ID, section.SectionNo, section.Description)
		fmt.Printf("Number of Questions : %d 	| Score		   : %.2f\n", section.NumQuestions, section.Score)
		fmt.Println("--------------------------------------------------")
	}
	sectionID := assetUtil.GetUintInput("Enter Section ID to update: ")
	for _, section := range examSections {
		if section.ID != sectionID {
			fmt.Println("Section ID not found.")
			return nil
		}
	}
	
	newDescription := assetUtil.GetStringInput("Enter new description: ")
	newNumQuestions := assetUtil.GetUintInput("Enter new number of questions: ")
	newScore := assetUtil.GetFloatInput("Enter new score: ")

	updateSection := &model.ExamSection{
		BaseModel: core.BaseModel{
			Model: gorm.Model{ID: sectionID},
		},
		ExamID: examID,
		SectionNo: sectionID,
		Description: newDescription,
		NumQuestions: newNumQuestions,
		Score: newScore,
	}
	err = menu.wrapper.ExamSectionController.UpdateByID(updateSection)
	if err != nil {
		fmt.Println("Error updating section:", err)
		return nil
	}
	fmt.Println("Section updated successfully.")
	return nil
}

func (menu *ExamMenuState) ListAllQuestionsByExamID() error {
	examID := assetUtil.GetUintInput("Enter Exam ID to retrieve questions: ")
	question, err := menu.wrapper.ExamSectionController.List(map[string]interface{}{"exam_id": examID}, "Questions")
	if err != nil {
		fmt.Println("Error retrieving questions:", err)
		return nil
	}
	if len(question) == 0 {
		fmt.Println("No questions found for this exam.")
		return nil
	}
	fmt.Printf("Questions for Exam ID [%d]: \n", examID)
	fmt.Println("--------------------------------------------------")
	for _, q := range question {
		fmt.Printf("[%d] Question: %s\n", q.Questions[q.ID].ID, q.Questions[q.ID].ActualQuestion)
	}
	fmt.Println("--------------------------------------------------")
    return nil
}

func (menu *ExamMenuState) DeleteExam() error {
	examID :=  assetUtil.GetUintInput("Enter Exam ID to delete: ")

	err := menu.wrapper.ExamController.DeleteByID(examID)
	if err != nil {
		fmt.Printf("Error deleting exam: %s\n", err.Error())
		return nil
	}

	err = menu.wrapper.ExamSectionController.DeleteByCondition(map[string]interface{}{"exam_id": examID})
	if err != nil {
		fmt.Printf("Error deleting exam sections: %s\n", err.Error())
		return nil
	}

	fmt.Println("Exam deleted successfully.")
	return nil
}

func (menu *ExamMenuState) PublishExam() error {
	examID := assetUtil.GetUintInput("Enter Exam ID to publish: ")

	exam, err := menu.wrapper.ExamController.RetrieveByID(examID)
	if err != nil {
		fmt.Println("Error retrieving exam:", err)
		return nil
	}

	if exam.ExamStatus != model.Draft {
		fmt.Println("Exam is not in Draft status, cannot publish")
		return nil
	}

	exam.ExamStatus = model.Publish

	err = menu.wrapper.ExamController.UpdateByID(exam)
	if err != nil {
		fmt.Println("Error publishing exam:", err)
		return nil
	}

	fmt.Println("Exam published successfully.")
	return nil
}

func (menu *ExamMenuState) HideExam() error {
	examID := assetUtil.GetUintInput("Enter Exam ID to hide: ")

	exam, err := menu.wrapper.ExamController.RetrieveByID(examID)
	if err != nil {
		fmt.Println("Error retrieving exam:", err)
		return nil
	}

	exam.ExamStatus = model.Hidden

	err = menu.wrapper.ExamController.UpdateByID(exam)
	if err != nil {
		fmt.Println("Error hiding exam:", err)
		return nil
	}

	fmt.Println("Exam hidden successfully.")
	return nil
}