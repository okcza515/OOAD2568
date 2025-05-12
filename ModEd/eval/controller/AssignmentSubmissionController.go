package controller

import (
	"ModEd/core"
	"ModEd/eval/model"
	"fmt"

	"gorm.io/gorm"
)

type AssignmentSubmissionController struct {
	db   *gorm.DB
	core *core.BaseController[*model.AssignmentSubmission]
}

type AssignmentSubmissionControllerInterface interface {
	CreateSubmission(submission *model.AssignmentSubmission) error
	GetSubmission(submissionID uint) (*model.AssignmentSubmission, error)
	GetSubmissionByAssignmentAndStudent(assignmentID uint, studentCode string) (*model.AssignmentSubmission, error)
	GetSubmissionsByAssignment(assignmentID uint) ([]*model.AssignmentSubmission, error)
	GetSubmissionsByStudent(studentCode string) ([]*model.AssignmentSubmission, error)
	UpdateSubmission(submission *model.AssignmentSubmission) error
	DeleteSubmission(submissionID uint) error
	SubmitAssignment(submission *model.AssignmentSubmission) error
}

func NewAssignmentSubmissionController(db *gorm.DB) *AssignmentSubmissionController {
	return &AssignmentSubmissionController{
		db:   db,
		core: core.NewBaseController[*model.AssignmentSubmission](db),
	}
}

func (c *AssignmentSubmissionController) CreateSubmission(submission *model.AssignmentSubmission) error {
	// Get the assignment to check its status
	var assignment model.Assignment
	if err := c.db.Where("assignment_id = ?", submission.AssignmentId).First(&assignment).Error; err != nil {
		return fmt.Errorf("assignment not found: %v", err)
	}

	// Initialize the assignment state if it's nil
	if assignment.State == nil {
		switch assignment.Status {
		case model.StatusDraft:
			assignment.State = &DraftState{}
		case model.StatusPublished:
			assignment.State = &PublishedState{}
		case model.StatusClosed:
			assignment.State = &ClosedState{}
		default:
			assignment.State = &DraftState{}
		}
	}

	// Check if the assignment is in a state that allows submissions
	if err := assignment.State.HandleSubmission(&assignment, submission); err != nil {
		return err
	}

	return c.core.Insert(submission)
}

func (c *AssignmentSubmissionController) GetSubmission(submissionID uint) (*model.AssignmentSubmission, error) {
	return c.core.RetrieveByID(submissionID)
}

func (c *AssignmentSubmissionController) GetSubmissionByAssignmentAndStudent(assignmentID uint, studentCode string) (*model.AssignmentSubmission, error) {
	condition := map[string]interface{}{
		"assignment_id": assignmentID,
		"student_code":  studentCode,
	}

	return c.core.RetrieveByCondition(condition)
}

func (c *AssignmentSubmissionController) GetSubmissionsByAssignment(assignmentID uint) ([]*model.AssignmentSubmission, error) {
	condition := map[string]interface{}{
		"assignment_id": assignmentID,
	}

	return c.core.List(condition)
}

func (c *AssignmentSubmissionController) GetSubmissionsByStudent(studentCode string) ([]*model.AssignmentSubmission, error) {
	condition := map[string]interface{}{
		"student_code": studentCode,
	}

	submissions, err := c.core.List(condition)
	if err != nil {
		return nil, err
	}

	// Filter submissions to only get assignments that are published
	var filteredSubmissions []*model.AssignmentSubmission
	for _, submission := range submissions {
		var assignment model.Assignment
		if err := c.db.Where("assignment_id = ?", submission.AssignmentId).First(&assignment).Error; err != nil {
			continue // Skip if assignment not found
		}
		filteredSubmissions = append(filteredSubmissions, submission)
	}

	return filteredSubmissions, nil
}

func (c *AssignmentSubmissionController) UpdateSubmission(submission *model.AssignmentSubmission) error {
	// Get the original submission to check if it exists
	_, err := c.GetSubmission(submission.ID)
	if err != nil {
		return fmt.Errorf("submission not found: %v", err)
	}

	// Get the assignment to check its status
	var assignment model.Assignment
	if err := c.db.Where("assignment_id = ?", submission.AssignmentId).First(&assignment).Error; err != nil {
		return fmt.Errorf("assignment not found: %v", err)
	}

	// Initialize the assignment state if it's nil
	if assignment.State == nil {
		switch assignment.Status {
		case model.StatusDraft:
			assignment.State = &DraftState{}
		case model.StatusPublished:
			assignment.State = &PublishedState{}
		case model.StatusClosed:
			assignment.State = &ClosedState{}
		default:
			assignment.State = &DraftState{}
		}
	}

	// Check if the assignment is in a state that allows submissions
	if err := assignment.State.HandleSubmission(&assignment, submission); err != nil {
		return err
	}

	return c.core.UpdateByID(submission)
}

func (c *AssignmentSubmissionController) DeleteSubmission(submissionID uint) error {
	submission, err := c.GetSubmission(submissionID)
	if err != nil {
		return fmt.Errorf("submission not found: %v", err)
	}

	// Get the assignment to check its status
	var assignment model.Assignment
	if err := c.db.Where("assignment_id = ?", submission.AssignmentId).First(&assignment).Error; err != nil {
		return fmt.Errorf("assignment not found: %v", err)
	}

	// Initialize the assignment state if it's nil
	if assignment.State == nil {
		switch assignment.Status {
		case model.StatusDraft:
			assignment.State = &DraftState{}
		case model.StatusPublished:
			assignment.State = &PublishedState{}
		case model.StatusClosed:
			assignment.State = &ClosedState{}
		default:
			assignment.State = &DraftState{}
		}
	}

	// Check if the assignment is in a state that allows modifications
	if assignment.Status != model.StatusPublished {
		return fmt.Errorf("cannot modify submission for non-published assignment")
	}

	return c.core.DeleteByID(submissionID)
}

func (c *AssignmentSubmissionController) SubmitAssignment(submission *model.AssignmentSubmission) error {
	// Get the assignment to check its status
	var assignment model.Assignment
	if err := c.db.Where("assignment_id = ?", submission.AssignmentId).First(&assignment).Error; err != nil {
		return fmt.Errorf("assignment not found: %v", err)
	}

	// Initialize the assignment state if it's nil
	if assignment.State == nil {
		switch assignment.Status {
		case model.StatusDraft:
			assignment.State = &DraftState{}
		case model.StatusPublished:
			assignment.State = &PublishedState{}
		case model.StatusClosed:
			assignment.State = &ClosedState{}
		default:
			assignment.State = &DraftState{}
		}
	}

	// Check if the assignment is in a state that allows submissions
	if err := assignment.State.HandleSubmission(&assignment, submission); err != nil {
		return err
	}

	// Set the submission as submitted
	submission.Submitted = true

	// Check if the submission already exists
	existingSubmission, err := c.GetSubmissionByAssignmentAndStudent(submission.AssignmentId, submission.StudentCode)
	if err == nil {
		// Update the existing submission
		existingSubmission.Answer = submission.Answer
		existingSubmission.Submitted = true
		return c.core.UpdateByID(existingSubmission)
	} else {
		// Create a new submission
		return c.core.Insert(submission)
	}
}
