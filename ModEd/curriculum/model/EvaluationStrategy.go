// MEP-1008
package model

import (
	projectModel "ModEd/project/model"
	"fmt"
)

type EvaluationStrategy interface {
	Evaluate(criteria []projectModel.AssessmentCriteria) float64
}

type PresentationEvaluationStrategy struct{}
type AssignmentEvaluationStrategy struct{}
type ReportEvaluationStrategy struct{}

type MarkedCriteria struct {
	projectModel.AssessmentCriteria
	IsCompleted bool
}

type WeightedCriteria struct {
	projectModel.AssessmentCriteria
	Weight float64
}

func (p *PresentationEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	fmt.Println("Evaluating presentation criteria...")

	score := 0.0
	scale := map[string]float64{
		"poor":    5,
		"average": 10,
		"good":    15,
	}

	for _, c := range criteria {
		if val, ok := scale[c.CriteriaName]; ok {
			score += val
		}
	}

	return score
}

func (a *AssignmentEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	fmt.Println("Evaluating assignment criteria...")

	mockedCriteria := []MarkedCriteria{
		{criteria[0], true},
		{criteria[1], false},
		{criteria[2], true},
	}

	score := 0.0
	for _, c := range mockedCriteria {
		if c.IsCompleted {
			score += 10
		}
	}

	return score
}

func (r *ReportEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {

	fmt.Println("Evaluating report criteria...")

	score := 0.0
	scale := map[string]float64{
		"poor":    5,
		"average": 10,
		"good":    15,
	}

	mockedCriteria := []MarkedCriteria{
		{criteria[0], true},
		{criteria[1], false},
		{criteria[2], true},
	}

	for _, c := range mockedCriteria {
		if c.IsCompleted {
			if val, ok := scale[c.CriteriaName]; ok {
				score += val
			} else {
				score += 10
			}
		}
	}
	return score
}
