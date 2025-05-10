package controller

import (
	"ModEd/eval/model"
)

type GradingStrategy interface {
    Grade(question model.Question, answer model.Answer) float64
}

type MultipleChoice struct{}

func (multipleChoice *MultipleChoice) Grade(question model.Question, answer model.Answer) float64 {
    if question.Correct_answer == answer.Answer {
        return question.Score
    }
    return 0
}

type ShortAnswer struct{}

func (shortAnswer *ShortAnswer) Grade(question model.Question, answer model.Answer) float64 {
    return 0
}

type TrueFalse struct{}

func (trueFalse *TrueFalse) Grade(question model.Question, answer model.Answer) float64 {
    if question.Correct_answer == answer.Answer {
        return question.Score
    }
    return 0
}

type Subjective struct{}

func (subjective *Subjective) Grade(question model.Question, answer model.Answer) float64 {
    return 0
}

type GradingStrategyFactory struct{}

func (f *GradingStrategyFactory) GetStrategy(questionType model.QuestionType) GradingStrategy {
    switch questionType {
	case "Multiple_choice":
        return &MultipleChoice{}
	case "Short_answer":
        return &ShortAnswer{}
	case "True_false":
        return &TrueFalse{}
    case "Subjective":
        return &Subjective{}
    default:
        return nil
    }
}
