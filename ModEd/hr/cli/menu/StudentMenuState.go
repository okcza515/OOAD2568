package menu

import (
	"ModEd/core/cli"
	coreHandler "ModEd/core/handler"
	hrHandler "ModEd/hr/cli/menu/handler"
	"ModEd/hr/controller"
	"fmt"
)

type StudentMenuState struct {
	manager           *cli.CLIMenuStateManager
	handlerContext    *coreHandler.HandlerContext
	studentController *controller.StudentHRController
}

// HandleUserInput implements cli.MenuState.
func (a *StudentMenuState) HandleUserInput(input string) error {
	err := a.handlerContext.HandleInput(input)
	if err != nil {
		return err
	}

	return nil
}

// Render implements cli.MenuState.
func (a *StudentMenuState) Render() {
	fmt.Println("=== Student Menu ===")
	a.handlerContext.ShowMenu()
	fmt.Println("exit:\tExit the program.")
}

func NewStudentMenuState(
	manager *cli.CLIMenuStateManager,
	studentCtrl *controller.StudentHRController,
	leaveStudentCtrl *controller.LeaveStudentHRController,
	resignStudentCtrl *controller.ResignationStudentHRController,
) *StudentMenuState {
	handlerContext := coreHandler.NewHandlerContext()

	// Pass the controller to your strategy/handler
	addStudentHandler := hrHandler.NewAddStudentStrategy(studentCtrl)
	listStudentHandler := hrHandler.NewListStudentStrategy(studentCtrl)
	deleteStudentHandler := hrHandler.NewDeleteStudentStrategy(studentCtrl)
	updateStudentInfoHandler := hrHandler.NewUpdateStudentInfoStrategy(studentCtrl)
	requestStudentLeaveHandler := hrHandler.NewRequestLeaveHandlerStrategy(leaveStudentCtrl.SubmitStudentLeaveRequest)
	requestStudentResignHandler := hrHandler.NewRequestResignationHandlerStrategy(resignStudentCtrl.SubmitResignationStudent)
	reviewLeaveHandler := hrHandler.NewReviewHandlerStrategy(leaveStudentCtrl.ReviewStudentLeaveRequest)
	reviewResignationHandler := hrHandler.NewReviewHandlerStrategy(resignStudentCtrl.ReviewStudentResignRequest)
	handlerContext.AddHandler("1", "Add new student", addStudentHandler)
	handlerContext.AddHandler("2", "List student", listStudentHandler)
	handlerContext.AddHandler("3", "Update student Info", updateStudentInfoHandler)
	handlerContext.AddHandler("4", "Delete student", deleteStudentHandler)
	handlerContext.AddHandler("5", "Request leave", requestStudentLeaveHandler)
	handlerContext.AddHandler("6", "Request resignation", requestStudentResignHandler)
	handlerContext.AddHandler("7", "Review leave", reviewLeaveHandler)
	handlerContext.AddHandler("8", "Review resignation", reviewResignationHandler)

	backHandler := coreHandler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_HR)))
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return &StudentMenuState{
		manager:           manager,
		handlerContext:    handlerContext,
		studentController: studentCtrl,
	}
}
