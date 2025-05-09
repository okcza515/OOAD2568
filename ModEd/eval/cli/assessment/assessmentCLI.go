// Assessment CLI Module
package assessment

import (
	"ModEd/eval/cli/assessment/handler"
	"ModEd/eval/controller"
	"fmt"

	"gorm.io/gorm"
)

func RunAssessmentModuleCLI(db *gorm.DB) {
	// Initialize controllers
	assessmentController := controller.NewAssessmentController()
	submissionController := controller.NewSubmissionController(db)
	resultController := controller.NewResultController(db)

	params := &handler.AssessmentCLIParams{
		AssessmentController: assessmentController,
		SubmissionController: submissionController,
		ResultController:     resultController,
	}

	mainState := handler.NewMainMenuState(params)
	stateManager := handler.NewMenuStateManager(mainState)

	// Run menu state manager
	err := stateManager.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func submissionMenu(db *gorm.DB, submissionController *controller.SubmissionController) {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Submission Management")

	ctx.AddHandler("1", "List Submissions", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Submissions =====")
			// Implementation would call submissionController methods
			fmt.Println("Submission listing will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("2", "Record Submission", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Record Submission =====")
			// Implementation would call submissionController methods
			fmt.Println("Submission recording will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("3", "Update Submission", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Update Submission =====")
			// Implementation would call submissionController methods
			fmt.Println("Submission update will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("4", "Delete Submission", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Delete Submission =====")
			// Implementation would call submissionController methods
			fmt.Println("Submission deletion will be implemented here")
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

func resultMenu(db *gorm.DB, resultController *controller.ResultController) {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Result Management")

	ctx.AddHandler("1", "List Results", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Results =====")
			// Implementation would call resultController methods
			fmt.Println("Result listing will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("2", "Record Result", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Record Result =====")
			// Implementation would call resultController methods
			fmt.Println("Result recording will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("3", "Update Result", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Update Result =====")
			// Implementation would call resultController methods
			fmt.Println("Result update will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("4", "Delete Result", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Delete Result =====")
			// Implementation would call resultController methods
			fmt.Println("Result deletion will be implemented here")
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
