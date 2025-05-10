package controller

import (
	"errors"

	"ModEd/core"
	assessmentModel "ModEd/eval/model"

	"gorm.io/gorm"
)

type AssessmentController struct {
	*core.BaseController[assessmentModel.Assessment]
	db *gorm.DB
}

func NewAssessmentController(db *gorm.DB) *AssessmentController {
	return &AssessmentController{
		db:             db,
		BaseController: core.NewBaseController[assessmentModel.Assessment](db),
	}
}

// CreateAssessment creates a new assessment
func (c *AssessmentController) CreateAssessment(assessment *assessmentModel.Assessment) error {
	return c.db.Create(assessment).Error
}

// GetAssessments retrieves all assessments
func (c *AssessmentController) GetAssessments() ([]assessmentModel.Assessment, error) {
	var assessments []assessmentModel.Assessment
	err := c.db.Find(&assessments).Error
	return assessments, err
}

// GetAssessmentByID retrieves an assessment by its ID
func (c *AssessmentController) GetAssessmentByID(id uint) (*assessmentModel.Assessment, error) {
	var assessment assessmentModel.Assessment
	err := c.db.First(&assessment, id).Error
	if err != nil {
		return nil, err
	}
	return &assessment, nil
}

// UpdateAssessment updates an existing assessment
func (c *AssessmentController) UpdateAssessment(assessment *assessmentModel.Assessment) error {
	return c.db.Save(assessment).Error
}

// UpdateAssessmentStatus changes the status of an assessment
func (c *AssessmentController) UpdateAssessmentStatus(assessmentID uint, newStatus assessmentModel.AssessmentStatus) error {
	assessment, err := c.RetrieveByID(assessmentID)
	if err != nil {
		return err
	}

	if assessment.State == nil {
		return errors.New("assessment state is not initialized")
	}

	return assessment.State.HandleStatusChange(&assessment, newStatus)
}

// DeleteAssessment deletes an assessment by its ID
func (c *AssessmentController) DeleteAssessment(id uint) error {
	return c.db.Delete(&assessmentModel.Assessment{}, id).Error
}

// CreateSubmission creates a new assessment submission
func (c *AssessmentController) CreateSubmission(submission *assessmentModel.AssessmentSubmission) error {
	return c.db.Create(submission).Error
}

// GetSubmissionsByAssessmentID retrieves all submissions for a specific assessment
func (c *AssessmentController) GetSubmissionsByAssessmentID(assessmentID uint) ([]assessmentModel.AssessmentSubmission, error) {
	var submissions []assessmentModel.AssessmentSubmission
	err := c.db.Where("assessment_id = ?", assessmentID).Find(&submissions).Error
	return submissions, err
}

// UpdateSubmission updates an existing submission
func (c *AssessmentController) UpdateSubmission(submission *assessmentModel.AssessmentSubmission) error {
	return c.db.Save(submission).Error
}

// SubmitAssessment handles a new submission for an assessment
func (c *AssessmentController) SubmitAssessment(assessmentID uint, submission *assessmentModel.AssessmentSubmission) error {
	assessment, err := c.RetrieveByID(assessmentID)
	if err != nil {
		return err
	}

	if assessment.State == nil {
		return errors.New("assessment state is not initialized")
	}

	// Set submission as submitted
	submission.Submitted = true

	// Handle submission based on assessment state
	err = assessment.State.HandleSubmission(&assessment, submission)
	if err != nil {
		return err
	}

	// Save the submission
	return c.db.Save(submission).Error
}

// SavePathFile saves the file information for a submission
func (c *AssessmentController) SavePathFile(submissionID uint, pathFile *assessmentModel.PathFile) error {
	var submission assessmentModel.AssessmentSubmission
	err := c.db.First(&submission, submissionID).Error
	if err != nil {
		return err
	}

	submission.Answers = *pathFile
	return c.db.Save(&submission).Error
}

// // ScoreSubmission updates the score and feedback for a submission
// func (c *AssessmentController) ScoreSubmission(submissionID uint, score float64, feedback string) error {
// 	var submission assessmentModel.AssessmentSubmission
// 	err := c.db.First(&submission, submissionID).Error
// 	if err != nil {
// 		return err
// 	}

// 	submission.Score = score
// 	submission.Feedback = feedback
// 	return c.DB.Save(&submission).Error
// }
