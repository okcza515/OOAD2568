package handler

import (
	assetUtil "ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	evalUtil "ModEd/eval/util"
	"fmt"

	"errors"

	"gorm.io/gorm"
)

type ExamMenuState struct {
	Manager                    *cli.CLIMenuStateManager
	wrapper                    *controller.ExamModuleWrapper
	handler                    *handler.HandlerContext
	backhandler                *handler.ChangeMenuHandlerStrategy
}

func NewExamMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.ExamModuleWrapper, examModuleMenuHandler cli.MenuState) *ExamMenuState {
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
	menu.handler.AddHandler("back", "Back to previous menu.", menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *ExamMenuState) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *ExamMenuState) PrintExamHeader() {
	fmt.Printf("\n%-5s %-20s %-15s %-15s %-30s %-25s %-25s %-10s\n",
		"ID", "Exam Name", "Instructor ID", "Class ID", "Description", "Start Date", "End Date", "Status")
	fmt.Printf("\n%-5s %-20s %-15s %-15s %-30s %-25s %-25s %-10s\n",
	"-----", "-------------", "-------------", "----------", "------------------", "-----------------", "-----------------", "----------")
}

func (menu *ExamMenuState) PrintExamRow(exam *model.Exam) {
	fmt.Printf("\n%-5d %-20s %-15d %-15d %-30s %-25s %-25s %-10s\n",
		exam.ID,
		truncateStr(exam.ExamName, 20),
		exam.InstructorID,
		exam.ClassID,
		truncateStr(exam.Description, 30),
		exam.StartDate.Format("2006-01-02 15:04:05"),
		exam.EndDate.Format("2006-01-02 15:04:05"),
		exam.ExamStatus,
	)
}

func truncateStr(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func (menu *ExamMenuState) PrintExamList(exams []*model.Exam) {
	if len(exams) == 0 {
		fmt.Println("\nNo exams found.")
		return
	}

	menu.PrintExamHeader()
	for _, exam := range exams {
		menu.PrintExamRow(exam)
	}
	fmt.Println()
}

func (menu *ExamMenuState) CreateExam() error {
	examName := evalUtil.PromptString("Enter Exam Name: ")
	instructorID, err := evalUtil.PromptUint("Enter Instructor ID: ")
	if err != nil {
		return errors.New("invalid instructor ID")
	}
	classID, err := evalUtil.PromptUint("Enter class ID: ")
	if err != nil {
		return errors.New("invalid class ID")
	}
	description := evalUtil.PromptString("Enter Exam Description: ")
	startDate, err := evalUtil.PromptDate("Enter Exam Start Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		return errors.New("invalid invalid date format")
	}
	endDate, err := evalUtil.PromptDate("Enter Exam End Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		return errors.New("invalid invalid date format")
	}

	numSections, err := evalUtil.PromptUint("Enter number of sections: ")
	if err != nil {
		return errors.New("invalid number of sections")
	}

	sectionDescriptions := make([]string, numSections)
	for i := 0; i < int(numSections); i++ {
		sectionDescriptions[i] = evalUtil.PromptString(fmt.Sprintf("Enter description for section %d: ", i+1))
	}

	sectionNumQuestions := make([]int, numSections)
	for i := 0; i < int(numSections); i++ {
		numQuestions, err := evalUtil.PromptUint(fmt.Sprintf("Enter number of questions for section %d: ", i+1))
		if err != nil {
			return errors.New("invalid number of questions")
		}
		sectionNumQuestions[i] = int(numQuestions)
	}

	sectionScores := make([]float64, numSections)
	for i := 0; i < int(numSections); i++ {
		score, err := evalUtil.PromptFloat(fmt.Sprintf("Enter score for section %d: ", i+1))
		if err != nil {
			return errors.New("invalid score")
		}
		sectionScores[i] = score
	}

	newExam := &model.Exam{
		ExamName:     examName,
		InstructorID: uint(instructorID),
		ClassID:      uint(classID),
		Description:  description,
		ExamStatus:   model.Draft,
		Attempt:      1,
		StartDate:    startDate,
		EndDate:      endDate,
	}

	if newExam.StartDate.After(newExam.EndDate) {
		return errors.New("start date cannot be after end date")
	}

	newExamSection := make([]*model.ExamSection, numSections)
	for i := 0; i < int(numSections); i++ {
		newExamSection[i] = &model.ExamSection{
			ExamID:       newExam.ID,
			SectionNo:    uint(i + 1),
			Description:  sectionDescriptions[i],
			NumQuestions: uint(sectionNumQuestions[i]),
			Score:        sectionScores[i],
		}
	}

	err = menu.wrapper.ExamController.Insert(newExam)
	if err != nil {
		return errors.New("error creating exam: " + err.Error())
	}
	err = menu.wrapper.ExamSectionController.InsertMany(newExamSection)
	if err != nil {
		return errors.New("error creating exam sections: " + err.Error())
	}
	println("Exam created successfully.")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *ExamMenuState) ListAllExams() error {
	exams, err := menu.wrapper.ExamController.List(nil)
	if err != nil {
		return errors.New("error retrieving exams: " + err.Error())
	}
	menu.PrintExamList(exams)
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *ExamMenuState) RetrieveExamByID() error {
	examID, err := evalUtil.PromptUint("Enter Exam ID to Retrieve: ")
	if err != nil {
		return errors.New("invalid exam ID")
	}
	exam, err := menu.wrapper.ExamController.RetrieveByID(uint(examID))
	if err != nil {
		return errors.New("error retrieving exam: " + err.Error())
	}
	examSections, err := menu.wrapper.ExamSectionController.List(map[string]interface{}{"exam_id": exam.ID})
	if err != nil {
		return errors.New("error retrieving exam sections: " + err.Error())
	}
	fmt.Println("Exam Details:")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Exam ID: %d\n", exam.ID)
	fmt.Printf("Exam Name: %s\n", exam.ExamName)
	fmt.Printf("Instructor ID: %d\n", exam.InstructorID)
	fmt.Printf("Class ID: %d\n", exam.ClassID)
	fmt.Printf("Description: %s\n", exam.Description)
	fmt.Printf("Start Date: %s\n", exam.StartDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("End Date: %s\n", exam.EndDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("Status: %s\n", exam.ExamStatus)
	for _, section := range examSections {
		fmt.Printf("Section ID: %d, Section No: %d, Description: %s\n", section.ID, section.SectionNo, section.Description)
		fmt.Printf("Number of Questions: %d, Score: %.2f\n", section.NumQuestions, section.Score)
		fmt.Println("--------------------------------------------------")
	}
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *ExamMenuState) UpdateExam() error {
	examID, err := evalUtil.PromptUint("Enter Exam ID to update: ")
	if err != nil {
		return errors.New("invalid exam ID")
	}

	exam, err := menu.wrapper.ExamController.RetrieveByID(uint(examID))
	if err != nil {
		return errors.New("error retrieving exam: " + err.Error())
	}

	if exam.ExamStatus != model.Draft {
		return errors.New("exam is not in Draft status, cannot update")
	}

	newExamName := evalUtil.PromptString("Enter Exam Name: ")
	newDescription := evalUtil.PromptString("Enter Exam Description: ")
	newStartDate, err := evalUtil.PromptDate("Enter Exam Start Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		return errors.New("invalid date format")
	}
	newEndDate, err := evalUtil.PromptDate("Enter Exam Start Date (YYYY-MM-DD H:M:S): ")
	if err != nil {
		return errors.New("invalid date format")
	}

	// mock up
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
		return errors.New("start date must be before end date")
	}

	err = menu.wrapper.ExamController.UpdateByID(updatedExam)
	if err != nil {
		return errors.New("error updating exam: " + err.Error())
	}

	fmt.Println("Exam updated successfully.")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *ExamMenuState) UpdateExamSection() error {
	examID, err := evalUtil.PromptUint("Enter Exam ID to Retrieve: ")
	if err != nil {
		return errors.New("invalid exam ID")
	}
	examSections, err := menu.wrapper.ExamSectionController.List(map[string]interface{}{"exam_id": examID})
	if err != nil {
		return errors.New("error retrieving exam sections: " + err.Error())
	}
	if err != nil {
		return errors.New("error retrieving sections: " + err.Error())
	}
	fmt.Printf("Exam Sections of Exam ID [%d]: \n", examID)
	for _, section := range examSections {
		fmt.Printf("Section ID: %d, Section No: %d, Description: %s\n", section.ID, section.SectionNo, section.Description)
		fmt.Printf("Number of Questions: %d, Score: %.2f\n", section.NumQuestions, section.Score)
		fmt.Println("--------------------------------------------------")
	}

	sectionID, err := evalUtil.PromptUint("Enter Section ID to update: ")
	if err != nil {
		return errors.New("invalid section ID")
	}
	for _, section := range examSections {
		if section.ID != uint(sectionID) {
			return errors.New("section ID not found")
		}
	}
	newDescription := evalUtil.PromptString("Enter new description: ")
	newNumQuestions, err := evalUtil.PromptUint("Enter new number of questions: ")
	if err != nil {
		return errors.New("invalid number of questions")
	}
	newScore, err := evalUtil.PromptFloat("Enter new score: ")
	if err != nil {
		return errors.New("invalid score")
	}

	updateSection := &model.ExamSection{
		BaseModel: core.BaseModel{
			Model: gorm.Model{ID: uint(sectionID)},
		},
		ExamID: uint(examID),
		SectionNo: uint(sectionID),
		Description: newDescription,
		NumQuestions: uint(newNumQuestions),
		Score: newScore,
	}
	err = menu.wrapper.ExamSectionController.UpdateByID(updateSection)
	if err != nil {
		return errors.New("error updating section: " + err.Error())
	}
	fmt.Println("Section updated successfully.")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *ExamMenuState) ListAllQuestionsByExamID() error {
	examID, err := evalUtil.PromptUint("Enter Exam ID to retrieve questions: ")
	if err != nil {
		return errors.New("invalid exam ID")
	}
	question, err := menu.wrapper.ExamSectionController.List(map[string]interface{}{"exam_id": examID}, "Questions")
	if err != nil {
		return errors.New("error retrieving questions: " + err.Error())
	}
	if len(question) == 0 {
		fmt.Println("No questions found for this exam.")
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
	}
	fmt.Printf("Questions for Exam ID [%d]: \n", examID)
	fmt.Println("--------------------------------------------------")
	for _, q := range question {
		fmt.Printf("[%d] Question: %s\n", q.Questions[q.ID].ID, q.Questions[q.ID].ActualQuestion)
	}
	fmt.Println("--------------------------------------------------")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
    return nil
}

func (menu *ExamMenuState) DeleteExam() error {
	examID, err := evalUtil.PromptUint("Enter Exam ID to delete: ")
	if err != nil {
		return errors.New("invalid exam ID")
	}

	err = menu.wrapper.ExamController.DeleteByID(uint(examID))
	if err != nil {
		return errors.New("error deleting exam: " + err.Error())
	}

	err = menu.wrapper.ExamSectionController.DeleteByCondition(map[string]interface{}{"exam_id": examID})
	if err != nil {
		return errors.New("error deleting exam sections: " + err.Error())
	}

	fmt.Println("Exam deleted successfully.")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *ExamMenuState) PublishExam() error {
	examID, err := evalUtil.PromptUint("Enter Exam ID to publish: ")
	if err != nil {
		return errors.New("invalid exam ID")
	}

	exam, err := menu.wrapper.ExamController.RetrieveByID(uint(examID))
	if err != nil {
		return errors.New("error retrieving exam: " + err.Error())
	}

	if exam.ExamStatus != model.Draft {
		return errors.New("exam is not in Draft status, cannot publish")
	}

	exam.ExamStatus = model.Publish

	err = menu.wrapper.ExamController.UpdateByID(exam)
	if err != nil {
		return errors.New("error publishing exam: " + err.Error())
	}

	fmt.Println("Exam published successfully.")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *ExamMenuState) HideExam() error {
	examID, err := evalUtil.PromptUint("Enter Exam ID to hide: ")
	if err != nil {
		return errors.New("invalid exam ID")
	}

	exam, err := menu.wrapper.ExamController.RetrieveByID(uint(examID))
	if err != nil {
		return errors.New("error retrieving exam: " + err.Error())
	}

	exam.ExamStatus = model.Hidden

	err = menu.wrapper.ExamController.UpdateByID(exam)
	if err != nil {
		return errors.New("error hiding exam: " + err.Error())
	}

	fmt.Println("Exam hidden successfully.")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}