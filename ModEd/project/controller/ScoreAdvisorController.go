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
