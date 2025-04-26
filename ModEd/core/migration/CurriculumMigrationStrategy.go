package migration

import (
	"ModEd/curriculum/model"
)

type CurriculumMigrationStrategy struct {
}

func (s *CurriculumMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.Curriculum{},
		&model.Class{},
		&model.Course{},
	}
}
