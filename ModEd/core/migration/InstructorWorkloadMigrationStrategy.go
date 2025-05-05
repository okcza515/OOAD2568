package migration

import (
	"ModEd/curriculum/model"
)

type InstructorWorkloadMigrationStrategy struct {
}

func (s *InstructorWorkloadMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.StudentAdvisor{},
		&model.StudentRequest{},
		&model.Meeting{},
		&model.MeetingAttendee{},
		&model.OnlineMeeting{},
		&model.ExternalMeeting{},
		&model.ProjectEvaluation{},
		&model.CoursePlan{},
		&model.ClassMaterial{},
	}
}
