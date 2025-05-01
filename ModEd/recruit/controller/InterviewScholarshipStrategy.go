package controller

import "fmt"

type ScholarshipEvaluationStrategy struct{}

func (s ScholarshipEvaluationStrategy) GetCriteria() []string {
	return []string{
		"Academic Performance",
		"Financial Need",
		"Motivation",
		"Community Involvement",
	}
}

func (s ScholarshipEvaluationStrategy) Evaluate(scores map[string]float64) (float64, error) {
	var total float64
	for _, criterion := range s.GetCriteria() {
		score, ok := scores[criterion]
		if !ok {
			return 0, fmt.Errorf("missing score for %s", criterion)
		}
		if score < 0 || score > 10 {
			return 0, fmt.Errorf("score for %s must be between 0 and 10", criterion)
		}
		total += score
	}
	return total, nil
}

func (s ScholarshipEvaluationStrategy) ProjectType() string {
	return "Scholarship"
}
