//MEP-1006

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
		&model.Assignment{},
		&model.AssignmentSubmission{},
	}
}
