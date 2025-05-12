package handler

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"errors"
	"fmt"
	"time"
)

type QuizMenuStateHandler struct {
	Manager                    *cli.CLIMenuStateManager
	wrapper                    *controller.EvalModuleWrapper
	EvalModuleMenuStateHandler cli.MenuState
	handler                    *handler.HandlerContext
	backhandler                *handler.ChangeMenuHandlerStrategy
}

func NewQuizMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.EvalModuleWrapper, evalModuleMenuStateHandler cli.MenuState) *QuizMenuStateHandler {
	return &QuizMenuStateHandler{
		Manager:                    manager,
		wrapper:                    wrapper,
		EvalModuleMenuStateHandler: evalModuleMenuStateHandler,
		handler:                    handler.NewHandlerContext(),
		backhandler:                handler.NewChangeMenuHandlerStrategy(manager, evalModuleMenuStateHandler),
	}
}

func (menu *QuizMenuStateHandler) Render() {
	menu.handler.SetMenuTitle("\nQuiz management menu:")
	menu.handler.AddHandler("1", "Create new quiz", handler.FuncStrategy{Action: menu.CreateQuiz})
	menu.handler.AddHandler("2", "Update quiz", handler.FuncStrategy{Action: menu.UpdateQuiz})
	menu.handler.AddHandler("3", "Delete quiz", handler.FuncStrategy{Action: menu.DeleteQuiz})
	menu.handler.AddHandler("4", "Get quiz by ID", handler.FuncStrategy{Action: menu.GetQuizByID})
	menu.handler.AddHandler("5", "List all quizzes", handler.FuncStrategy{Action: menu.ListAllQuizzes})
	menu.handler.AddBackHandler(menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *QuizMenuStateHandler) HandlerUserInput(input string) error {
	err := menu.handler.HandleInput(input)
	if err != nil {
		return err
	}
	return nil
}

func (menu *QuizMenuStateHandler) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *QuizMenuStateHandler) printQuizTable(quizzes []*model.Quiz) {
	if len(quizzes) == 0 {
		fmt.Println("\nNo quizzes found.")
		return
	}

	fmt.Printf("\n%-5s %-20s %-15s %-15s %-10s %-10s", "ID", "Title", "Status", "Start Date", "End Date", "Attempts")
	fmt.Printf("\n%-5s %-20s %-15s %-15s %-10s %-10s", "---", "-----", "------", "----------", "--------", "--------")

	for _, quiz := range quizzes {
		fmt.Printf("\n%-5d %-20s %-15s %-15s %-10s %-10d",
			quiz.ID,
			quiz.Title,
			quiz.Status,
			quiz.StartDate.Format("2006-01-02"),
			quiz.EndDate.Format("2006-01-02"),
			quiz.Attempts)
	}
	fmt.Println()
}

func (menu *QuizMenuStateHandler) CreateQuiz() error {
	title := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Quiz Title: ",
		FieldNameText: "Title",
	}).(string)

	description := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Quiz Description: ",
		FieldNameText: "Description",
	}).(string)

	status := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Quiz Status (active/hidden): ",
		FieldNameText: "Status",
	}).(string)

	startDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Start Date (YYYY-MM-DD): ",
		FieldNameText: "Start Date",
	}).(string)
	startDate, _ := time.Parse("2006-01-02", startDateStr)

	endDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter End Date (YYYY-MM-DD): ",
		FieldNameText: "End Date",
	}).(string)
	endDate, _ := time.Parse("2006-01-02", endDateStr)

	attempts := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Number of Attempts: ",
		FieldNameText: "Attempts",
	}).(uint)

	instructorID := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Instructor ID: ",
		FieldNameText: "Instructor ID",
	}).(uint)

	courseID := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Course ID: ",
		FieldNameText: "Course ID",
	}).(uint)

	quiz := &model.Quiz{
		Title:        title,
		Description:  description,
		Status:       status,
		StartDate:    startDate,
		EndDate:      endDate,
		Attempts:     attempts,
		InstructorID: instructorID,
		CourseID:     courseID,
	}

	createdQuiz, err := menu.wrapper.QuizController.CreateQuiz(quiz)
	if err != nil {
		return errors.New("failed to create quiz")
	}

	fmt.Printf("\nQuiz created successfully with ID: %d", createdQuiz.ID)
	return nil
}

func (menu *QuizMenuStateHandler) UpdateQuiz() error {
	quizID := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Quiz ID to update: ",
		FieldNameText: "Quiz ID",
	}).(uint)

	existingQuiz, err := menu.wrapper.QuizController.GetQuiz(quizID)
	if err != nil {
		return errors.New("failed to retrieve quiz")
	}

	fmt.Printf("Current title: %s\n", existingQuiz.Title)
	title := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new Quiz Title (or press Enter to keep current): ",
		FieldNameText: "Title",
	}).(string)
	if title == "" {
		title = existingQuiz.Title
	}

	fmt.Printf("Current description: %s\n", existingQuiz.Description)
	description := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new Quiz Description (or press Enter to keep current): ",
		FieldNameText: "Description",
	}).(string)
	if description == "" {
		description = existingQuiz.Description
	}

	fmt.Printf("Current status: %s\n", existingQuiz.Status)
	status := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new Quiz Status (active/hidden) (or press Enter to keep current): ",
		FieldNameText: "Status",
	}).(string)
	if status == "" {
		status = existingQuiz.Status
	}

	fmt.Printf("Current start date: %s\n", existingQuiz.StartDate.Format("2006-01-02"))
	startDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new Start Date (YYYY-MM-DD) (or press Enter to keep current): ",
		FieldNameText: "Start Date",
	}).(string)
	startDate := existingQuiz.StartDate
	if startDateStr != "" {
		startDate, _ = time.Parse("2006-01-02", startDateStr)
	}

	fmt.Printf("Current end date: %s\n", existingQuiz.EndDate.Format("2006-01-02"))
	endDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new End Date (YYYY-MM-DD) (or press Enter to keep current): ",
		FieldNameText: "End Date",
	}).(string)
	endDate := existingQuiz.EndDate
	if endDateStr != "" {
		endDate, _ = time.Parse("2006-01-02", endDateStr)
	}

	fmt.Printf("Current attempts: %d\n", existingQuiz.Attempts)
	attemptsStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new Number of Attempts (or press Enter to keep current): ",
		FieldNameText: "Attempts",
	}).(string)
	attempts := existingQuiz.Attempts
	if attemptsStr != "" {
		attemptsParsed, ok := core.UintInputStep{}.Validate(attemptsStr)
		if ok {
			attempts = attemptsParsed.(uint)
		}
	}

	quiz := &model.Quiz{
		Title:        title,
		Description:  description,
		Status:       status,
		StartDate:    startDate,
		EndDate:      endDate,
		Attempts:     attempts,
		InstructorID: existingQuiz.InstructorID,
		CourseID:     existingQuiz.CourseID,
	}

	updatedQuiz, err := menu.wrapper.QuizController.UpdateQuiz(quiz)
	if err != nil {
		return errors.New("failed to update quiz")
	}

	fmt.Printf("\nQuiz updated successfully with ID: %d", updatedQuiz.ID)
	return nil
}

func (menu *QuizMenuStateHandler) DeleteQuiz() error {
	quizID := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Quiz ID to delete: ",
		FieldNameText: "Quiz ID",
	}).(uint)

	err := menu.wrapper.QuizController.DeleteQuiz(quizID)
	if err != nil {
		return errors.New("failed to delete quiz")
	}

	fmt.Printf("\nQuiz with ID %d has been marked as hidden", quizID)
	return nil
}

func (menu *QuizMenuStateHandler) GetQuizByID() error {
	quizID := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Quiz ID: ",
		FieldNameText: "Quiz ID",
	}).(uint)

	quiz, err := menu.wrapper.QuizController.GetQuiz(quizID)
	if err != nil {
		return errors.New("failed to retrieve quiz")
	}

	fmt.Printf("\nQuiz Details:")
	fmt.Printf("\nID: %d", quiz.ID)
	fmt.Printf("\nTitle: %s", quiz.Title)
	fmt.Printf("\nDescription: %s", quiz.Description)
	fmt.Printf("\nStatus: %s", quiz.Status)
	fmt.Printf("\nStart Date: %s", quiz.StartDate)
	fmt.Printf("\nEnd Date: %s", quiz.EndDate)
	fmt.Printf("\nAttempts: %d", quiz.Attempts)
	fmt.Printf("\nInstructor ID: %d", quiz.InstructorID)
	fmt.Printf("\nCourse ID: %d", quiz.CourseID)
	fmt.Println()
	return nil
}

func (menu *QuizMenuStateHandler) ListAllQuizzes() error {
	quizzes, err := menu.wrapper.QuizController.GetAllQuizzes()
	if err != nil {
		return errors.New("failed to retrieve quizzes")
	}

	fmt.Println("\nAll Quizzes:")
	menu.printQuizTable(quizzes)
	return nil
}
