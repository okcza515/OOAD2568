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

func (c *AssessmentController) CreateAssessment(assessment *assessmentModel.Assessment) error {
	return c.db.Create(assessment).Error
}

func (c *AssessmentController) GetAssessments() ([]assessmentModel.Assessment, error) {
	var assessments []assessmentModel.Assessment
	err := c.db.Find(&assessments).Error
	return assessments, err
}

func (c *AssessmentController) GetAssessmentByID(id uint) (*assessmentModel.Assessment, error) {
	var assessment assessmentModel.Assessment
	err := c.db.First(&assessment, id).Error
	if err != nil {
		return nil, err
	}
	return &assessment, nil
}

func (c *AssessmentController) UpdateAssessment(assessment *assessmentModel.Assessment) error {
	return c.db.Save(assessment).Error
}

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

func (c *AssessmentController) DeleteAssessment(id uint) error {
	return c.db.Delete(&assessmentModel.Assessment{}, id).Error
}

// อาจจะไม่มี
func (c *AssessmentController) CreateSubmission(submission *assessmentModel.AssessmentSubmission) error {
	return c.db.Create(submission).Error
}

// อาจจะไม่มี
func (c *AssessmentController) GetSubmissionsByAssessmentID(assessmentID uint) ([]assessmentModel.AssessmentSubmission, error) {
	var submissions []assessmentModel.AssessmentSubmission
	err := c.db.Where("assessment_id = ?", assessmentID).Find(&submissions).Error
	return submissions, err
}

// อาจจะไม่มี
func (c *AssessmentController) UpdateSubmission(submission *assessmentModel.AssessmentSubmission) error {
	return c.db.Save(submission).Error
}

// อาจจะไม่มี
func (c *AssessmentController) SubmitAssessment(assessmentID uint, submission *assessmentModel.AssessmentSubmission) error {
	assessment, err := c.RetrieveByID(assessmentID)
	if err != nil {
		return err
	}

	if assessment.State == nil {
		return errors.New("assessment state is not initialized")
	}

	submission.Submitted = true

	err = assessment.State.HandleSubmission(&assessment, submission)
	if err != nil {
		return err
	}

	return c.db.Save(submission).Error
}

// อาจจะไม่มี
func (c *AssessmentController) SavePathFile(submissionID uint, pathFile *assessmentModel.PathFile) error {
	var submission assessmentModel.AssessmentSubmission
	err := c.db.First(&submission, submissionID).Error
	if err != nil {
		return err
	}

	submission.Answers = *pathFile
	return c.db.Save(&submission).Error
}
