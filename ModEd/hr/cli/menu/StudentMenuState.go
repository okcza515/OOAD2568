package menu

import (
	"ModEd/core/cli"
	coreHandler "ModEd/core/handler"
	hrHandler "ModEd/hr/cli/menu/handler"
	"ModEd/hr/controller"
	"fmt"
)

type StudentMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *coreHandler.HandlerContext
}

func (a *StudentMenuState) HandleUserInput(input string) error {
	err := a.handlerContext.HandleInput(input)
	if err != nil {
		return err
	}

	return nil
}

func (a *StudentMenuState) Render() {
	fmt.Println("=== Student Menu ===")
	a.handlerContext.ShowMenu()
	fmt.Println("exit:\tExit the program.")
	fmt.Println()
}

func NewStudentMenuState(
	cliManager *cli.CLIMenuStateManager,
	hrManager *controller.HRControllerManager,
) *StudentMenuState {
	handlerContext := coreHandler.NewHandlerContext()

	addStudentHandler := hrHandler.NewAddStudentStrategy(hrManager.StudentCtrl)
	listStudentHandler := hrHandler.NewListStudentStrategy(hrManager.StudentCtrl)
	deleteStudentHandler := hrHandler.NewDeleteStudentStrategy(hrManager.StudentCtrl)
	updateStudentInfoHandler := hrHandler.NewUpdateStudentInfoStrategy(hrManager.StudentCtrl)
	requestStudentLeaveHandler := hrHandler.NewRequestLeaveHandlerStrategy(hrManager.LeaveStudentCtrl.SubmitStudentLeaveRequest)
	requestStudentResignHandler := hrHandler.NewRequestResignationHandlerStrategy(hrManager.ResignStudentCtrl.SubmitResignationStudent)
	reviewLeaveHandler := hrHandler.NewReviewHandlerStrategy(hrManager.LeaveStudentCtrl.ReviewStudentLeaveRequest)
	reviewResignationHandler := hrHandler.NewReviewHandlerStrategy(hrManager.ResignStudentCtrl.ReviewStudentResignRequest)

	handlerContext.AddHandler("1", "Add new student", addStudentHandler)
	handlerContext.AddHandler("2", "List student", listStudentHandler)
	handlerContext.AddHandler("3", "Update student Info", updateStudentInfoHandler)
	handlerContext.AddHandler("4", "Delete student", deleteStudentHandler)
	handlerContext.AddHandler("5", "Request leave", requestStudentLeaveHandler)
	handlerContext.AddHandler("6", "Request resignation", requestStudentResignHandler)
	handlerContext.AddHandler("7", "Review leave", reviewLeaveHandler)
	handlerContext.AddHandler("8", "Review resignation", reviewResignationHandler)

	hrMainMenuState := cliManager.GetState(string(MENU_HR))
	backHandler := coreHandler.NewChangeMenuHandlerStrategy(cliManager, hrMainMenuState)
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return &StudentMenuState{
		manager:        cliManager,
		handlerContext: handlerContext,
	}
}
