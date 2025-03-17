package model

import (
	"errors"

	"github.com/google/uuid"
)

type ProjectEvaluationStrategy interface {
	Evaluate(evaluation Evaluation) (float64, string, error)
}

type Evaluation struct {
	ID             uuid.UUID
	TaskID         uuid.UUID
	GroupID        uuid.UUID
	AssignmentType string
	Score          float64
	Comment        string

	strategy ProjectEvaluationStrategy
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
