package controller

import (
	"fmt"
	"time"

	assessmentModel "ModEd/eval/model"
)

type DraftState struct{}

func (s *DraftState) HandleSubmission(assessment *assessmentModel.Assessment, submission *assessmentModel.AssessmentSubmission) error {
	return fmt.Errorf("cannot submit to a draft assessment")
}

func (s *DraftState) HandleStatusChange(assessment *assessmentModel.Assessment, newStatus assessmentModel.AssessmentStatus) error {
	if newStatus == assessmentModel.StatusPublished {
		assessment.Status = newStatus
		assessment.State = &PublishedState{}
		return nil
	}
	return fmt.Errorf("invalid status transition from draft")
}

type PublishedState struct{}

func (s *PublishedState) HandleSubmission(assessment *assessmentModel.Assessment, submission *assessmentModel.AssessmentSubmission) error {
	if time.Now().After(assessment.DueDate) {
		return fmt.Errorf("submission deadline has passed")
	}
	return nil
}

func (s *PublishedState) HandleStatusChange(assessment *assessmentModel.Assessment, newStatus assessmentModel.AssessmentStatus) error {
	if newStatus == assessmentModel.StatusClosed {
		assessment.Status = newStatus
		assessment.State = &ClosedState{}
		return nil
	}
	return fmt.Errorf("invalid status transition from published")
}

type ClosedState struct{}

func (s *ClosedState) HandleSubmission(assessment *assessmentModel.Assessment, submission *assessmentModel.AssessmentSubmission) error {
	return fmt.Errorf("cannot submit to a closed assessment")
}

func (s *ClosedState) HandleStatusChange(assessment *assessmentModel.Assessment, newStatus assessmentModel.AssessmentStatus) error {
	return fmt.Errorf("cannot change status of a closed assessment")
}
