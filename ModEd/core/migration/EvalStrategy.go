package migration

import "ModEd/eval/model"

// MEP-1007 Evaluation

type EvalStrategy struct {}

func (s *EvalStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.Exam{},
		&model.ExamSection{},
		&model.Question{},
		&model.MultipleChoiceAnswer{},
		&model.ShortAnswer{},
		&model.TrueFalseAnswer{},
		&model.Submission{},
		&model.MultipleChoiceAnswerSubmission{},
		&model.ShortAnswerSubmission{},
		&model.TrueFalseAnswerSubmission{},
	}
}