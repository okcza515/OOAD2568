// MEP-1008
package model

import (
	projectModel "ModEd/project/model"
	"fmt"
)

type EvaluationStrategy interface {
	Evaluate(criteria []projectModel.AssessmentCriteria) (float64, string)
}

type PresentationEvaluationStrategy struct{}
type AssignmentEvaluationStrategy struct{}
type ReportEvaluationStrategy struct{}

type MarkedCriteria struct {
	projectModel.AssessmentCriteria
	IsPass bool
}

type WeightedCriteria struct {
	projectModel.AssessmentCriteria
	Weight float64
}

func (p *PresentationEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) (float64, string) {
	fmt.Println("Evaluate Presentation")

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

	comment := "Should be improved"

	return score, comment
}

func (a *AssignmentEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) (float64, string) {
	fmt.Println("Evaluate Assignment")

	mockedCriteria := []MarkedCriteria{
		{criteria[0], true},
		{criteria[1], false},
		{criteria[2], true},
	}

	score := 0.0
	for _, c := range mockedCriteria {
		if c.IsPass {
			score += 10
		}
	}

	comment := "Good job!"
	return score, comment
}

func (r *ReportEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) (float64, string) {

	fmt.Println("Evaluate Report")

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
		if c.IsPass {
			if val, ok := scale[c.CriteriaName]; ok {
				score += val
			} else {
				score += 10
			}
		}
	}
	return score, ""
}
