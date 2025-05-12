// MEP-1007
package controller

type GradingContext struct {
	Strategy GradingStrategy
}

func (g *GradingContext) SetStrategy(strategy GradingStrategy) {
	g.Strategy = strategy
}

func (g *GradingContext) Grade(submissionID uint) (float64, error) {
	return g.Strategy.Grade(submissionID)
}
