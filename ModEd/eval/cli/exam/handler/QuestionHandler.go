package handler

//MEP-1007

import (
	assetUtil "ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	assetUtilLocal "ModEd/eval/util"
	// "errors"
	"fmt"
	// "strconv"
	// "time"
)

type QuestionMenuState struct {
	Manager                        *cli.CLIMenuStateManager
	wrapper                        *controller.ExamModuleWrapper
	QuestionModuleMenuStateHandler cli.MenuState
	handler                        *handler.HandlerContext
	backhandler                    *handler.ChangeMenuHandlerStrategy
}

func NewQuestionMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.ExamModuleWrapper, questionModuleMenuStateHandler cli.MenuState) *QuestionMenuState {
	return &QuestionMenuState{
		Manager:                        manager,
		wrapper:                        wrapper,
		QuestionModuleMenuStateHandler: questionModuleMenuStateHandler,
		handler:                        handler.NewHandlerContext(),
		backhandler:                    handler.NewChangeMenuHandlerStrategy(manager, questionModuleMenuStateHandler),
	}
}

func (menu *QuestionMenuState) Render() {
	menu.handler.SetMenuTitle("\nQuestion management menu:")
	menu.handler.AddHandler("1", "Create a new question.", handler.FuncStrategy{Action: menu.CreateQuestion})
	menu.handler.AddHandler("2", "Delete a question.", handler.FuncStrategy{Action: menu.DeleteQuestion})
	menu.handler.AddHandler("3", "Uadete questions.", handler.FuncStrategy{Action: menu.UpdateQuestion})
	menu.handler.AddHandler("4", "Show correct answer By ID.", handler.FuncStrategy{Action: menu.ShowCorrectAnswerByQyestionID})
	menu.handler.AddHandler("5", "Show question by section ID.", handler.FuncStrategy{Action: menu.ShowQuestionBySectionID})
	menu.handler.AddHandler("b", "Back to previous menu.", menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *QuestionMenuState) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *QuestionMenuState) CreateQuestion() error {
	sectionID := assetUtil.GetUintInput("Enter section ID: ")
	score := assetUtil.GetFloatInput("Enter examination score: ")
	actualQuestion := assetUtilLocal.GetStringInput("Enter the question text: ")

	fmt.Println(`Select question type:
	1 = Multiple Choice
	2 = Short Answer
	3 = True/False`)
	selectType := assetUtil.GetStringInput("Enter choice: ")
	var questionType model.QuestionType
	switch selectType {
	case "1":
		questionType = model.MultipleChoiceQuestion
	case "2":
		questionType = model.ShortAnswerQuestion
	case "3":
		questionType = model.TrueFalseQuestion
	default:
		return fmt.Errorf("invalid selection for QuestionType")
	}

	fmt.Println("You selected:", sectionID, score, actualQuestion, questionType)
	question := &model.Question{
		SectionID:      sectionID,
		Score:          score,
		ActualQuestion: actualQuestion,
		QuestionType:   questionType,
	}

	if err := menu.wrapper.QuestionController.Insert(question); err != nil {
		return fmt.Errorf("failed to insert question: %w", err)
	}
	fmt.Printf("Question created with ID: %d\n", question.ID)
	switch questionType {
	case model.ShortAnswerQuestion:
		expectedAnswer := assetUtil.GetStringInput("Enter expected short answer: ")
		answer := &model.ShortAnswer{
			QuestionID:     question.ID,
			ExpectedAnswer: expectedAnswer,
		}
		if err := menu.wrapper.ShortAnswerController.Insert(answer); err != nil {
			return fmt.Errorf("failed to insert short answer: %w", err)
		}

	case model.TrueFalseQuestion:
		expectedStr := assetUtil.GetStringInput("Enter expected answer (true/false): ")
		isExpected := expectedStr == "true"
		answer := &model.TrueFalseAnswer{
			QuestionID: question.ID,
			IsExpected: isExpected,
		}
		if err := menu.wrapper.TrueFalseAnswerController.Insert(answer); err != nil {
			return fmt.Errorf("failed to insert true/false answer: %w", err)
		}

	case model.MultipleChoiceQuestion:
		numChoices := assetUtil.GetUintInput("How many choices? (e.g. 4): ")

		for i := 0; i < int(numChoices); i++ {
			label := assetUtil.GetStringInput(fmt.Sprintf("Enter choice #%d text: ", i+1))
			isExpectedStr := assetUtil.GetStringInput(fmt.Sprintf("Is this the correct answer for choice #%d? (true/false): ", i+1))
			isExpected := isExpectedStr == "true"

			answer := &model.MultipleChoiceAnswer{
				QuestionID:  question.ID,
				AnswerLabel: label,
				IsExpected:  isExpected,
			}
			if err := menu.wrapper.MultipleChoiceAnswerController.Insert(answer); err != nil {
				return fmt.Errorf("failed to insert multiple choice answer: %w", err)
			}
		}
	}

	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *QuestionMenuState) DeleteQuestion() error {

	questionID := assetUtil.GetUintInput("Enter question ID to delete: ")
	if err := menu.wrapper.QuestionController.DeleteByID(questionID); err != nil {
		return fmt.Errorf("failed to delete question: %w", err)
	}
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *QuestionMenuState) UpdateQuestion() error {

	questionID := assetUtil.GetUintInput("Enter question ID to update: ")

	question, err := menu.wrapper.QuestionController.RetrieveByID(questionID)
	if err != nil {
		return fmt.Errorf("failed to retrieve question: %w", err)
	}
	newQuestion := assetUtilLocal.GetStringInput("Enter new question text: ")

	fmt.Println(`Select question type:
1 = Multiple Choice
2 = Short Answer
3 = True/False`)
	selectType := assetUtil.GetStringInput("Enter choice: ")
	var questionType model.QuestionType
	switch selectType {
	case "1":
		questionType = model.MultipleChoiceQuestion
	case "2":
		questionType = model.ShortAnswerQuestion
	case "3":
		questionType = model.TrueFalseQuestion
	default:
		return fmt.Errorf("invalid selection for QuestionType")
	}

	question.ActualQuestion = newQuestion
	question.QuestionType = questionType

	if err := menu.wrapper.QuestionController.UpdateByID(question); err != nil {
		return fmt.Errorf("failed to update question: %w", err)
	}
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *QuestionMenuState) ShowCorrectAnswerByQyestionID() error {
	questionIDRaw := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter question ID to view correct answer:",
		FieldNameText: "QuestionID",
	})
	questionID, ok := questionIDRaw.(uint)
	if !ok {
		return fmt.Errorf("invalid input for QuestionID")
	}

	question, err := menu.wrapper.QuestionController.RetrieveByID(questionID)
	if err != nil {
		return fmt.Errorf("failed to retrieve question: %w", err)
	}
	question_type := question.QuestionType
	if question_type == "MultiplechoiseQuestion" {
		ans, _ := menu.wrapper.QuestionController.RetrieveByID(questionID, "MultipleChoiceAnswers")
		for _, result := range ans.MultipleChoiceAnswers {
			if result.IsExpected {
				fmt.Printf("Correct answer for question %s: %s\n", question.ActualQuestion, result.AnswerLabel)
			}
		}

	} else if question_type == "ShortAnswerQuestion" {
		ans, _ := menu.wrapper.QuestionController.RetrieveByID(questionID, "ShortAnswerQuestion")
		fmt.Printf("Correct answer for question %s: %s\n", question.ActualQuestion, ans.ShortAnswer.ExpectedAnswer)
	} else if question_type == "TrueFalseAnswerQuestion" {
		ans, _ := menu.wrapper.QuestionController.RetrieveByID(questionID, "TruefalseAnswer")
		getanswer := "True"
		if !ans.TruefalseAnswer.IsExpected {
			getanswer = "False"
		}
		fmt.Printf("Correct answer for question %s: %s\n", question.ActualQuestion, getanswer)
	}
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil

}

func (menu *QuestionMenuState) ShowQuestionBySectionID() error {
	sectionID := assetUtil.GetUintInput("Enter section ID: ")

	questions, err := menu.wrapper.QuestionController.RetrieveByID(sectionID)
	if err != nil {
		return fmt.Errorf("failed to retrieve questions: %w", err)
	}

	fmt.Printf("Question ID: %d, Question: %s QuestionType %s\n", questions.ID, questions.ActualQuestion, questions.QuestionType)
	Answer, err := menu.wrapper.MultipleChoiceAnswerController.List(map[string]interface{}{
		"question_id": questions.ID})

	for _, answer := range Answer {
		fmt.Printf("Answer Label: %s\n", answer.AnswerLabel)
	}
	if err != nil {
		return fmt.Errorf("failed to retrieve answers: %w", err)
	}
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}
