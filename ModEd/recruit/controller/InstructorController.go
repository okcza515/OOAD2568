// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"
	"fmt"

	"gorm.io/gorm"
)

type InstructorController struct {
	*core.BaseController
	DB *gorm.DB
}

func CreateInstructorController(db *gorm.DB) *InstructorController {
	return &InstructorController{
		BaseController: core.NewBaseController("Interview", db),
		DB:             db,
	}
}

func (ctrl *InstructorController) GetInterviewsByInstructor(instructorID uint) ([]model.Interview, error) {
	var interviews []model.Interview
	err := ctrl.DB.Where("instructor_id = ?", instructorID).Find(&interviews).Error
	return interviews, err
}

func (ctrl *InstructorController) EvaluateApplicant(interviewID uint, score float64) error {
	result := ctrl.DB.Model(&model.Interview{}).
		Where("id = ?", interviewID).
		Update("interview_score", score)

	if result.RowsAffected == 0 {
		return fmt.Errorf("no interview found with ID %d", interviewID)
	}
	return result.Error
}
