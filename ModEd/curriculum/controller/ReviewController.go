// MEP-1009 Student Internship
package controller

import (
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ReviewController struct {
	*BaseScoreController[model.SupervisorReview]
}

func NewReviewController(connector *gorm.DB) *ReviewController {
	return &ReviewController{
		BaseScoreController: &BaseScoreController[model.SupervisorReview]{Connector: connector},
	}
}

func (rc *ReviewController) UpdateReviewScore(studentID string, supervisorScore int, mentorScore int) error {
	scoreFields := map[string]interface{}{
		"InstructorScore": supervisorScore,
		"MentorScore":     mentorScore,
	}

	return rc.UpdateScore(studentID, scoreFields, func(db *gorm.DB, studentID string) (uint, error) {
		var application model.InternshipApplication
		if err := db.Where("student_code = ?", studentID).Last(&application).Error; err != nil {
			return 0, err
		}
		return application.SupervisorReviewId, nil
	})
}
