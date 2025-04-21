package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type ScoreAdvisorController struct {
	*core.BaseController
	db *gorm.DB
}

func NewScoreAdvisorController(db *gorm.DB) *ScoreAdvisorController {
	return &ScoreAdvisorController{
		db:             db,
		BaseController: core.NewBaseController("score_advisors", db),
	}
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
	// Ensure the score implements core.RecordInterface
	if record, ok := score.(core.RecordInterface); ok {
		return c.Insert(record)
	}
	return fmt.Errorf("score does not implement core.RecordInterface")
}

func (c *ScoreAdvisorController) UpdateAdvisorScore(scoreType string, score interface{}) error {
	switch scoreType {
	case "assignment":
		if s, ok := score.(*model.ScoreAssignmentAdvisor); ok {
			return c.UpdateByID(s)
		}
	case "assessment":
		if s, ok := score.(*model.ScoreAssessmentAdvisor); ok {
			return c.UpdateByID(s)
		}
	case "presentation":
		if s, ok := score.(*model.ScorePresentationAdvisor); ok {
			return c.UpdateByID(s)
		}
	case "report":
		if s, ok := score.(*model.ScoreReportAdvisor); ok {
			return c.UpdateByID(s)
		}
	default:
		return fmt.Errorf("invalid score type: %s", scoreType)
	}
	return fmt.Errorf("failed to cast score to the correct type for scoreType: %s", scoreType)
}

func (c *ScoreAdvisorController) DeleteAdvisorScore(scoreType string, score interface{}) error {
	switch scoreType {
	case "assignment":
		if s, ok := score.(*model.ScoreAssignmentAdvisor); ok {
			return c.db.Where("id = ?", s.ID).Delete(&model.ScoreAssignmentAdvisor{}).Error
		}
	case "assessment":
		if s, ok := score.(*model.ScoreAssessmentAdvisor); ok {
			return c.db.Where("id = ?", s.ID).Delete(&model.ScoreAssessmentAdvisor{}).Error
		}
	case "presentation":
		if s, ok := score.(*model.ScorePresentationAdvisor); ok {
			return c.db.Where("id = ?", s.ID).Delete(&model.ScorePresentationAdvisor{}).Error
		}
	case "report":
		if s, ok := score.(*model.ScoreReportAdvisor); ok {
			return c.db.Where("id = ?", s.ID).Delete(&model.ScoreReportAdvisor{}).Error
		}
	default:
		return fmt.Errorf("invalid score type: %s", scoreType)
	}
	return fmt.Errorf("failed to cast score to the correct type for scoreType: %s", scoreType)
}
