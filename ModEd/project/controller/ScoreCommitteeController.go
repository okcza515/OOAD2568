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
