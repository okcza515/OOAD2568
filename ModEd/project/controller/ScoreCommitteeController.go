package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type ScoreCommitteeController struct {
	*core.BaseController
	db *gorm.DB
}

func NewScoreCommitteeController(db *gorm.DB) *ScoreCommitteeController {
	return &ScoreCommitteeController{
		db:             db,
		BaseController: core.NewBaseController("score_committees", db),
	}
}

func (c *ScoreCommitteeController) ListAllCommitteeScores(scoreType string) (interface{}, error) {
	var scores interface{}
	switch scoreType {
	case "assignment":
		scores = &[]model.ScoreAssignmentCommittee{}
	case "assessment":
		scores = &[]model.ScoreAssessmentCommittee{}
	case "presentation":
		scores = &[]model.ScorePresentationCommittee{}
	case "report":
		scores = &[]model.ScoreReportCommittee{}
	default:
		return nil, fmt.Errorf("invalid score type: %s", scoreType)
	}

	err := c.db.Find(scores).Error
	return scores, err
}

func (c *ScoreCommitteeController) ListCommitteeScoresByCondition(scoreType string, condition string, value interface{}) (interface{}, error) {
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

	err := c.db.Where(condition, value).Find(scores).Error
	return scores, err
}

func (c *ScoreCommitteeController) RetrieveCommitteeScore(scoreType string, id uint) (interface{}, error) {
	var score interface{}
	switch scoreType {
	case "assignment":
		score = &model.ScoreAssignmentCommittee{}
	case "assessment":
		score = &model.ScoreAssessmentCommittee{}
	case "presentation":
		score = &model.ScorePresentationCommittee{}
	case "report":
		score = &model.ScoreReportCommittee{}
	default:
		return nil, fmt.Errorf("invalid score type: %s", scoreType)
	}

	err := c.db.First(score, id).Error
	return score, err
}

func (c *ScoreCommitteeController) InsertCommitteeScore(score interface{}) error {
	// Ensure the score implements core.RecordInterface
	if record, ok := score.(core.RecordInterface); ok {
		return c.Insert(record)
	}
	return fmt.Errorf("score does not implement core.RecordInterface")
}

func (c *ScoreCommitteeController) UpdateCommitteeScore(scoreType string, score interface{}) error {
	switch scoreType {
	case "assignment":
		if s, ok := score.(*model.ScoreAssignmentCommittee); ok {
			return c.UpdateByID(s)
		}
	case "assessment":
		if s, ok := score.(*model.ScoreAssessmentCommittee); ok {
			return c.UpdateByID(s)
		}
	case "presentation":
		if s, ok := score.(*model.ScorePresentationCommittee); ok {
			return c.UpdateByID(s)
		}
	case "report":
		if s, ok := score.(*model.ScoreReportCommittee); ok {
			return c.UpdateByID(s)
		}
	default:
		return fmt.Errorf("invalid score type: %s", scoreType)
	}
	return fmt.Errorf("failed to cast score to the correct type for scoreType: %s", scoreType)
}

func (c *ScoreCommitteeController) DeleteCommitteeScore(scoreType string, score interface{}) error {
	switch scoreType {
	case "assignment":
		if s, ok := score.(*model.ScoreAssignmentCommittee); ok {
			return c.db.Where("id = ?", s.ID).Delete(&model.ScoreAssignmentCommittee{}).Error
		}
	case "assessment":
		if s, ok := score.(*model.ScoreAssessmentCommittee); ok {
			return c.db.Where("id = ?", s.ID).Delete(&model.ScoreAssessmentCommittee{}).Error
		}
	case "presentation":
		if s, ok := score.(*model.ScorePresentationCommittee); ok {
			return c.db.Where("id = ?", s.ID).Delete(&model.ScorePresentationCommittee{}).Error
		}
	case "report":
		if s, ok := score.(*model.ScoreReportCommittee); ok {
			return c.db.Where("id = ?", s.ID).Delete(&model.ScoreReportCommittee{}).Error
		}
	default:
		return fmt.Errorf("invalid score type: %s", scoreType)
	}
	return fmt.Errorf("failed to cast score to the correct type for scoreType: %s", scoreType)
}
