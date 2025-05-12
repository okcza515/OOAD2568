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
	"strconv"
	"time"
)

type AssignmentMenuStateHandler struct {
	Manager                    *cli.CLIMenuStateManager
	wrapper                    *controller.EvalModuleWrapper
	EvalModuleMenuStateHandler cli.MenuState
	handler                    *handler.HandlerContext
	backhandler                *handler.ChangeMenuHandlerStrategy
}

func NewAssignmentMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.EvalModuleWrapper, evalModuleMenuStateHandler cli.MenuState) *AssignmentMenuStateHandler {
	return &AssignmentMenuStateHandler{
		Manager:                    manager,
		wrapper:                    wrapper,
		EvalModuleMenuStateHandler: evalModuleMenuStateHandler,
		handler:                    handler.NewHandlerContext(),
		backhandler:                handler.NewChangeMenuHandlerStrategy(manager, evalModuleMenuStateHandler),
	}
}

func (menu *AssignmentMenuStateHandler) Render() {
	util.ClearScreen()
	menu.handler.SetMenuTitle("\nAssignment management menu:")
	menu.handler.AddHandler("1", "Create a new assignment.", handler.FuncStrategy{Action: menu.CreateAssignment})
	menu.handler.AddHandler("2", "View all assignments.", handler.FuncStrategy{Action: menu.ViewAllAssignments})
	menu.handler.AddHandler("3", "View assignment by ID.", handler.FuncStrategy{Action: menu.ViewAssignmentByID})
	menu.handler.AddHandler("4", "Update assignment details.", handler.FuncStrategy{Action: menu.UpdateAssignment})
	menu.handler.AddHandler("5", "Change assignment status.", handler.FuncStrategy{Action: menu.ChangeAssignmentStatus})
	menu.handler.AddHandler("6", "Delete an assignment.", handler.FuncStrategy{Action: menu.DeleteAssignment})
	menu.handler.AddHandler("7", "Assignment Submission Menu", handler.FuncStrategy{Action: menu.GoToSubmissionMenu})
	menu.handler.AddHandler("back", "Back to previous menu.", menu.backhandler)

	menu.handler.ShowMenu()
}

func (menu *AssignmentMenuStateHandler) HandleUserInput(input string) error {
	util.ClearScreen()
	return menu.handler.HandleInput(input)
}

func (menu *AssignmentMenuStateHandler) printAssignmentHeader() {
	fmt.Printf("\n%-5s %-20s %-15s %-15s %-10s %-10s", "ID", "Title", "Publish Date", "Due Date", "Status", "Class ID")
	fmt.Printf("\n%-5s %-20s %-15s %-15s %-10s %-10s", "---", "-----", "------------", "--------", "------", "--------")
}

func (menu *AssignmentMenuStateHandler) printAssignmentRow(assignment *model.Assignment) {
	publishDate := assignment.PublishDate.Format("2006-01-02")
	dueDate := assignment.DueDate.Format("2006-01-02")

	fmt.Printf("\n%-5d %-20s %-15s %-15s %-10s %-10v",
		assignment.AssignmentId,
		truncateStr(assignment.Title, 20),
		publishDate,
		dueDate,
		assignment.Status,
		assignment.ClassId)
}

// Helper function to truncate strings for display
func truncateStr(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func (menu *AssignmentMenuStateHandler) printAssignmentList(assignments []*model.Assignment) {
	if len(assignments) == 0 {
		fmt.Println("\nNo assignments found.")
		return
	}

	menu.printAssignmentHeader()
	for _, assignment := range assignments {
		menu.printAssignmentRow(assignment)
	}
	fmt.Println()
}

func (menu *AssignmentMenuStateHandler) CreateAssignment() error {
	title := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assignment title:",
		FieldNameText: "Title",
	}).(string)

	description := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assignment description:",
		FieldNameText: "Description",
	}).(string)

	publishDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter publish date (YYYY-MM-DD):",
		FieldNameText: "Publish Date",
	}).(string)
	publishDate, err := time.Parse("2006-01-02", publishDateStr)
	if err != nil {
		return errors.New("invalid date format for publish date")
	}

	dueDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter due date (YYYY-MM-DD):",
		FieldNameText: "Due Date",
	}).(string)
	dueDate, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		return errors.New("invalid date format for due date")
	}

	classIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter class ID:",
		FieldNameText: "Class ID",
	}).(string)
	classIdUint, err := strconv.ParseUint(classIdStr, 10, 64)
	if err != nil {
		return errors.New("invalid class ID")
	}

	instructorCode := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter instructor code:",
		FieldNameText: "Instructor Code",
	}).(string)

	assignment := &model.Assignment{
		Title:          title,
		Description:    description,
		PublishDate:    publishDate,
		DueDate:        dueDate,
		Status:         model.StatusDraft,
		ClassId:        uint(classIdUint),
		InstructorCode: instructorCode,
		State:          &controller.DraftState{},
	}

	assignmentId, err := menu.wrapper.AssignmentController.CreateAssignment(assignment)
	if err != nil {
		return errors.New("failed to create assignment: " + err.Error())
	}

	fmt.Printf("\nAssignment created successfully with ID: %d\n", assignmentId)
	fmt.Println("Assignment has been saved to CSV file automatically.")
	return nil
}

func (menu *AssignmentMenuStateHandler) ViewAllAssignments() error {
	assignments, err := menu.wrapper.AssignmentController.GetAssignments()
	if err != nil {
		return errors.New("failed to retrieve assignments: " + err.Error())
	}

	fmt.Println("\nAll Assignments:")
	menu.printAssignmentList(assignments)
	return nil
}

func (menu *AssignmentMenuStateHandler) ViewAssignmentByID() error {
	assignmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assignment ID:",
		FieldNameText: "Assignment ID",
	}).(string)

	assignmentId, err := strconv.ParseUint(assignmentIdStr, 10, 64)
	if err != nil {
		return errors.New("invalid assignment ID")
	}

	assignment, err := menu.wrapper.AssignmentController.GetAssignment(uint(assignmentId))
	if err != nil {
		return errors.New("failed to retrieve assignment: " + err.Error())
	}

	fmt.Printf("\nAssignment Details (ID: %d):\n", assignmentId)
	fmt.Printf("Title: %s\n", assignment.Title)
	fmt.Printf("Description: %s\n", assignment.Description)
	fmt.Printf("Publish Date: %s\n", assignment.PublishDate.Format("2006-01-02"))
	fmt.Printf("Due Date: %s\n", assignment.DueDate.Format("2006-01-02"))
	fmt.Printf("Status: %s\n", assignment.Status)
	fmt.Printf("Class ID: %v\n", assignment.ClassId)
	fmt.Printf("Instructor Code: %s\n", assignment.InstructorCode)
	return nil
}

func (menu *AssignmentMenuStateHandler) UpdateAssignment() error {
	assignmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assignment ID to update:",
		FieldNameText: "Assignment ID",
	}).(string)

	assignmentId, err := strconv.ParseUint(assignmentIdStr, 10, 64)
	if err != nil {
		return errors.New("invalid assignment ID")
	}

	assignment, err := menu.wrapper.AssignmentController.GetAssignment(uint(assignmentId))
	if err != nil {
		return errors.New("failed to retrieve assignment: " + err.Error())
	}

	fmt.Printf("Current title: %s\n", assignment.Title)
	title := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new title (or press Enter to keep current):",
		FieldNameText: "Title",
	}).(string)
	if title == "" {
		title = assignment.Title
	}

	fmt.Printf("Current description: %s\n", assignment.Description)
	description := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new description (or press Enter to keep current):",
		FieldNameText: "Description",
	}).(string)
	if description == "" {
		description = assignment.Description
	}

	fmt.Printf("Current publish date: %s\n", assignment.PublishDate.Format("2006-01-02"))
	publishDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new publish date (YYYY-MM-DD) (or press Enter to keep current):",
		FieldNameText: "Publish Date",
	}).(string)

	publishDate := assignment.PublishDate
	if publishDateStr != "" {
		var err error
		publishDate, err = time.Parse("2006-01-02", publishDateStr)
		if err != nil {
			return errors.New("invalid date format for publish date")
		}
	}

	fmt.Printf("Current due date: %s\n", assignment.DueDate.Format("2006-01-02"))
	dueDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new due date (YYYY-MM-DD) (or press Enter to keep current):",
		FieldNameText: "Due Date",
	}).(string)

	dueDate := assignment.DueDate
	if dueDateStr != "" {
		var err error
		dueDate, err = time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			return errors.New("invalid date format for due date")
		}
	}

	fmt.Printf("Current class ID: %v\n", assignment.ClassId)
	classIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new class ID (or press Enter to keep current):",
		FieldNameText: "Class ID",
	}).(string)

	classId := assignment.ClassId
	if classIdStr != "" {
		classIdUint, err := strconv.ParseUint(classIdStr, 10, 64)
		if err != nil {
			return errors.New("invalid class ID")
		}
		classId = uint(classIdUint)
	}

	fmt.Printf("Current instructor code: %s\n", assignment.InstructorCode)
	instructorCodeStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new instructor code (or press Enter to keep current):",
		FieldNameText: "Instructor Code",
	}).(string)

	instructorCode := assignment.InstructorCode
	if instructorCodeStr != "" {
		instructorCode = instructorCodeStr
	}

	updatedAssignment := &model.Assignment{
		AssignmentId:   uint(assignmentId),
		Title:          title,
		Description:    description,
		PublishDate:    publishDate,
		DueDate:        dueDate,
		Status:         assignment.Status,
		ClassId:        classId,
		InstructorCode: instructorCode,
		State:          assignment.State,
	}

	_, err = menu.wrapper.AssignmentController.UpdateAssignment(updatedAssignment)
	if err != nil {
		return errors.New("failed to update assignment: " + err.Error())
	}

	fmt.Println("\nAssignment updated successfully")
	fmt.Println("Changes have been saved to CSV file automatically.")
	return nil
}

func (menu *AssignmentMenuStateHandler) ChangeAssignmentStatus() error {
	assignmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assignment ID:",
		FieldNameText: "Assignment ID",
	}).(string)

	assignmentId, err := strconv.ParseUint(assignmentIdStr, 10, 64)
	if err != nil {
		return errors.New("invalid assignment ID")
	}

	assignment, err := menu.wrapper.AssignmentController.GetAssignment(uint(assignmentId))
	if err != nil {
		return errors.New("Failed to retrieve assignment: " + err.Error())
	}

	fmt.Printf("\nCurrent status: %s\n", assignment.Status)
	fmt.Println("Available status options:")
	fmt.Println("1. Draft")
	fmt.Println("2. Published")
	fmt.Println("3. Closed")

	choice := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Select new status (1-3):",
		FieldNameText: "Status Choice",
	}).(string)

	var newStatus model.AssignmentStatus
	switch choice {
	case "1":
		newStatus = model.StatusDraft
	case "2":
		newStatus = model.StatusPublished
	case "3":
		newStatus = model.StatusClosed
	default:
		return errors.New("invalid status choice")
	}

	err = menu.wrapper.AssignmentController.UpdateAssignmentStatus(uint(assignmentId), newStatus)
	if err != nil {
		return errors.New("failed to update assignment status: " + err.Error())
	}

	fmt.Printf("\nAssignment status updated to: %s\n", newStatus)
	fmt.Println("Changes have been saved to CSV file.")
	return nil
}

func (menu *AssignmentMenuStateHandler) DeleteAssignment() error {
	assignmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assignment ID to delete:",
		FieldNameText: "Assignment ID",
	}).(string)

	assignmentId, err := strconv.ParseUint(assignmentIdStr, 10, 64)
	if err != nil {
		return errors.New("invalid assignment ID")
	}

	confirmation := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    fmt.Sprintf("Are you sure you want to delete assignment %d? (yes/no):", assignmentId),
		FieldNameText: "Confirmation",
	}).(string)

	if confirmation != "yes" {
		fmt.Println("\nDeletion cancelled")
		return nil
	}

	_, err = menu.wrapper.AssignmentController.DeleteAssignment(uint(assignmentId))
	if err != nil {
		return errors.New("Failed to delete assignment: " + err.Error())
	}

	fmt.Println("\nAssignment deleted successfully")
	return nil
}

func (menu *AssignmentMenuStateHandler) GoToSubmissionMenu() error {
	util.ClearScreen()
	// Create the assignment submission menu handler
	submissionHandler := NewAssignmentSubmissionMenuStateHandler(menu.Manager, menu.wrapper, menu)
	// Set it as the current state
	menu.Manager.SetState(submissionHandler)
	return nil
}
