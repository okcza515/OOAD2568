// Quiz CLI Module
package quiz

import (
	"ModEd/eval/cli/quiz/handler"
	"ModEd/eval/controller"
	"fmt"

	"gorm.io/gorm"
)

func RunQuizModuleCLI(db *gorm.DB) {
	// Initialize controllers
	quizController := controller.NewQuizController(db)
	questionController := controller.NewQuestionController(db)

	params := &handler.QuizCLIParams{
		QuizController:     quizController,
		QuestionController: questionController,
	}

	mainState := handler.NewMainMenuState(params)
	stateManager := handler.NewMenuStateManager(mainState)

	// Run menu state manager
	err := stateManager.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func manageQuizQuestions(db *gorm.DB, questionController *controller.QuestionController) {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Quiz Questions Management")

	ctx.AddHandler("1", "List Questions", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Quiz Questions =====")
			// Implementation would call questionController methods
			fmt.Println("Question listing will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("2", "Add Question", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Add Question =====")
			// Implementation would call questionController methods
			fmt.Println("Question addition will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("3", "Update Question", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Update Question =====")
			// Implementation would call questionController methods
			fmt.Println("Question update will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("4", "Delete Question", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Delete Question =====")
			// Implementation would call questionController methods
			fmt.Println("Question deletion will be implemented here")
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
