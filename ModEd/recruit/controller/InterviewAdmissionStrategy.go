package controller

import "fmt"

type AdmissionEvaluationStrategy struct{}

func (s AdmissionEvaluationStrategy) GetCriteria() []string {
	return []string{"Communication", "Motivation", "Personality Attitude", "Analytical Thinking"}
}

func (s AdmissionEvaluationStrategy) Evaluate(scores map[string]float64) (float64, error) {
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

func (s AdmissionEvaluationStrategy) ProjectType() string {
	return "Admission"
}
