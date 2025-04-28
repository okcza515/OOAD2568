package controller

import (
	"ModEd/core"

	"gorm.io/gorm"
)

type ScoreCommitteeController[T core.RecordInterface] struct {
	*core.BaseController[T]
	db             *gorm.DB
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

func (c *ScoreCommitteeController[T]) RetrieveCommitteeScore(id uint) (T, error) {
	return c.RetrieveByID(id)
}

func (c *ScoreCommitteeController[T]) InsertCommitteeScore(score T) error {
	return c.Insert(score)
}

func (c *ScoreCommitteeController[T]) UpdateCommitteeScore(score T) error {
	return c.UpdateByID(score)
}

func (c *ScoreCommitteeController[T]) DeleteCommitteeScore(id uint) error {
	return c.DeleteByID(id)
}
