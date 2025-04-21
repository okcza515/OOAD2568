package controller

import (
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type ScoreAdvisorController struct {
	db *gorm.DB
}

func NewScoreAdvisorController(db *gorm.DB) *ScoreAdvisorController {
	return &ScoreAdvisorController{db: db}
}

func (c *ScoreAdvisorController) ListAllAdvisorScores(scoreType string) (interface{}, error) {
	var scores interface{}
	switch scoreType {
	case "assignment":
		scores = &[]model.ScoreAssignmentAdvisor{}
	case "presentation":
		scores = &[]model.ScorePresentationAdvisor{}
	case "report":
		scores = &[]model.ScoreReportAdvisor{}
	case "assessment":
		scores = &[]model.ScoreAssessmentAdvisor{}
	default:
		return nil, fmt.Errorf("invalid score type: %s", scoreType)
	}

	err := c.db.Find(scores).Error
	return scores, err
}

func (c *ScoreAdvisorController) RetrieveAdvisorScore(scoreType string, id uint) (interface{}, error) {
	var score interface{}
	switch scoreType {
	case "assignment":
		score = &model.ScoreAssignmentAdvisor{}
	case "presentation":
		score = &model.ScorePresentationAdvisor{}
	case "report":
		score = &model.ScoreReportAdvisor{}
	case "assessment":
		score = &model.ScoreAssessmentAdvisor{}
	default:
		return nil, fmt.Errorf("invalid score type: %s", scoreType)
	}

	err := c.db.First(score, id).Error
	return score, err
}

func (c *ScoreAdvisorController) InsertAdvisorScore(score interface{}) error {
	return c.db.Create(score).Error
}

func (c *ScoreAdvisorController) UpdateAdvisorScore(score interface{}) error {
	return c.db.Save(score).Error
}

func (c *ScoreAdvisorController) DeleteAdvisorScore(score interface{}, id uint) error {
	if err := c.db.Where("id = ?", id).First(score).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}
	return c.db.Delete(score).Error
}
