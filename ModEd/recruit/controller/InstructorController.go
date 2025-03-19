// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InstructorController struct {
	DB *gorm.DB
}

func CreateInstructorController(db *gorm.DB) *InstructorController {
	return &InstructorController{DB: db}
}

// Get all interviews assigned to an instructor
func (ctrl *InstructorController) GetInterviewsByInstructor(instructorID uint) ([]model.Interview, error) {
	var interviews []model.Interview
	err := ctrl.DB.Where("instructor_id = ?", instructorID).Find(&interviews).Error
	return interviews, err
}

// Evaluate an applicant by updating interview score
func (ctrl *InstructorController) EvaluateApplicant(interviewID uuid.UUID, score float64) error {
	return ctrl.DB.Model(&model.Interview{}).
		Where("id = ?", interviewID).
		Update("interview_score", score).Error
}
