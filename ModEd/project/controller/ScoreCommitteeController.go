package controller

import (
	"ModEd/core"

	"gorm.io/gorm"
)

type ScoreCommitteeController[T core.RecordInterface] struct {
	*core.BaseController[T]
	db *gorm.DB
}

func NewScoreCommitteeController[T core.RecordInterface](db *gorm.DB) *ScoreCommitteeController[T] {
	return &ScoreCommitteeController[T]{
		db:             db,
		BaseController: core.NewBaseController[T](db),
	}
}

func (c *ScoreCommitteeController[T]) ListAllCommitteeScores() ([]T, error) {
	return c.List(nil)
}

func (c *ScoreCommitteeController[T]) ListCommitteeScoresByCondition(condition string, value interface{}) ([]T, error) {
	return c.List(map[string]interface{}{condition: value})
}
