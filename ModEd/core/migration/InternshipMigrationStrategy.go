package migration

import (
	"ModEd/curriculum/model"
)

type InternshipMigrationStrategy struct {
}

func (s *InternshipMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.InternStudent{},
		&model.Company{},
		&model.SupervisorReview{},
		&model.InternshipReport{},
		&model.InternshipApplication{},
		&model.InternshipMentor{},
		&model.InternshipCriteria{},
		&model.InternshipInformation{},
		&model.InternshipResultEvaluation{},
	}
}
