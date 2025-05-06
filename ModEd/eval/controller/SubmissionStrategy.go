package controller

import (
	"fmt"

	assessmentModel "ModEd/eval/model"
)

type QuizSubmissionStrategy struct{}

func (s *QuizSubmissionStrategy) ValidateSubmission(submission *assessmentModel.AssessmentSubmission) error {
	if submission.Answers == "" {
		return fmt.Errorf("quiz answers cannot be empty")
	}
	return nil
}

func (s *QuizSubmissionStrategy) ProcessSubmission(submission *assessmentModel.AssessmentSubmission) error {
	return nil
}

type AssignmentSubmissionStrategy struct{}

func (s *AssignmentSubmissionStrategy) ValidateSubmission(submission *assessmentModel.AssessmentSubmission) error {
	if submission.Answers == "" {
		return fmt.Errorf("assignment submission cannot be empty")
	}
	return nil
}

func (s *AssignmentSubmissionStrategy) ProcessSubmission(submission *assessmentModel.AssessmentSubmission) error {
	return nil
}
