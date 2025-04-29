package controller

import (
	"ModEd/core"

	"gorm.io/gorm"
)

type ScoreAdvisorController[T core.RecordInterface] struct {
	*core.BaseController[T]
	db *gorm.DB
}

func NewScoreAdvisorController[T core.RecordInterface](db *gorm.DB) *ScoreAdvisorController[T] {
	return &ScoreAdvisorController[T]{
		db:             db,
		BaseController: core.NewBaseController[T](db),
	}
}

func (c *ScoreAdvisorController[T]) ListAllAdvisorScores() ([]T, error) {
	return c.List(nil)
}

func (c *ScoreAdvisorController[T]) ListAdvisorScoresByCondition(condition string, value interface{}) ([]T, error) {
	return c.List(map[string]interface{}{condition: value})
}

func (c *ScoreAdvisorController[T]) RetrieveAdvisorScore(id uint) (T, error) {
	return c.RetrieveByID(id)
}

func (c *ScoreAdvisorController[T]) InsertAdvisorScore(score T) error {
	return c.Insert(score)
}

func (c *ScoreAdvisorController[T]) UpdateAdvisorScore(score T) error {
	return c.UpdateByID(score)
}

func (c *ScoreAdvisorController[T]) DeleteAdvisorScore(id uint) error {
	return c.DeleteByID(id)
}
