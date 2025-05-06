package controller

import (
	"fmt"
	"time"

	assessmentModel "ModEd/eval/model"
)

type AssessmentController struct {
	assessments []*assessmentModel.Assessment
}

func NewAssessmentController() *AssessmentController {
	return &AssessmentController{
		assessments: make([]*assessmentModel.Assessment, 0),
	}
}

func (c *AssessmentController) CreateAssessment(assessmentType assessmentModel.AssessmentType) assessmentModel.AssessmentBuilder {
	return assessmentModel.NewAssessmentBuilder(assessmentType)
}

func (c *AssessmentController) AddAssessment(assessment *assessmentModel.Assessment) {
	c.assessments = append(c.assessments, assessment)
}

func (c *AssessmentController) GetAssessmentByID(id uint) (*assessmentModel.Assessment, error) {
	for _, a := range c.assessments {
		if a.AssessmentId == id {
			return a, nil
		}
	}
	return nil, fmt.Errorf("assessment not found")
}

func (c *AssessmentController) UpdateAssessmentStatus(assessment *assessmentModel.Assessment, newStatus assessmentModel.AssessmentStatus) error {
	if assessment.State == nil {
		assessment.State = &DraftState{}
	}
	if err := assessment.State.HandleStatusChange(assessment, newStatus); err != nil {
		return err
	}
	for _, observer := range assessment.Observers {
		observer.OnStatusChanged(assessment, assessment.Status)
	}
	return nil
}

func (c *AssessmentController) SubmitAssessment(assessment *assessmentModel.Assessment, submission *assessmentModel.AssessmentSubmission) error {
	if assessment.State == nil {
		assessment.State = &DraftState{}
	}
	return assessment.State.HandleSubmission(assessment, submission)
}

func (c *AssessmentController) AddObserver(assessment *assessmentModel.Assessment, observer assessmentModel.AssessmentObserver) {
	if assessment.Observers == nil {
		assessment.Observers = make([]assessmentModel.AssessmentObserver, 0)
	}
	assessment.Observers = append(assessment.Observers, observer)
}

func (c *AssessmentController) CreateTimedAssessment(assessment *assessmentModel.Assessment, timeLimit time.Duration) assessmentModel.AssessmentDecorator {
	return assessmentModel.NewTimedAssessmentDecorator(assessment, timeLimit)
}

func (c *AssessmentController) ValidateQuizSubmission(submission *assessmentModel.AssessmentSubmission) error {
	strategy := &QuizSubmissionStrategy{}
	return strategy.ValidateSubmission(submission)
}

func (c *AssessmentController) ValidateAssignmentSubmission(submission *assessmentModel.AssessmentSubmission) error {
	strategy := &AssignmentSubmissionStrategy{}
	return strategy.ValidateSubmission(submission)
}

func (c *AssessmentController) ProcessQuizSubmission(submission *assessmentModel.AssessmentSubmission) error {
	strategy := &QuizSubmissionStrategy{}
	return strategy.ProcessSubmission(submission)
}

func (c *AssessmentController) ProcessAssignmentSubmission(submission *assessmentModel.AssessmentSubmission) error {
	strategy := &AssignmentSubmissionStrategy{}
	return strategy.ProcessSubmission(submission)
}
