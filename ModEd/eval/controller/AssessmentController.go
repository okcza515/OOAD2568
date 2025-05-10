// MEP-1002
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	// "time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AssessmentController struct {
	db   *gorm.DB
	core *core.BaseController[*model.Assessment]
}

type AssessmentControllerInterface interface {
	CreateAssessment(assessment *model.Assessment) (assessmentId uint, err error)
	GetAssessment(assessmentId uint, preload ...string) (assessment *model.Assessment, err error)
	GetAssessments(preload ...string) (assessments []*model.Assessment, err error)
	GetAssessmentsByClass(classId uint, preload ...string) (assessments []*model.Assessment, err error)
	GetAssessmentsByInstructor(instructorCode string, preload ...string) (assessments []*model.Assessment, err error)
	UpdateAssessment(updatedAssessment *model.Assessment) (*model.Assessment, error)
	DeleteAssessment(assessmentId uint) (assessment *model.Assessment, err error)
	UpdateAssessmentStatus(assessmentID uint, newStatus model.AssessmentStatus) error
	// CreateSubmission(assessmentId uint, submission *model.AssessmentSubmission) error
	// GetSubmissionsByAssessment(assessmentId uint) ([]model.AssessmentSubmission, error)
	// UpdateSubmission(submission *model.AssessmentSubmission) error
}

func NewAssessmentController(db *gorm.DB) AssessmentControllerInterface {
	return &AssessmentController{
		db:   db,
		core: core.NewBaseController[*model.Assessment](db),
	}
}

func (c *AssessmentController) CreateAssessment(assessment *model.Assessment) (assessmentId uint, err error) {
	if err := c.core.Insert(assessment); err != nil {
		return 0, err
	}
	return assessment.AssessmentId, nil
}

func (c *AssessmentController) GetAssessment(assessmentId uint, preload ...string) (assessment *model.Assessment, err error) {
	assessment, err = c.core.RetrieveByCondition(map[string]interface{}{"assessment_id": assessmentId}, preload...)
	if err != nil {
		return nil, err
	}
	return assessment, nil
}

func (c *AssessmentController) GetAssessments(preload ...string) (assessments []*model.Assessment, err error) {
	assessments, err = c.core.List(nil, preload...)
	if err != nil {
		return nil, err
	}
	return assessments, nil
}

func (c *AssessmentController) GetAssessmentsByClass(classId uint, preload ...string) (assessments []*model.Assessment, err error) {
	condition := map[string]interface{}{"class_id": classId}
	assessments, err = c.core.List(condition, preload...)
	if err != nil {
		return nil, err
	}
	return assessments, nil
}

func (c *AssessmentController) GetAssessmentsByInstructor(instructorCode string, preload ...string) (assessments []*model.Assessment, err error) {
	condition := map[string]interface{}{"instructor_code": instructorCode}
	assessments, err = c.core.List(condition, preload...)
	if err != nil {
		return nil, err
	}
	return assessments, nil
}

func (c *AssessmentController) UpdateAssessment(updatedAssessment *model.Assessment) (assessment *model.Assessment, err error) {
	assessment, err = c.core.RetrieveByCondition(map[string]interface{}{"assessment_id": updatedAssessment.AssessmentId})
	if err != nil {
		return nil, err
	}

	assessment.Title = updatedAssessment.Title
	assessment.Description = updatedAssessment.Description
	assessment.PublishDate = updatedAssessment.PublishDate
	assessment.DueDate = updatedAssessment.DueDate
	assessment.Status = updatedAssessment.Status
	assessment.ClassId = updatedAssessment.ClassId
	assessment.InstructorCode = updatedAssessment.InstructorCode

	if err := c.core.UpdateByCondition(map[string]interface{}{"assessment_id": updatedAssessment.AssessmentId}, assessment); err != nil {
		return nil, err
	}
	return assessment, nil
}

func (c *AssessmentController) DeleteAssessment(assessmentId uint) (assessment *model.Assessment, err error) {
	assessment, err = c.core.RetrieveByCondition(map[string]interface{}{"assessment_id": assessmentId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"assessment_id": assessmentId}); err != nil {
		return nil, err
	}
	return assessment, nil
}

func (c *AssessmentController) UpdateAssessmentStatus(assessmentID uint, newStatus model.AssessmentStatus) error {
	assessment, err := c.GetAssessment(assessmentID)
	if err != nil {
		return err
	}

	if assessment.State == nil {
		return errors.New("assessment state is not initialized")
	}

	return assessment.State.HandleStatusChange(assessment, newStatus)
}

// func (c *AssessmentController) CreateSubmission(assessmentId uint, submission *model.AssessmentSubmission) error {
// 	// Check if the assessment exists
// 	_, err := c.GetAssessment(assessmentId, "Submission")
// 	if err != nil {
// 		return err
// 	}

// 	submission.UpdatedAt = time.Now()

// 	// Add submission to assessment
// 	return c.db.Create(submission).Error
// }

// func (c *AssessmentController) GetSubmissionsByAssessment(assessmentId uint) ([]model.AssessmentSubmission, error) {
// 	var submissions []model.AssessmentSubmission
// 	err := c.db.Where("assessment_id = ?", assessmentId).Find(&submissions).Error
// 	return submissions, err
// }

// func (c *AssessmentController) UpdateSubmission(submission *model.AssessmentSubmission) error {
// 	submission.UpdatedAt = time.Now()
// 	return c.db.Save(submission).Error
// }
