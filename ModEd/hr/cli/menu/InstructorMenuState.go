package menu

import (
	"ModEd/core/cli"
	coreHandler "ModEd/core/handler"
	hrHandler "ModEd/hr/cli/menu/handler"
	"ModEd/hr/controller"
	"fmt"
)

type InstructorMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *coreHandler.HandlerContext
}

func (a *InstructorMenuState) HandleUserInput(input string) error {
	err := a.handlerContext.HandleInput(input)
	if err != nil {
		return err
	}

	return nil
}

func (a *InstructorMenuState) Render() {
	fmt.Println("=== Instructor Menu ===")
	a.handlerContext.ShowMenu()
	fmt.Println("exit:\tExit the program.")
	fmt.Println()
}

func NewInstructorMenuState(
	cliManager *cli.CLIMenuStateManager,
	hrManager *controller.HRControllerManager,
) *InstructorMenuState {
	handlerContext := coreHandler.NewHandlerContext()

	addInstructorHandler := hrHandler.NewAddInstructorStrategy(hrManager.InstructorCtrl)
	listInstructorHandler := hrHandler.NewListInstructorStrategy(hrManager.InstructorCtrl)
	updateInstructorHandler := hrHandler.NewUpdateInstructorInfoStrategy(hrManager.InstructorCtrl)
	requestInstructorLeaveHandler := hrHandler.NewRequestLeaveHandlerStrategy(hrManager.LeaveInstructorCtrl.SubmitInstructorLeaveRequest)
	requestInstructorResignHandler := hrHandler.NewRequestResignationHandlerStrategy(hrManager.ResignInstructorCtrl.SubmitResignationInstructor)
	requestInstructorRaiseHandler := hrHandler.NewRequestRaiseHandlerStrategy(hrManager.RaiseCtrl)
	reviewLeaveHandler := hrHandler.NewReviewHandlerStrategy(hrManager.LeaveInstructorCtrl.ReviewInstructorLeaveRequest)
	reviewResignationHandler := hrHandler.NewReviewHandlerStrategy(hrManager.ResignInstructorCtrl.ReviewInstructorResignRequest)
	reviewRaiseHandler := hrHandler.NewReviewHandlerStrategy(hrManager.RaiseCtrl.ReviewInstructorRaiseRequest)
	deleteInstructorHandler := hrHandler.NewDeleteInstructorStrategy(hrManager.InstructorCtrl)

	handlerContext.AddHandler("1", "Add new instructor", addInstructorHandler)
	handlerContext.AddHandler("2", "List instructor", listInstructorHandler)
	handlerContext.AddHandler("3", "Update instructor Info", updateInstructorHandler)
	handlerContext.AddHandler("4", "Delete instructor", deleteInstructorHandler)
	handlerContext.AddHandler("5", "Request leave", requestInstructorLeaveHandler)
	handlerContext.AddHandler("6", "Request resignation", requestInstructorResignHandler)
	handlerContext.AddHandler("7", "Request raise", requestInstructorRaiseHandler)
	handlerContext.AddHandler("8", "Review leave", reviewLeaveHandler)
	handlerContext.AddHandler("9", "Review resignation", reviewResignationHandler)
	handlerContext.AddHandler("10", "Review raise", reviewRaiseHandler)

	hrMainMenuState := cliManager.GetState(string(MENU_HR))
	backHandler := coreHandler.NewChangeMenuHandlerStrategy(cliManager, hrMainMenuState)
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return &InstructorMenuState{
		manager:        cliManager,
		handlerContext: handlerContext,
	}
}
