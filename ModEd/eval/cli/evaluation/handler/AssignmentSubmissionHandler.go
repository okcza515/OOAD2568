//MEP-1006 Quiz and Assignment

package handler

import (
	"ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"
	"strconv"
	"time"
)

type AssignmentSubmissionMenuStateHandler struct {
	Manager                    *cli.CLIMenuStateManager
	wrapper                    *controller.EvalModuleWrapper
	EvalModuleMenuStateHandler cli.MenuState
	handler                    *handler.HandlerContext
	backhandler                *handler.ChangeMenuHandlerStrategy
}

func NewAssignmentSubmissionMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.EvalModuleWrapper, evalModuleMenuStateHandler cli.MenuState) *AssignmentSubmissionMenuStateHandler {
	return &AssignmentSubmissionMenuStateHandler{
		Manager:                    manager,
		wrapper:                    wrapper,
		EvalModuleMenuStateHandler: evalModuleMenuStateHandler,
		handler:                    handler.NewHandlerContext(),
		backhandler:                handler.NewChangeMenuHandlerStrategy(manager, evalModuleMenuStateHandler),
	}
}

func (menu *AssignmentSubmissionMenuStateHandler) Render() {
	util.ClearScreen()
	menu.handler.SetMenuTitle("\nAssignment Submission Menu")
	menu.handler.AddHandler("1", "View Available Assignments", handler.FuncStrategy{Action: menu.ViewPublishedAssignments})
	menu.handler.AddHandler("2", "Create/Update Draft Submission", handler.FuncStrategy{Action: menu.CreateDraftSubmission})
	menu.handler.AddHandler("3", "Submit Assignment", handler.FuncStrategy{Action: menu.SubmitAssignment})
	menu.handler.AddHandler("4", "View My Submissions", handler.FuncStrategy{Action: menu.ViewMySubmissions})
	menu.handler.AddHandler("5", "View All Submissions for Assignment (Admin)", handler.FuncStrategy{Action: menu.ViewSubmissionsByAssignment})
	menu.handler.AddHandler("6", "Delete Submission", handler.FuncStrategy{Action: menu.DeleteSubmission})
	menu.handler.AddBackHandler(menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *AssignmentSubmissionMenuStateHandler) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *AssignmentSubmissionMenuStateHandler) ViewPublishedAssignments() error {
	// Get all assignments
	assignments, err := menu.wrapper.AssignmentController.GetAssignments()
	if err != nil {
		return fmt.Errorf("error retrieving assignments: %v", err)
	}

	// Filter to only published assignments
	var publishedAssignments []*model.Assignment
	for _, assignment := range assignments {
		if assignment.Status == model.StatusPublished {
			publishedAssignments = append(publishedAssignments, assignment)
		}
	}

	if len(publishedAssignments) == 0 {
		fmt.Println("\nNo published assignments available for submission.")
		util.PressEnterToContinue()
		util.ClearScreen()
		return nil
	}

	// Display the published assignments
	fmt.Println("\nPublished Assignments Available for Submission:")
	fmt.Printf("\n%-5s %-20s %-30s %-15s %-15s",
		"ID", "Title", "Description", "Publish Date", "Due Date")
	fmt.Printf("\n%-5s %-20s %-30s %-15s %-15s",
		"--", "-----", "-----------", "------------", "--------")

	for _, assignment := range publishedAssignments {
		// Truncate description if too long
		description := assignment.Description
		if len(description) > 27 {
			description = description[:27] + "..."
		}

		fmt.Printf("\n%-5d %-20s %-30s %-15s %-15s",
			assignment.AssignmentId,
			assignment.Title,
			description,
			assignment.PublishDate.Format("2006-01-02"),
			assignment.DueDate.Format("2006-01-02"))
	}

	fmt.Println("\n")
	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *AssignmentSubmissionMenuStateHandler) CreateDraftSubmission() error {
	assignmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Assignment ID:",
		FieldNameText: "Assignment ID",
	}).(string)

	assignmentId, err := strconv.ParseUint(assignmentIdStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid assignment ID: %v", err)
	}

	// Get assignment to check if it exists
	assignment, err := menu.wrapper.AssignmentController.GetAssignment(uint(assignmentId))
	if err != nil {
		return fmt.Errorf("assignment not found: %v", err)
	}

	// Verify assignment is published
	if assignment.Status != model.StatusPublished {
		return fmt.Errorf("cannot submit to an assignment that is not published")
	}

	studentCode := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter your Student Code:",
		FieldNameText: "Student Code",
	}).(string)

	answer := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter your Answer:",
		FieldNameText: "Answer",
	}).(string)

	// Create submission object
	submission := &model.AssignmentSubmission{
		AssignmentId: uint(assignmentId),
		StudentCode:  studentCode,
		Answer:       answer,
		Submitted:    false, // Draft mode
		UpdatedAt:    time.Now(),
	}

	// Try to get existing submission
	existingSubmission, err := menu.wrapper.AssignmentSubmissionController.GetSubmissionByAssignmentAndStudent(
		uint(assignmentId), studentCode)

	if err == nil {
		// Update existing submission
		submission.ID = existingSubmission.ID
		err = menu.wrapper.AssignmentSubmissionController.UpdateSubmission(submission)
		if err != nil {
			return fmt.Errorf("error updating submission: %v", err)
		}
		fmt.Println("\nDraft submission updated successfully!")
	} else {
		// Create new submission
		err = menu.wrapper.AssignmentSubmissionController.CreateSubmission(submission)
		if err != nil {
			return fmt.Errorf("error creating submission: %v", err)
		}
		fmt.Println("\nDraft submission created successfully!")
	}

	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *AssignmentSubmissionMenuStateHandler) SubmitAssignment() error {
	assignmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Assignment ID:",
		FieldNameText: "Assignment ID",
	}).(string)

	assignmentId, err := strconv.ParseUint(assignmentIdStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid assignment ID: %v", err)
	}

	// Get assignment to check if it exists
	assignment, err := menu.wrapper.AssignmentController.GetAssignment(uint(assignmentId))
	if err != nil {
		return fmt.Errorf("assignment not found: %v", err)
	}

	// Verify assignment is published
	if assignment.Status != model.StatusPublished {
		return fmt.Errorf("cannot submit to an assignment that is not published")
	}

	studentCode := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter your Student Code:",
		FieldNameText: "Student Code",
	}).(string)

	// Check if there's an existing draft or get new answer
	existingSubmission, err := menu.wrapper.AssignmentSubmissionController.GetSubmissionByAssignmentAndStudent(
		uint(assignmentId), studentCode)

	var answer string
	if err == nil && !existingSubmission.Submitted {
		// We have a draft
		useExistingAnswer := core.ExecuteUserInputStep(core.StringInputStep{
			PromptText:    fmt.Sprintf("Use existing answer: \"%s\"? (yes/no):", existingSubmission.Answer),
			FieldNameText: "Use Existing",
		}).(string)

		if useExistingAnswer == "yes" {
			answer = existingSubmission.Answer
		} else {
			answer = core.ExecuteUserInputStep(core.StringInputStep{
				PromptText:    "Enter your Answer:",
				FieldNameText: "Answer",
			}).(string)
		}
	} else {
		// No draft, get new answer
		answer = core.ExecuteUserInputStep(core.StringInputStep{
			PromptText:    "Enter your Answer:",
			FieldNameText: "Answer",
		}).(string)
	}

	// Create submission object
	submission := &model.AssignmentSubmission{
		AssignmentId: uint(assignmentId),
		StudentCode:  studentCode,
		Answer:       answer,
		Submitted:    true, // Mark as submitted
		UpdatedAt:    time.Now(),
	}

	// Submit the assignment
	err = menu.wrapper.AssignmentSubmissionController.SubmitAssignment(submission)
	if err != nil {
		return fmt.Errorf("error submitting assignment: %v", err)
	}

	fmt.Println("\nAssignment submitted successfully!")
	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *AssignmentSubmissionMenuStateHandler) ViewMySubmissions() error {
	studentCode := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter your Student Code:",
		FieldNameText: "Student Code",
	}).(string)

	submissions, err := menu.wrapper.AssignmentSubmissionController.GetSubmissionsByStudent(studentCode)
	if err != nil {
		return fmt.Errorf("error retrieving submissions: %v", err)
	}

	if len(submissions) == 0 {
		fmt.Println("\nNo submissions found.")
		util.PressEnterToContinue()
		util.ClearScreen()
		return nil
	}

	menu.displaySubmissionTableHeader()
	for _, submission := range submissions {
		menu.displaySubmission(submission)
	}
	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *AssignmentSubmissionMenuStateHandler) ViewSubmissionsByAssignment() error {
	assignmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Assignment ID:",
		FieldNameText: "Assignment ID",
	}).(string)

	assignmentId, err := strconv.ParseUint(assignmentIdStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid assignment ID: %v", err)
	}

	submissions, err := menu.wrapper.AssignmentSubmissionController.GetSubmissionsByAssignment(uint(assignmentId))
	if err != nil {
		return fmt.Errorf("error retrieving submissions: %v", err)
	}

	if len(submissions) == 0 {
		fmt.Println("\nNo submissions found for this assignment.")
		util.PressEnterToContinue()
		util.ClearScreen()
		return nil
	}

	menu.displaySubmissionTableHeader()
	for _, submission := range submissions {
		menu.displaySubmission(submission)
	}
	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *AssignmentSubmissionMenuStateHandler) DeleteSubmission() error {
	submissionIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Submission ID to delete:",
		FieldNameText: "Submission ID",
	}).(string)

	submissionId, err := strconv.ParseUint(submissionIdStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid submission ID: %v", err)
	}

	confirmation := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    fmt.Sprintf("Are you sure you want to delete submission %d? (yes/no):", submissionId),
		FieldNameText: "Confirmation",
	}).(string)

	if confirmation != "yes" {
		fmt.Println("\nDeletion cancelled")
		util.PressEnterToContinue()
		util.ClearScreen()
		return nil
	}

	err = menu.wrapper.AssignmentSubmissionController.DeleteSubmission(uint(submissionId))
	if err != nil {
		return fmt.Errorf("failed to delete submission: %v", err)
	}

	fmt.Println("\nSubmission deleted successfully")
	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *AssignmentSubmissionMenuStateHandler) displaySubmissionTableHeader() {
	fmt.Printf("\n%-5s %-15s %-15s %-40s %-12s %-20s",
		"ID", "Assignment ID", "Student Code", "Answer", "Status", "Updated At")
	fmt.Printf("\n%-5s %-15s %-15s %-40s %-12s %-20s",
		"---", "-------------", "------------", "------", "------", "----------")
}

func (menu *AssignmentSubmissionMenuStateHandler) displaySubmission(submission *model.AssignmentSubmission) {
	status := "Draft"
	if submission.Submitted {
		status = "Submitted"
	}

	// Truncate answer if it's too long for display
	answer := submission.Answer
	if len(answer) > 37 {
		answer = answer[:37] + "..."
	}

	fmt.Printf("\n%-5d %-15d %-15s %-40s %-12s %-20s",
		submission.ID,
		submission.AssignmentId,
		submission.StudentCode,
		answer,
		status,
		submission.UpdatedAt.Format("2006-01-02 15:04:05"))
}
