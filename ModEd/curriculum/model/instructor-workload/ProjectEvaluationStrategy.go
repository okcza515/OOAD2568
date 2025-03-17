package model

import (
	"errors"

	"github.com/google/uuid"
)

type ProjectEvaluationStrategy interface {
	Evaluate(evaluation ProjectEvaluation) (float64, string, error)
}

type ProjectEvaluation struct {
	ID             uuid.UUID
	TaskID         uuid.UUID
	GroupID        uuid.UUID
	AssignmentType string
	Score          float64
	Comment        string

	strategy ProjectEvaluationStrategy
}

func (e *ProjectEvaluation) SetEvaluationStrategy(strategy ProjectEvaluationStrategy) {
	e.strategy = strategy
}

func (e *ProjectEvaluation) ExecuteEvaluation() (float64, string, error) {
	if e.strategy == nil {
		return 0, "", errors.New("Evaluation strategy is not set")
	}
	return e.strategy.Evaluate(*e)
}
