package controller

import (
	"errors"

	assessmentModel "ModEd/eval/model"

	"gorm.io/gorm"
)

type AssessmentController struct {
	DB *gorm.DB
}

func NewAssessmentController(db *gorm.DB) *AssessmentController {
	return &AssessmentController{
		DB: db,
	}
}

// CreateAssessment creates a new assessment
func (c *AssessmentController) CreateAssessment(assessment *assessmentModel.Assessment) error {
	return c.DB.Create(assessment).Error
}

// GetAssessments retrieves all assessments
func (c *AssessmentController) GetAssessments() ([]assessmentModel.Assessment, error) {
	var assessments []assessmentModel.Assessment
	err := c.DB.Find(&assessments).Error
	return assessments, err
}

// GetAssessmentByID retrieves an assessment by its ID
func (c *AssessmentController) GetAssessmentByID(id uint) (*assessmentModel.Assessment, error) {
	var assessment assessmentModel.Assessment
	err := c.DB.First(&assessment, id).Error
	if err != nil {
		return nil, err
	}
	return &assessment, nil
}

// UpdateAssessment updates an existing assessment
func (c *AssessmentController) UpdateAssessment(assessment *assessmentModel.Assessment) error {
	return c.DB.Save(assessment).Error
}

// UpdateAssessmentStatus changes the status of an assessment
func (c *AssessmentController) UpdateAssessmentStatus(assessmentID uint, newStatus assessmentModel.AssessmentStatus) error {
	assessment, err := c.GetAssessmentByID(assessmentID)
	if err != nil {
		return err
	}

	if assessment.State == nil {
		return errors.New("assessment state is not initialized")
	}

	return assessment.State.HandleStatusChange(assessment, newStatus)
}

// DeleteAssessment deletes an assessment by its ID
func (c *AssessmentController) DeleteAssessment(id uint) error {
	return c.DB.Delete(&assessmentModel.Assessment{}, id).Error
}

// CreateSubmission creates a new assessment submission
func (c *AssessmentController) CreateSubmission(submission *assessmentModel.AssessmentSubmission) error {
	return c.DB.Create(submission).Error
}

// GetSubmissionsByAssessmentID retrieves all submissions for a specific assessment
func (c *AssessmentController) GetSubmissionsByAssessmentID(assessmentID uint) ([]assessmentModel.AssessmentSubmission, error) {
	var submissions []assessmentModel.AssessmentSubmission
	err := c.DB.Where("assessment_id = ?", assessmentID).Find(&submissions).Error
	return submissions, err
}

// UpdateSubmission updates an existing submission
func (c *AssessmentController) UpdateSubmission(submission *assessmentModel.AssessmentSubmission) error {
	return c.DB.Save(submission).Error
}

// SubmitAssessment handles a new submission for an assessment
func (c *AssessmentController) SubmitAssessment(assessmentID uint, submission *assessmentModel.AssessmentSubmission) error {
	assessment, err := c.GetAssessmentByID(assessmentID)
	if err != nil {
		return err
	}

	if assessment.State == nil {
		return errors.New("assessment state is not initialized")
	}

	// Set submission as submitted
	submission.Submitted = true

	// Handle submission based on assessment state
	err = assessment.State.HandleSubmission(assessment, submission)
	if err != nil {
		return err
	}

	// Save the submission
	return c.DB.Save(submission).Error
}

// SavePathFile saves the file information for a submission
func (c *AssessmentController) SavePathFile(submissionID uint, pathFile *assessmentModel.PathFile) error {
	var submission assessmentModel.AssessmentSubmission
	err := c.DB.First(&submission, submissionID).Error
	if err != nil {
		return err
	}

	submission.Answers = *pathFile
	return c.DB.Save(&submission).Error
}

// // ScoreSubmission updates the score and feedback for a submission
// func (c *AssessmentController) ScoreSubmission(submissionID uint, score float64, feedback string) error {
// 	var submission assessmentModel.AssessmentSubmission
// 	err := c.DB.First(&submission, submissionID).Error
// 	if err != nil {
// 		return err
// 	}

// 	submission.Score = score
// 	submission.Feedback = feedback
// 	return c.DB.Save(&submission).Error
// }
