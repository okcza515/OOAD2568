package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	hrHandler "ModEd/hr/cli/menu/handler"
	"ModEd/hr/controller"
	"fmt"
)

type DatabaseMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func (a *DatabaseMenuState) HandleUserInput(input string) error {
	err := a.handlerContext.HandleInput(input)
	if err != nil {
		return err
	}

	return nil
}

func (a *DatabaseMenuState) Render() {
	fmt.Println("=== Database Menu ===")
	a.handlerContext.ShowMenu()
	fmt.Println("exit:\tExit the program.")
	fmt.Println()
}

func NewDatabaseMenuState(
	cliManager *cli.CLIMenuStateManager,
	hrManager *controller.HRControllerManager,
) *DatabaseMenuState {

	handlerContext := handler.NewHandlerContext()

	pullStudentHandler := hrHandler.NewPullHandlerStrategy(hrManager.StudentCtrl.MigrateStudentRecords)
	pullInstructorHandler := hrHandler.NewPullHandlerStrategy(hrManager.InstructorCtrl.MigrateInstructorRecords)
	importStudentHandler := hrHandler.NewImportHandlerStrategy(hrManager.StudentCtrl.ImportStudents)
	importInstructorHandler := hrHandler.NewImportHandlerStrategy(hrManager.InstructorCtrl.ImportInstructors)
	exportStudentHandler := hrHandler.NewExportHandlerStrategy(hrManager.StudentCtrl.ExportStudents)
	exportInstructorHandler := hrHandler.NewExportHandlerStrategy(hrManager.InstructorCtrl.ExportInstructors)
	exportStudentLeaveHandler := hrHandler.NewExportHandlerStrategy(hrManager.LeaveStudentCtrl.ExportStudentLeaveRequests)
	exportInstructorLeaveHandler := hrHandler.NewExportHandlerStrategy(hrManager.LeaveInstructorCtrl.ExportInstructorLeaveRequests)
	exportStudentResignHandler := hrHandler.NewExportHandlerStrategy(hrManager.ResignStudentCtrl.ExportStudentResignRequests)
	exportInstructorResignHandler := hrHandler.NewExportHandlerStrategy(hrManager.ResignInstructorCtrl.ExportInstructorResignRequests)
	exportInstructorRaiseHandler := hrHandler.NewExportHandlerStrategy(hrManager.RaiseCtrl.ExportInstructorRaiseRequests)

	handlerContext.AddHandler("1", "Pull student data", pullStudentHandler)
	handlerContext.AddHandler("2", "Pull instructor data", pullInstructorHandler)
	handlerContext.AddHandler("3", "Import student data", importStudentHandler)
	handlerContext.AddHandler("4", "Import instructor data", importInstructorHandler)
	handlerContext.AddHandler("5", "Export student data", exportStudentHandler)
	handlerContext.AddHandler("6", "Export instructor data", exportInstructorHandler)
	handlerContext.AddHandler("7", "Export student leave requests", exportStudentLeaveHandler)
	handlerContext.AddHandler("8", "Export instructor leave requests", exportInstructorLeaveHandler)
	handlerContext.AddHandler("9", "Export student resignation requests", exportStudentResignHandler)
	handlerContext.AddHandler("10", "Export instructor resignation requests", exportInstructorResignHandler)
	handlerContext.AddHandler("11", "Export instructor raise requests", exportInstructorRaiseHandler) // Corrected typo from exportInsturctorRaiseHandler

	hrMainMenuState := cliManager.GetState(string(MENU_HR))
	backHandler := handler.NewChangeMenuHandlerStrategy(cliManager, hrMainMenuState)
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return &DatabaseMenuState{
		manager:        cliManager,
		handlerContext: handlerContext,
	}
}
