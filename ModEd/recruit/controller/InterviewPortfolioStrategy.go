package controller

import "fmt"

type PortfolioEvaluationStrategy struct{}

func (s PortfolioEvaluationStrategy) GetCriteria() []string {
	return []string{"Communication", "Creativity", "ProjectKnowledge", "Analytical Thinking"}
}

func (s PortfolioEvaluationStrategy) Evaluate(scores map[string]float64) (float64, error) {
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

func (s PortfolioEvaluationStrategy) ProjectType() string {
	return "portfolio"
}
