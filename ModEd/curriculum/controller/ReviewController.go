// MEP-1009 Student Internship
package controller

import (
	"ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type ReviewController struct {
	*BaseScoreController[model.SupervisorReview]
	Fetcher *Fetcher
}

func NewReviewController(connector *gorm.DB) *ReviewController {
	return &ReviewController{
		BaseScoreController: &BaseScoreController[model.SupervisorReview]{Connector: connector},
		Fetcher:             NewFetcher(connector),
	}
}

func (rc *ReviewController) UpdateReviewScore(studentID string, supervisorScore int, mentorScore int) error {
	reviewID, err := rc.Fetcher.FetchIDByStudentID(studentID)
	if err != nil {
		return fmt.Errorf("failed to fetch review ID for student '%s': %w", studentID, err)
	}

	scoreFields := map[string]interface{}{
		"InstructorScore": supervisorScore,
		"MentorScore":     mentorScore,
	}

	return rc.UpdateScoreByID(reviewID, scoreFields)
}
