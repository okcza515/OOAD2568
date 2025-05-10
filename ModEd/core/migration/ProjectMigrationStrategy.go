package migration

//

import (
	"ModEd/project/model"
)

type ProjectMigrationStrategy struct {
}

func (s *ProjectMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.SeniorProject{},
		&model.GroupMember{},
		&model.Advisor{},
		&model.Assessment{},
		&model.AssessmentCriteria{},
		&model.Assignment{},
		&model.Committee{},
		&model.Presentation{},
		&model.Progress{},
		&model.Report{},
		&model.ScoreAssessmentAdvisor{},
		&model.ScoreAssessmentCommittee{},
		&model.ScoreAssignmentAdvisor{},
		&model.ScoreAssignmentCommittee{},
		&model.ScorePresentationAdvisor{},
		&model.ScorePresentationCommittee{},
		&model.ScoreReportAdvisor{},
		&model.ScoreReportCommittee{},
	}
}