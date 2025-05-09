package evaluation

import (
	"ModEd/core/cli"
	"ModEd/eval/cli/evaluation/handler"
	"fmt"

	"gorm.io/gorm"
)

func RunEvaluationModuleCLI(db *gorm.DB) {
	// Create CLI state manager
	manager := cli.NewCLIMenuManager()

	// Create evaluation module handler
	evalHandler := handler.NewEvalModuleHandler(manager, db)

	// Set initial state
	manager.SetState(evalHandler)

	// Run menu state manager
	for {
		evalHandler.Render()
		var input string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		err := evalHandler.HandleUserInput(input)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
