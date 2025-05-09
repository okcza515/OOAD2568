package progress

import (
	"ModEd/eval/cli/progress/handler"
	"ModEd/eval/controller"
	"fmt"

	"gorm.io/gorm"
)

func RunProgressModuleCLI(db *gorm.DB) {
	// Initialize controllers
	progressController := controller.NewProgressController(db)

	params := &handler.ProgressCLIParams{
		ProgressController: progressController,
	}

	mainState := handler.NewMainMenuState(params)
	stateManager := handler.NewMenuStateManager(mainState)

	// Run menu state manager
	err := stateManager.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
