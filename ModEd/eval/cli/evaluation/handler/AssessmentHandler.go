package handler

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	curriculumModel "ModEd/curriculum/model"
	"ModEd/eval/controller"
	"ModEd/eval/model"

	"errors"
	"fmt"
	"strconv"
	"time"
)

type AssessmentMenuStateHandler struct {
	Manager                    *cli.CLIMenuStateManager
	wrapper                    *controller.EvalModuleWrapper
	EvalModuleMenuStateHandler cli.MenuState
	handler                    *handler.HandlerContext
	backhandler                *handler.ChangeMenuHandlerStrategy
}

func NewAssessmentMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.EvalModuleWrapper, evalModuleMenuStateHandler cli.MenuState) *AssessmentMenuStateHandler {
	return &AssessmentMenuStateHandler{
		Manager:                    manager,
		wrapper:                    wrapper,
		EvalModuleMenuStateHandler: evalModuleMenuStateHandler,
		handler:                    handler.NewHandlerContext(),
		backhandler:                handler.NewChangeMenuHandlerStrategy(manager, evalModuleMenuStateHandler),
	}
}

func (menu *AssessmentMenuStateHandler) Render() {
	menu.handler.SetMenuTitle("\nAssessment management menu:")
	menu.handler.AddHandler("1", "Create a new assessment.", handler.FuncStrategy{Action: menu.CreateAssessment})
	menu.handler.AddHandler("2", "View all assessments.", handler.FuncStrategy{Action: menu.ViewAllAssessments})
	menu.handler.AddHandler("3", "View assessment by ID.", handler.FuncStrategy{Action: menu.ViewAssessmentByID})
	menu.handler.AddHandler("4", "Update assessment details.", handler.FuncStrategy{Action: menu.UpdateAssessment})
	menu.handler.AddHandler("5", "Change assessment status.", handler.FuncStrategy{Action: menu.ChangeAssessmentStatus})
	menu.handler.AddHandler("6", "Delete an assessment.", handler.FuncStrategy{Action: menu.DeleteAssessment})
	menu.handler.AddHandler("b", "Back to previous menu.", menu.backhandler)

	menu.handler.ShowMenu()
}

func (menu *AssessmentMenuStateHandler) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *AssessmentMenuStateHandler) printAssessmentHeader() {
	fmt.Printf("\n%-5s %-20s %-15s %-15s %-10s %-10s", "ID", "Title", "Publish Date", "Due Date", "Status", "Class ID")
	fmt.Printf("\n%-5s %-20s %-15s %-15s %-10s %-10s", "---", "-----", "------------", "--------", "------", "--------")
}

func (menu *AssessmentMenuStateHandler) printAssessmentRow(assessment *model.Assessment) {
	publishDate := assessment.PublishDate.Format("2006-01-02")
	dueDate := assessment.DueDate.Format("2006-01-02")

	fmt.Printf("\n%-5d %-20s %-15s %-15s %-10s %-10v",
		assessment.AssessmentId,
		assessment.Title,
		publishDate,
		dueDate,
		assessment.Status,
		assessment.ClassId)
}

func (menu *AssessmentMenuStateHandler) printAssessmentList(assessments []*model.Assessment) {
	if len(assessments) == 0 {
		fmt.Println("\nNo assessments found.")
		return
	}

	menu.printAssessmentHeader()
	for _, assessment := range assessments {
		menu.printAssessmentRow(assessment)
	}
	fmt.Println()
}

func (menu *AssessmentMenuStateHandler) CreateAssessment() error {
	title := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assessment title:",
		FieldNameText: "Title",
	}).(string)

	description := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assessment description:",
		FieldNameText: "Description",
	}).(string)

	publishDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter publish date (YYYY-MM-DD):",
		FieldNameText: "Publish Date",
	}).(string)
	publishDate, err := time.Parse("2006-01-02", publishDateStr)
	if err != nil {
		return errors.New("Invalid date format for publish date")
	}

	dueDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter due date (YYYY-MM-DD):",
		FieldNameText: "Due Date",
	}).(string)
	dueDate, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		return errors.New("Invalid date format for due date")
	}

	classIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter class ID:",
		FieldNameText: "Class ID",
	}).(string)
	classIdUint, err := strconv.ParseUint(classIdStr, 10, 64)
	if err != nil {
		return errors.New("Invalid class ID")
	}

	instructorCode := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter instructor code:",
		FieldNameText: "Instructor Code",
	}).(string)

	assessment := &model.Assessment{
		Title:       title,
		Description: description,
		PublishDate: publishDate,
		DueDate:     dueDate,
		Status:      model.StatusDraft,
		State:       &controller.DraftState{},
	}

	// We need to fetch and set the actual Class and Instructor objects
	var classObj curriculumModel.Class
	classObj.ClassId = uint(classIdUint)
	assessment.ClassId = classObj

	var instructorObj commonModel.Instructor
	instructorObj.InstructorCode = instructorCode
	assessment.InstructorCode = instructorObj

	assessmentId, err := menu.wrapper.AssessmentController.CreateAssessment(assessment)
	if err != nil {
		return errors.New("Failed to create assessment: " + err.Error())
	}

	fmt.Printf("\nAssessment created successfully with ID: %d\n", assessmentId)
	return nil
}

func (menu *AssessmentMenuStateHandler) ViewAllAssessments() error {
	assessments, err := menu.wrapper.AssessmentController.GetAssessments()
	if err != nil {
		return errors.New("Failed to retrieve assessments: " + err.Error())
	}

	fmt.Println("\nAll Assessments:")
	menu.printAssessmentList(assessments)
	return nil
}

func (menu *AssessmentMenuStateHandler) ViewAssessmentByID() error {
	assessmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assessment ID:",
		FieldNameText: "Assessment ID",
	}).(string)

	assessmentId, err := strconv.ParseUint(assessmentIdStr, 10, 64)
	if err != nil {
		return errors.New("Invalid assessment ID")
	}

	assessment, err := menu.wrapper.AssessmentController.GetAssessment(uint(assessmentId))
	if err != nil {
		return errors.New("Failed to retrieve assessment: " + err.Error())
	}

	fmt.Printf("\nAssessment Details (ID: %d):\n", assessmentId)
	fmt.Printf("Title: %s\n", assessment.Title)
	fmt.Printf("Description: %s\n", assessment.Description)
	fmt.Printf("Publish Date: %s\n", assessment.PublishDate.Format("2006-01-02"))
	fmt.Printf("Due Date: %s\n", assessment.DueDate.Format("2006-01-02"))
	fmt.Printf("Status: %s\n", assessment.Status)
	fmt.Printf("Class ID: %v\n", assessment.ClassId)
	fmt.Printf("Instructor Code: %v\n", assessment.InstructorCode)
	return nil
}

func (menu *AssessmentMenuStateHandler) UpdateAssessment() error {
	assessmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assessment ID to update:",
		FieldNameText: "Assessment ID",
	}).(string)

	assessmentId, err := strconv.ParseUint(assessmentIdStr, 10, 64)
	if err != nil {
		return errors.New("Invalid assessment ID")
	}

	assessment, err := menu.wrapper.AssessmentController.GetAssessment(uint(assessmentId))
	if err != nil {
		return errors.New("Failed to retrieve assessment: " + err.Error())
	}

	fmt.Printf("Current title: %s\n", assessment.Title)
	title := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new title (or press Enter to keep current):",
		FieldNameText: "Title",
	}).(string)
	if title == "" {
		title = assessment.Title
	}

	fmt.Printf("Current description: %s\n", assessment.Description)
	description := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new description (or press Enter to keep current):",
		FieldNameText: "Description",
	}).(string)
	if description == "" {
		description = assessment.Description
	}

	fmt.Printf("Current publish date: %s\n", assessment.PublishDate.Format("2006-01-02"))
	publishDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new publish date (YYYY-MM-DD) (or press Enter to keep current):",
		FieldNameText: "Publish Date",
	}).(string)

	publishDate := assessment.PublishDate
	if publishDateStr != "" {
		var err error
		publishDate, err = time.Parse("2006-01-02", publishDateStr)
		if err != nil {
			return errors.New("Invalid date format for publish date")
		}
	}

	fmt.Printf("Current due date: %s\n", assessment.DueDate.Format("2006-01-02"))
	dueDateStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new due date (YYYY-MM-DD) (or press Enter to keep current):",
		FieldNameText: "Due Date",
	}).(string)

	dueDate := assessment.DueDate
	if dueDateStr != "" {
		var err error
		dueDate, err = time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			return errors.New("Invalid date format for due date")
		}
	}

	fmt.Printf("Current class ID: %v\n", assessment.ClassId)
	classIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new class ID (or press Enter to keep current):",
		FieldNameText: "Class ID",
	}).(string)

	classId := assessment.ClassId
	if classIdStr != "" {
		classIdUint, err := strconv.ParseUint(classIdStr, 10, 64)
		if err != nil {
			return errors.New("Invalid class ID")
		}

		var classObj curriculumModel.Class
		classObj.ClassId = uint(classIdUint)
		classId = classObj
	}

	fmt.Printf("Current instructor code: %v\n", assessment.InstructorCode)
	instructorCodeStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter new instructor code (or press Enter to keep current):",
		FieldNameText: "Instructor Code",
	}).(string)

	instructorCode := assessment.InstructorCode
	if instructorCodeStr != "" {
		var instructorObj commonModel.Instructor
		instructorObj.InstructorCode = instructorCodeStr
		instructorCode = instructorObj
	}

	updatedAssessment := &model.Assessment{
		AssessmentId:   uint(assessmentId),
		Title:          title,
		Description:    description,
		PublishDate:    publishDate,
		DueDate:        dueDate,
		Status:         assessment.Status,
		ClassId:        classId,
		InstructorCode: instructorCode,
		State:          assessment.State,
	}

	_, err = menu.wrapper.AssessmentController.UpdateAssessment(updatedAssessment)
	if err != nil {
		return errors.New("Failed to update assessment: " + err.Error())
	}

	fmt.Println("\nAssessment updated successfully")
	return nil
}

func (menu *AssessmentMenuStateHandler) ChangeAssessmentStatus() error {
	assessmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assessment ID:",
		FieldNameText: "Assessment ID",
	}).(string)

	assessmentId, err := strconv.ParseUint(assessmentIdStr, 10, 64)
	if err != nil {
		return errors.New("Invalid assessment ID")
	}

	assessment, err := menu.wrapper.AssessmentController.GetAssessment(uint(assessmentId))
	if err != nil {
		return errors.New("Failed to retrieve assessment: " + err.Error())
	}

	fmt.Printf("\nCurrent status: %s\n", assessment.Status)
	fmt.Println("Available status options:")
	fmt.Println("1. Draft")
	fmt.Println("2. Published")
	fmt.Println("3. Closed")

	choice := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Select new status (1-3):",
		FieldNameText: "Status Choice",
	}).(string)

	var newStatus model.AssessmentStatus
	switch choice {
	case "1":
		newStatus = model.StatusDraft
	case "2":
		newStatus = model.StatusPublished
	case "3":
		newStatus = model.StatusClosed
	default:
		return errors.New("Invalid status choice")
	}

	err = menu.wrapper.AssessmentController.UpdateAssessmentStatus(uint(assessmentId), newStatus)
	if err != nil {
		return errors.New("Failed to update assessment status: " + err.Error())
	}

	fmt.Printf("\nAssessment status updated to: %s\n", newStatus)
	return nil
}

func (menu *AssessmentMenuStateHandler) DeleteAssessment() error {
	assessmentIdStr := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter assessment ID to delete:",
		FieldNameText: "Assessment ID",
	}).(string)

	assessmentId, err := strconv.ParseUint(assessmentIdStr, 10, 64)
	if err != nil {
		return errors.New("Invalid assessment ID")
	}

	confirmation := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    fmt.Sprintf("Are you sure you want to delete assessment %d? (yes/no):", assessmentId),
		FieldNameText: "Confirmation",
	}).(string)

	if confirmation != "yes" {
		fmt.Println("\nDeletion cancelled")
		return nil
	}

	_, err = menu.wrapper.AssessmentController.DeleteAssessment(uint(assessmentId))
	if err != nil {
		return errors.New("Failed to delete assessment: " + err.Error())
	}

	fmt.Println("\nAssessment deleted successfully")
	return nil
}
