package migration

import (
	"ModEd/hr/model"
)

type HRMigrationStrategy struct {
}

func (s *HRMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.InstructorInfo{},
		&model.StudentInfo{},
		&model.RequestLeaveInstructor{},
		&model.RequestLeaveStudent{},
		&model.RequestResignationInstructor{},
		&model.RequestResignationStudent{},
		&model.RequestRaiseInstructor{},
	}
}
