//MEP-1006

package controller

import (
	"fmt"
	"time"

	assignmentModel "ModEd/eval/model"
)

type DraftState struct{}

func (s *DraftState) HandleSubmission(assignment *assignmentModel.Assignment, submission *assignmentModel.AssignmentSubmission) error {
	return fmt.Errorf("cannot submit to a draft assignment")
}

func (s *DraftState) HandleStatusChange(assignment *assignmentModel.Assignment, newStatus assignmentModel.AssignmentStatus) error {
	if newStatus == assignmentModel.StatusPublished {
		assignment.Status = newStatus
		assignment.State = &PublishedState{}
		return nil
	}
	return fmt.Errorf("invalid status transition from draft")
}

type PublishedState struct{}

func (s *PublishedState) HandleSubmission(assignment *assignmentModel.Assignment, submission *assignmentModel.AssignmentSubmission) error {
	if time.Now().After(assignment.DueDate) {
		return fmt.Errorf("submission deadline has passed")
	}
	return nil
}

func (s *PublishedState) HandleStatusChange(assignment *assignmentModel.Assignment, newStatus assignmentModel.AssignmentStatus) error {
	if newStatus == assignmentModel.StatusClosed {
		assignment.Status = newStatus
		assignment.State = &ClosedState{}
		return nil
	}
	return fmt.Errorf("invalid status transition from published")
}

type ClosedState struct{}

func (s *ClosedState) HandleSubmission(assignment *assignmentModel.Assignment, submission *assignmentModel.AssignmentSubmission) error {
	return fmt.Errorf("cannot submit to a closed assignment")
}

func (s *ClosedState) HandleStatusChange(assignment *assignmentModel.Assignment, newStatus assignmentModel.AssignmentStatus) error {
	return fmt.Errorf("cannot change status of a closed assignment")
}
