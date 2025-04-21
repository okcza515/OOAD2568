package controller

import (
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type ScoreCommitteeController struct {
	db *gorm.DB
}

func NewScoreCommitteeController(db *gorm.DB) *ScoreCommitteeController {
	return &ScoreCommitteeController{db: db}
}

func (c *ScoreCommitteeController) ListAllCommitteeScores(scoreType string) (interface{}, error) {
	var scores interface{}
	switch scoreType {
	case "assignment":
		scores = &[]model.ScoreAssignmentCommittee{}
	case "presentation":
		scores = &[]model.ScorePresentationCommittee{}
	case "report":
		scores = &[]model.ScoreReportCommittee{}
	case "assessment":
		scores = &[]model.ScoreAssessmentCommittee{}
	default:
		return nil, fmt.Errorf("invalid score type: %s", scoreType)
	}

	err := c.db.Find(scores).Error
	return scores, err
}

func (c *ScoreCommitteeController) RetrieveCommitteeScore(scoreType string, id uint) (interface{}, error) {
	var score interface{}
	switch scoreType {
	case "assignment":
		score = &model.ScoreAssignmentCommittee{}
	case "presentation":
		score = &model.ScorePresentationCommittee{}
	case "report":
		score = &model.ScoreReportCommittee{}
	case "assessment":
		score = &model.ScoreAssessmentCommittee{}
	default:
		return nil, fmt.Errorf("invalid score type: %s", scoreType)
	}

	err := c.db.First(score, id).Error
	return score, err
}

func (c *ScoreCommitteeController) InsertCommitteeScore(score interface{}) error {
	return c.db.Create(score).Error
}

func (c *ScoreCommitteeController) UpdateCommitteeScore(score interface{}) error {
	return c.db.Save(score).Error
}

func (c *ScoreCommitteeController) DeleteCommitteeScore(score interface{}, id uint) error {
	if err := c.db.Where("id = ?", id).First(score).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}
	return c.db.Delete(score).Error
}
