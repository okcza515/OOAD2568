package handler

import (
	"ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"errors"
	"fmt"
)

type ProgressMenuStateHandler struct {
	Manager                    *cli.CLIMenuStateManager
	wrapper                    *controller.EvalModuleWrapper
	EvalModuleMenuStateHandler cli.MenuState
	handler                    *handler.HandlerContext
	backhandler                *handler.ChangeMenuHandlerStrategy
}

func NewProgressMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.EvalModuleWrapper, evalModuleMenuStateHandler cli.MenuState) *ProgressMenuStateHandler {
	return &ProgressMenuStateHandler{
		Manager:                    manager,
		wrapper:                    wrapper,
		EvalModuleMenuStateHandler: evalModuleMenuStateHandler,
		handler:                    handler.NewHandlerContext(),
		backhandler:                handler.NewChangeMenuHandlerStrategy(manager, evalModuleMenuStateHandler),
	}
}

func (menu *ProgressMenuStateHandler) Render() {
	util.ClearScreen()
	menu.handler.SetMenuTitle("\nProgress tracking menu:")
	menu.handler.AddHandler("1", "List all progress in each assignment.", handler.FuncStrategy{Action: menu.GetAllProgress})
	menu.handler.AddHandler("2", "Get progress by student ID.", handler.FuncStrategy{Action: menu.GetProgressByStudentCode})
	menu.handler.AddHandler("3", "Get progress by submission status.", handler.FuncStrategy{Action: menu.GetProgressByStatus})
	menu.handler.AddHandler("4", "Get Assignment submit count.", handler.FuncStrategy{Action: menu.GetAssignmentSubmitCount})
	menu.handler.AddBackHandler(menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *ProgressMenuStateHandler) HandlerUserInput(input string) error {
	util.ClearScreen()
	err := menu.handler.HandleInput(input)
	if err != nil {
		return err
	}
	return nil
}

func (menu *ProgressMenuStateHandler) HandleUserInput(input string) error {
	util.ClearScreen()
	return menu.handler.HandleInput(input)
}

func (menu *ProgressMenuStateHandler) getAllProgressTableHeader() {
	fmt.Printf("\n%-5s %-10s %-15s %-15s %-10s", "ID", "Assignment ID", "Student Code", "Submitted", "Total Submit")
	fmt.Printf("\n%-5s %-10s %-15s %-15s %-10s", "---", "------------", "-----------", "---------", "-----------")
}

func (menu *ProgressMenuStateHandler) printProgressTable(progressList []model.Progress) {
	if len(progressList) == 0 {
		fmt.Println("\nNo progress records found.")
		return
	}

	menu.getAllProgressTableHeader()
	for _, progress := range progressList {
		fmt.Printf("\n%-5d %-10d %-15s %-15v %-10d",
			progress.ID,
			progress.AssignmentId,
			progress.StudentCode,
			progress.IsSubmitted,
			progress.TotalSubmit)
	}
	fmt.Println()
}

func (menu *ProgressMenuStateHandler) GetAllProgress() error {
	assignmentId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Assignment ID: ",
		FieldNameText: "Assignment ID",
	}).(uint)

	progressList, err := menu.wrapper.ProgressController.List(map[string]interface{}{
		"assignment_id": assignmentId,
	})
	if err != nil {
		return errors.New("failed to retrieve progress list")
	}

	fmt.Printf("\nAssignment %d Progress List:", assignmentId)
	menu.printProgressTable(progressList)
	return nil
}

func (menu *ProgressMenuStateHandler) GetProgressByStudentCode() error {
	assignmentId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Assignment ID: ",
		FieldNameText: "Assignment ID",
	}).(uint)

	studentCode := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Student Code: ",
		FieldNameText: "Student Code",
	}).(string)

	progressList, err := menu.wrapper.ProgressController.List(map[string]interface{}{
		"assignment_id": assignmentId,
		"student_code":  studentCode,
	})
	if err != nil {
		return errors.New("failed to retrieve progress list")
	}

	fmt.Printf("\nAssignment %d Progress List for Student %s:", assignmentId, studentCode)
	menu.printProgressTable(progressList)
	return nil
}

func (menu *ProgressMenuStateHandler) GetProgressByStatus() error {
	assignmentId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Assignment ID: ",
		FieldNameText: "Assignment ID",
	}).(uint)

	statusStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Submission Status (true/false): ",
		FieldNameText: "Submitted",
	}).(string)

	submitted := statusStr == "true"

	progressList, err := menu.wrapper.ProgressController.List(map[string]interface{}{
		"assignment_id": assignmentId,
		"is_submitted":  submitted,
	})
	if err != nil {
		return errors.New("failed to retrieve progress list")
	}

	fmt.Printf("\nAssignment %d Progress List with Submission Status %v:", assignmentId, submitted)
	menu.printProgressTable(progressList)
	return nil
}

func (menu *ProgressMenuStateHandler) GetAssignmentSubmitCount() error {
	assignmentId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Assignment ID: ",
		FieldNameText: "Assignment ID",
	}).(uint)

	progressList, err := menu.wrapper.ProgressController.List(map[string]interface{}{
		"assignment_id": assignmentId,
	})
	if err != nil {
		return errors.New("failed to retrieve progress list")
	}

	statusCount := make(map[bool]int)
	for _, progress := range progressList {
		statusCount[progress.IsSubmitted]++
	}

	fmt.Printf("\nAssignment %d Submission Count:", assignmentId)
	for submitted, count := range statusCount {
		fmt.Printf("\nSubmitted: %v - Count: %d", submitted, count)
	}
	fmt.Println()
	return nil
}
