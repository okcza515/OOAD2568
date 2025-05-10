package migration

import (
	"ModEd/eval/model"
)

type QuizStrategy struct {
}

func (s *QuizStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.Quiz{},
		&model.Progress{},
		&model.Evaluation{},
		&model.Assessment{},
		&model.AssessmentSubmission{},
	}
}
