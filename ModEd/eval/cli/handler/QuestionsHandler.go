package handler

//MEP-1007

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"

	// "errors"
	"fmt"
	// "strconv"
	// "time"
)

type QuestionMenuStateHandler struct {
	Manager                        *cli.CLIMenuStateManager
	wrapper                        *controller.ExamModuleWrapper
	QuestionModuleMenuStateHandler cli.MenuState
	handler                        *handler.HandlerContext
	backhandler                    *handler.ChangeMenuHandlerStrategy
}

func NewQuestionMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.ExamModuleWrapper, questionModuleMenuStateHandler cli.MenuState) *QuestionMenuStateHandler {
	return &QuestionMenuStateHandler{
		Manager:                        manager,
		wrapper:                        wrapper,
		QuestionModuleMenuStateHandler: questionModuleMenuStateHandler,
		handler:                        handler.NewHandlerContext(),
		backhandler:                    handler.NewChangeMenuHandlerStrategy(manager, questionModuleMenuStateHandler),
	}
}

func (menu *QuestionMenuStateHandler) Render() {
	menu.handler.SetMenuTitle("\nQuestion management menu:")
	menu.handler.AddHandler("1", "Create a new question.", handler.FuncStrategy{Action: menu.CreateQuestion})
	menu.handler.AddHandler("2", "Delete a question.", handler.FuncStrategy{Action: menu.DeleteQuestion})
	menu.handler.AddHandler("3", "Uadete questions.", handler.FuncStrategy{Action: menu.UpdateQuestion})
	menu.handler.AddHandler("4", "Show correct answer By ID.", handler.FuncStrategy{Action: menu.ShowCorrectAnswerByQyestionID})
	menu.handler.AddHandler("b", "Back to previous menu.", menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *QuestionMenuStateHandler) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *QuestionMenuStateHandler) CreateQuestion() error {
	sectionIDRaw := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter examination section_id:",
		FieldNameText: "SectionID",
	})
	sectionID, ok := sectionIDRaw.(uint)
	if !ok {
		return fmt.Errorf("invalid input for SectionID")
	}

	scoreRaw := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter examination score:",
		FieldNameText: "Score",
	})
	score, ok := scoreRaw.(float64)
	if !ok {
		return fmt.Errorf("invalid input for Score")
	}

	actualQuestionRaw := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter examination actual_question:",
		FieldNameText: "ActualQuestion",
	})
	actualQuestion, ok := actualQuestionRaw.(string)
	if !ok {
		return fmt.Errorf("invalid input for ActualQuestion")
	}

	questionTypeRaw := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter examination question_type:",
		FieldNameText: "QuestionType",
	})
	questionType, ok := questionTypeRaw.(string)
	if !ok {
		return fmt.Errorf("invalid input for QuestionType")
	}

	question := &model.Question{
		SectionID:      sectionID,
		Score:          score,
		ActualQuestion: actualQuestion,
		QuestionType:   model.QuestionType(questionType),
	}

	if err := menu.wrapper.QuestionController.Insert(question); err != nil {
		return fmt.Errorf("failed to insert question: %w", err)
	}

	return nil
}

func (menu *QuestionMenuStateHandler) DeleteQuestion() error {
	questionIDRaw := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter question ID to delete:",
		FieldNameText: "QuestionID",
	})
	questionID, ok := questionIDRaw.(uint)
	if !ok {
		return fmt.Errorf("invalid input for QuestionID")
	}

	if err := menu.wrapper.QuestionController.DeleteByID(questionID); err != nil {
		return fmt.Errorf("failed to delete question: %w", err)
	}

	return nil
}

func (menu *QuestionMenuStateHandler) UpdateQuestion() error {
	questionIDRaw := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter question ID to update:",
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

	newQuestionTypeRaw := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new question type:",
		FieldNameText: "NewQuestionType",
	})
	newQuestionType, ok := newQuestionTypeRaw.(string)
	if !ok {
		return fmt.Errorf("invalid input for NewQuestionType")
	}

	question.QuestionType = model.QuestionType(newQuestionType)

	if err := menu.wrapper.QuestionController.UpdateByID(question); err != nil {
		return fmt.Errorf("failed to update question: %w", err)
	}

	return nil
}

func (menu *QuestionMenuStateHandler) ShowCorrectAnswerByQyestionID() error {
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
			if result.IsExpected == true {
				fmt.Printf("Correct answer for question %s: %s\n", question.ActualQuestion, result.AnswerLabel)
			}
		}

	} else if question_type == "ShortAnswerQuestion" {
		ans, _ := menu.wrapper.QuestionController.RetrieveByID(questionID, "ShortAnswerQuestion")
		fmt.Printf("Correct answer for question %s: %s\n", question.ActualQuestion, ans.ShortAnswer.ExpectedAnswer)
	} else if question_type == "TrueFalseAnswerQuestion" {
		ans, _ := menu.wrapper.QuestionController.RetrieveByID(questionID, "TruefalseAnswer")
		getanswer := "True"
		if ans.TruefalseAnswer.IsExpected == false {
			getanswer = "False"
		}
		fmt.Printf("Correct answer for question %s: %s\n", question.ActualQuestion, getanswer)
	}

	return nil

}
