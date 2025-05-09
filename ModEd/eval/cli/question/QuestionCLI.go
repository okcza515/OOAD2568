package question

import (
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"fmt"
)

type QuestionCLIParams struct {
	QuestionController *controller.QuestionController
	AnswerController   *controller.AnswerController
}

func RunQuestionCLI(params *QuestionCLIParams) {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Question Management")

	// Question menu items
	ctx.AddHandler("1", "List Questions", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Questions =====")
			// Implementation would call params.QuestionController methods
			fmt.Println("Question listing will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("2", "Create Question", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Create Question =====")
			// Implementation would call params.QuestionController methods
			fmt.Println("Question creation will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("3", "Update Question", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Update Question =====")
			// Implementation would call params.QuestionController methods
			fmt.Println("Question update will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("4", "Delete Question", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Delete Question =====")
			// Implementation would call params.QuestionController methods
			fmt.Println("Question deletion will be implemented here")
			return nil
		},
	})

	// Answers menu items
	ctx.AddHandler("5", "Manage Answers", handler.FuncStrategy{
		Action: func() error {
			manageQuestionAnswers(params)
			return nil
		},
	})

	ctx.AddBackHandler(handler.FuncStrategy{
		Action: func() error {
			return nil
		},
	})

	for {
		ctx.ShowMenu()
		var userInput string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&userInput)
		if err := ctx.HandleInput(userInput); err != nil {
			fmt.Println("Error:", err)
		}
		if userInput == "back" {
			break
		}
	}
}

func manageQuestionAnswers(params *QuestionCLIParams) {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Question Answers Management")

	ctx.AddHandler("1", "List Answers", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Question Answers =====")
			// Implementation would call params.AnswerController methods
			fmt.Println("Answer listing will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("2", "Add Answer", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Add Answer =====")
			// Implementation would call params.AnswerController methods
			fmt.Println("Answer addition will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("3", "Update Answer", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Update Answer =====")
			// Implementation would call params.AnswerController methods
			fmt.Println("Answer update will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("4", "Delete Answer", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Delete Answer =====")
			// Implementation would call params.AnswerController methods
			fmt.Println("Answer deletion will be implemented here")
			return nil
		},
	})

	ctx.AddBackHandler(handler.FuncStrategy{
		Action: func() error {
			return nil
		},
	})

	for {
		ctx.ShowMenu()
		var userInput string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&userInput)
		if err := ctx.HandleInput(userInput); err != nil {
			fmt.Println("Error:", err)
		}
		if userInput == "back" {
			break
		}
	}
}
