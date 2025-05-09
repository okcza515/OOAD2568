package handler

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	progress "ModEd/eval/cli/progress/handler"
	"ModEd/eval/controller"

	"gorm.io/gorm"
)

type EvalModuleMenuStateHandler struct {
	Manager                  *cli.CLIMenuStateManager
	wrapper                  *controller.EvalModuleWrapper
	handlerContext           *handler.HandlerContext
	ProgressMenuStateHandler *progress.ProgressMenuStateHandler
}

func NewEvalModuleHandler(manager *cli.CLIMenuStateManager, db *gorm.DB) *EvalModuleMenuStateHandler {
	wrapper := controller.NewEvalModuleWrapper(db)
	handlerContext := handler.NewHandlerContext()

	evalModuleHandler := &EvalModuleMenuStateHandler{
		Manager:        manager,
		wrapper:        wrapper,
		handlerContext: handlerContext,
	}

	evalModuleHandler.ProgressMenuStateHandler = progress.NewProgressMenuStateHandler(manager, wrapper, evalModuleHandler)

	handlerContext.SetMenuTitle("Evaluation Module Menu")
	handlerContext.AddHandler("1", "Progress", handler.FuncStrategy{
		Action: func() error {
			manager.SetState(evalModuleHandler.ProgressMenuStateHandler)
			return nil
		},
	})
	handlerContext.AddHandler("exit", "Exit the Evaluation Module", handler.FuncStrategy{
		Action: func() error {
			return nil
		},
	})

	return evalModuleHandler
}

func (handler *EvalModuleMenuStateHandler) Render() {
	handler.handlerContext.ShowMenu()
}

func (handler *EvalModuleMenuStateHandler) HandleUserInput(input string) error {
	return handler.handlerContext.HandleInput(input)
}
