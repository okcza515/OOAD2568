package handler

import (
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
	menu.handler.SetMenuTitle("\nProgress tracking menu:")
	menu.handler.AddHandler("1", "List all progress in each assessment.", handler.FuncStrategy{Action: menu.GetAllProgress})
	menu.handler.AddHandler("2", "Get progress by student ID.", handler.FuncStrategy{Action: menu.GetProgressByStudentCode})
	menu.handler.AddHandler("3", "Get progress by assessment status.", handler.FuncStrategy{Action: menu.GetProgressByStatus})
	menu.handler.AddHandler("4", "Get Assessment submit count.", handler.FuncStrategy{Action: menu.GetAssessmentSubmitCount})

	menu.handler.ShowMenu()
}

func (menu *ProgressMenuStateHandler) HandlerUserInput(input string) error {
	err := menu.handler.HandleInput(input)
	if err != nil {
		return err
	}
	return nil
}

func (menu *ProgressMenuStateHandler) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *ProgressMenuStateHandler) getAllProgressTableHeader() {
	fmt.Printf("\n%-5s %-10s %-15s %-15s %-10s", "ID", "Assessment ID", "Student Code", "Type", "Status")
	fmt.Printf("\n%-5s %-10s %-15s %-15s %-10s", "---", "------------", "-----------", "------", "------")
}

func (menu *ProgressMenuStateHandler) printProgressTable(progressList []model.Progress) {
	if len(progressList) == 0 {
		fmt.Println("\nNo progress records found.")
		return
	}

	menu.getAllProgressTableHeader()
	for _, progress := range progressList {
		fmt.Printf("\n%-5d %-10d %-15s %-10s %-15s",
			progress.ID,
			progress.AssessmentId,
			progress.StudentCode,
			progress.Type,
			progress.Status)
	}
	fmt.Println()
}

func (menu *ProgressMenuStateHandler) GetAllProgress() error {
	assessmentId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Assessment ID:",
		FieldNameText: "Assessment ID",
	}).(uint)

	progressList, err := menu.wrapper.ProgressController.List(map[string]interface{}{
		"assessment_id": assessmentId,
	})
	if err != nil {
		return errors.New("Failed to retrieve progress list.")
	}

	fmt.Printf("\nAssessment %d Progress List:", assessmentId)
	menu.printProgressTable(progressList)
	return nil
}

func (menu *ProgressMenuStateHandler) GetProgressByStudentCode() error {
	assessmentId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Assessment ID:",
		FieldNameText: "Assessment ID",
	}).(uint)

	studentCode := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Student Code:",
		FieldNameText: "Student Code",
	}).(string)

	progressList, err := menu.wrapper.ProgressController.List(map[string]interface{}{
		"assessment_id": assessmentId,
		"student_code":  studentCode,
	})
	if err != nil {
		return errors.New("Failed to retrieve progress list.")
	}

	fmt.Printf("\nAssessment %d Progress List for Student %s:", assessmentId, studentCode)
	menu.printProgressTable(progressList)
	return nil
}

func (menu *ProgressMenuStateHandler) GetProgressByStatus() error {
	assessmentId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Assessment ID:",
		FieldNameText: "Assessment ID",
	}).(uint)

	status := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Status:",
		FieldNameText: "Status",
	}).(string)

	progressList, err := menu.wrapper.ProgressController.List(map[string]interface{}{
		"assessment_id": assessmentId,
		"status":        status,
	})
	if err != nil {
		return errors.New("Failed to retrieve progress list.")
	}

	fmt.Printf("\nAssessment %d Progress List with Status %s:", assessmentId, status)
	menu.printProgressTable(progressList)
	return nil
}

func (menu *ProgressMenuStateHandler) GetAssessmentSubmitCount() error {
	assessmentId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Assessment ID:",
		FieldNameText: "Assessment ID",
	}).(uint)

	progressList, err := menu.wrapper.ProgressController.List(map[string]interface{}{
		"assessment_id": assessmentId,
	})
	if err != nil {
		return errors.New("Failed to retrieve progress list.")
	}

	statusCount := make(map[model.AssessmentStatus]int)
	for _, progress := range progressList {
		statusCount[progress.Status]++
	}

	fmt.Printf("\nAssessment %d Submission Count:", assessmentId)
	for status, count := range statusCount {
		fmt.Printf("\n%s: %d", status, count)
	}
	fmt.Printf("\nTotal Submissions: %d", len(progressList))
	fmt.Println()
	return nil
}
