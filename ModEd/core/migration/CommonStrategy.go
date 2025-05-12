package migration

import (
	"ModEd/common/model"
)

type CommonStrategy struct {
}

func (s *CommonStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.Department{},
		&model.Faculty{},
		&model.Instructor{},
		&model.InternationalStudent{},
		&model.RegularStudent{},
		&model.Student{},
	}
}
