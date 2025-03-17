package model

import (
	"errors"

	"github.com/google/uuid"
)

type ProjectEvaluationStrategy interface {
	Evaluate(evaluation Evaluation) (float64, string, error)
}

type Evaluation struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TaskID         uuid.UUID `gorm:"type:uuid;not null"`
	GroupID        uuid.UUID `gorm:"type:uuid;not null"`
	AssignmentType string    `gorm:"not null"`
	Score          float64   `gorm:"not null"`
	Comment        string    `gorm:"-"`

	strategy ProjectEvaluationStrategy `gorm:"-"`
}

func (e *Evaluation) SetEvaluationStrategy(strategy ProjectEvaluationStrategy) {
	e.strategy = strategy
}

func (e *Evaluation) ExecuteEvaluation() (float64, string, error) {
	if e.strategy == nil {
		return 0, "", errors.New("Evaluation strategy is not set")
	}
	return e.strategy.Evaluate(*e)
}
