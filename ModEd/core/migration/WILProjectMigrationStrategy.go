package migration

import (
	"ModEd/curriculum/model"
)

type WILProjectMigrationStrategy struct {
}

func (s *WILProjectMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.WILProjectCourse{},
		&model.WILProjectClass{},
		&model.WILProjectMember{},
		&model.WILProjectApplication{},
		&model.WILProject{},
		&model.IndependentStudy{},
	}
}
