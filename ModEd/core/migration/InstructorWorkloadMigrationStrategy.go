package migration

import (
	"ModEd/curriculum/model"
)

type InstructorWorkloadMigrationStrategy struct {
}

func (s *InstructorWorkloadMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.ClassLecture{},
		&model.StudentAdvisor{},
		&model.StudentRequest{},
		&model.Meeting{},
		&model.OnlineMeeting{},
		&model.ExternalMeeting{},
		&model.ProjectEvaluation{},
	}
}
