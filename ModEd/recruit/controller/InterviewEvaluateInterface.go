// MEP-1003 Student Recruitment
package controller

import "fmt"

type InterviewEvaluate interface {
	GetCriteria() []string
	Evaluate(scores map[string]float64) (float64, error)
	ProjectType() string
}

var evaluationStrategies = map[string]InterviewEvaluate{
	"Portfolio":   PortfolioEvaluationStrategy{},
	"Scholarship": ScholarshipEvaluationStrategy{},
	"Quota":       QuotaEvaluationStrategy{},
	"Admission":   AdmissionEvaluationStrategy{},
}

func GetStrategyByRoundName(roundName string) (InterviewEvaluate, error) {
	strategy, exists := evaluationStrategies[roundName]
	if !exists {
		return nil, fmt.Errorf("unsupported evaluation strategy for round: %s", roundName)
	}
	return strategy, nil
}
