// MEP-1007
package controller

import "ModEd/eval/model"

type GradingContext struct {
	strategyMap map[model.QuestionType]IGradingStrategy
}

func NewGradingContext() *GradingContext {
	return &GradingContext{
		strategyMap: make(map[model.QuestionType]IGradingStrategy),
	}
}
func (g *GradingContext) Grade(questionType model.QuestionType, submissionID uint) (float64, error) {
	return g.strategyMap[questionType].Grade(submissionID)
}

func (g *GradingContext) AddGradingStrategy(questionType model.QuestionType, strategy IGradingStrategy) {
	g.strategyMap[questionType] = strategy
}

func (g *GradingContext) GradeAll(submissionID uint) (float64, error) {
	var totalScore float64
	for _, v := range g.strategyMap {
		score, err := v.Grade(submissionID)
		if err != nil {
			return 0, err
		}
		totalScore += score
	}

	return totalScore, nil
}
