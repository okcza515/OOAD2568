package model

import (
	"errors"
)

type ProjectEvaluationStrategy interface {
	Evaluate(evaluation ProjectEvaluation) (float64, string, error)
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
