package handler

//MEP-1007

import (
	assetUtil "ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	assetUtilLocal "ModEd/eval/util"
	"fmt"
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
	menu.handler.AddHandler("4", "Show correct answer by question ID.", handler.FuncStrategy{Action: menu.ShowCorrectAnswerByQuestionID})
	menu.handler.AddHandler("5", "Show question by question ID.", handler.FuncStrategy{Action: menu.ShowQuestionByQuestionID})
	menu.handler.AddHandler("6", "Show all questions.", handler.FuncStrategy{Action: menu.ShowAllQuestions})
	menu.handler.AddHandler("b", "Back to previous menu.", menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *QuestionMenuState) HandleUserInput(input string) error {
	menu.handler.HandleInput(input)
	if input == "back" {
		assetUtil.ClearScreen()
		return nil
	}

	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
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
		expectedAnswer := assetUtilLocal.GetStringInput("Enter expected short answer: ")
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
			label := assetUtilLocal.GetStringInput(fmt.Sprintf("Enter choice #%d text: ", i+1))
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
	return nil
}

func (menu *QuestionMenuState) DeleteQuestion() error {

	questionID := assetUtil.GetUintInput("Enter question ID to delete: ")
	if err := menu.wrapper.QuestionController.DeleteByQuestionID(questionID); err != nil {
		return fmt.Errorf("failed to delete question: %w", err)
	}
	return nil
}

func (menu *QuestionMenuState) UpdateQuestion() error {
	questionID := assetUtil.GetUintInput("Enter question ID to update: ")

	question, err := menu.wrapper.QuestionController.RetrieveByID(questionID)
	if err != nil {
		return fmt.Errorf("failed to retrieve question: %w", err)
	}

	newQuestion := assetUtilLocal.GetStringInput("Enter new question text: ")
	changeAnswer := assetUtil.GetStringInput("Do you want to change the answer? (yes/no): ")

	if changeAnswer == "yes" {
	fmt.Println(`Select question type:
1 = Multiple Choice
2 = Short Answer
3 = True/False`)
	selectType := assetUtil.GetStringInput("Enter choice: ")

	var newQuestionType model.QuestionType
	switch selectType {
	case "1":
		newQuestionType = model.MultipleChoiceQuestion
	case "2":
		newQuestionType = model.ShortAnswerQuestion
	case "3":
		newQuestionType = model.TrueFalseQuestion
	default:
		return fmt.Errorf("invalid selection for QuestionType")
	}

	question.ActualQuestion = newQuestion

	if question.QuestionType != newQuestionType {
		if err := menu.wrapper.QuestionController.UpdateQuestionType(newQuestionType, question); err != nil {
			return err
		}
	}

	if err := menu.wrapper.QuestionController.UpdateByID(question); err != nil {
		return fmt.Errorf("failed to update question: %w", err)
	}

	fmt.Printf("Question updated with ID: %d\n", question.ID)
	switch newQuestionType {
	case model.ShortAnswerQuestion:
		expectedAnswer := assetUtilLocal.GetStringInput("Enter expected short answer: ")
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
			label := assetUtilLocal.GetStringInput(fmt.Sprintf("Enter choice #%d text: ", i+1))
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
}
	return nil
}

func (menu *QuestionMenuState) ShowCorrectAnswerByQuestionID() error {
	questionID := assetUtil.GetUintInput("Enter question ID to view correct answer: ")

	question, err := menu.wrapper.QuestionController.RetrieveByID(questionID)
	if err != nil {
		return fmt.Errorf("failed to retrieve question: %w", err)
	}
	question_type := question.QuestionType
	if question_type == model.MultipleChoiceQuestion {
		ans, _ := menu.wrapper.QuestionController.RetrieveByID(questionID, "MultipleChoiceAnswers")
		for _, result := range ans.MultipleChoiceAnswers {
			if result.IsExpected {
				fmt.Printf("Correct answer for question %s: %s\n", question.ActualQuestion, result.AnswerLabel)
			}
		}

	} else if question_type == model.ShortAnswerQuestion {
		ans, _ := menu.wrapper.QuestionController.RetrieveByID(questionID, "ShortAnswerQuestion")
		fmt.Printf("Correct answer for question %s: %s\n", question.ActualQuestion, ans.ShortAnswer.ExpectedAnswer)
	} else if question_type == model.TrueFalseQuestion {
		ans, _ := menu.wrapper.QuestionController.RetrieveByID(questionID, "TruefalseAnswer")
		getanswer := "True"
		if !ans.TruefalseAnswer.IsExpected {
			getanswer = "False"
		}
		fmt.Printf("Correct answer for question %s: %s\n", question.ActualQuestion, getanswer)
	}
	return nil

}

func (menu *QuestionMenuState) ShowQuestionByQuestionID() error {
	questionID := assetUtil.GetUintInput("Enter question ID: ")

	question, err := menu.wrapper.QuestionController.RetrieveByID(questionID)
	if err != nil {
		return fmt.Errorf("failed to retrieve question: %w", err)
	}

	fmt.Printf("Question ID: %d\nQuestion: %s\nQuestionType: %s\n", question.ID, question.ActualQuestion, question.QuestionType)

	switch question.QuestionType {
	case model.MultipleChoiceQuestion:
		answers, err := menu.wrapper.MultipleChoiceAnswerController.List(map[string]interface{}{"question_id": question.ID})
		if err != nil {
			return fmt.Errorf("failed to retrieve multiple choice answers: %w", err)
		}
		for _, answer := range answers {
			fmt.Printf("- Label: %s | IsCorrect: %v\n", answer.AnswerLabel, answer.IsExpected)
		}

	case model.ShortAnswerQuestion:
		answer, err := menu.wrapper.ShortAnswerController.RetrieveByCondition(map[string]interface{}{"question_id": question.ID})
		if err != nil {
			return fmt.Errorf("failed to retrieve short answer: %w", err)
		}
		fmt.Printf("- Expected Answer: %s\n", answer.ExpectedAnswer)

	case model.TrueFalseQuestion:
		answer, err := menu.wrapper.TrueFalseAnswerController.RetrieveByCondition(map[string]interface{}{"question_id": question.ID})
		if err != nil {
			return fmt.Errorf("failed to retrieve true/false answer: %w", err)
		}
		fmt.Printf("- Correct Answer: %v\n", answer.IsExpected)
	}
	return nil
}


func (menu *QuestionMenuState) ShowAllQuestions() error {

	questions, err := menu.wrapper.QuestionController.List(map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("failed to list all  questions: %w", err)
	}

	for _, question := range questions {
		fmt.Printf("Question ID: %d Question: %s\n", question.ID, question.ActualQuestion)
	}
	if err != nil {
		return fmt.Errorf("failed to retrieve answers: %w", err)
	}
	return nil
}
