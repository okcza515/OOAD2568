package controller

import (
	"ModEd/project/model"

	"gorm.io/gorm"
)

type InstanceStorer struct {
	DB                         *gorm.DB
	Advisor                    *AdvisorController
	Assessment                 *AssessmentController
	AssessmentCriteria         *AssessmentCriteriaController
	AssessmentCriteriaLink     *AssessmentCriteriaLinkController
	Assignment                 *AssignmentController
	Committee                  *CommitteeController
	GroupMember                *GroupMemberController
	Presentation               *PresentationController
	Progress                   *ProgressController
	Report                     *ReportController
	SeniorProject              *SeniorProjectController
	ScoreAssessmentAdvisor     *ScoreAdvisorController[*model.ScoreAssessmentAdvisor]
	ScoreAssessmentCommittee   *ScoreCommitteeController[*model.ScoreAssessmentCommittee]
	ScoreAssignmentAdvisor     *ScoreAdvisorController[*model.ScoreAssignmentAdvisor]
	ScoreAssignmentCommittee   *ScoreCommitteeController[*model.ScoreAssignmentCommittee]
	ScorePresentationAdvisor   *ScoreAdvisorController[*model.ScorePresentationAdvisor]
	ScorePresentationCommittee *ScoreCommitteeController[*model.ScorePresentationCommittee]
	ScoreReportAdvisor         *ScoreAdvisorController[*model.ScoreReportAdvisor]
	ScoreReportCommittee       *ScoreCommitteeController[*model.ScoreReportCommittee]
}

func CreateInstance(db *gorm.DB) *InstanceStorer {
	return &InstanceStorer{
		DB:                         db,
		Advisor:                    NewAdvisorController(db),
		Assessment:                 NewAssessmentController(db),
		AssessmentCriteria:         NewAssessmentCriteriaController(db),
		AssessmentCriteriaLink:     NewAssessmentCriteriaLinkController(db),
		Assignment:                 NewAssignmentController(db),
		Committee:                  NewCommitteeController(db),
		GroupMember:                NewGroupMemberController(db),
		Presentation:               NewPresentationController(db),
		Progress:                   NewProgressController(db),
		Report:                     NewReportController(db),
		SeniorProject:              NewSeniorProjectController(db),
		ScoreAssessmentAdvisor:     NewScoreAdvisorController[*model.ScoreAssessmentAdvisor](db),
		ScoreAssessmentCommittee:   NewScoreCommitteeController[*model.ScoreAssessmentCommittee](db),
		ScoreAssignmentAdvisor:     NewScoreAdvisorController[*model.ScoreAssignmentAdvisor](db),
		ScoreAssignmentCommittee:   NewScoreCommitteeController[*model.ScoreAssignmentCommittee](db),
		ScorePresentationAdvisor:   NewScoreAdvisorController[*model.ScorePresentationAdvisor](db),
		ScorePresentationCommittee: NewScoreCommitteeController[*model.ScorePresentationCommittee](db),
		ScoreReportAdvisor:         NewScoreAdvisorController[*model.ScoreReportAdvisor](db),
		ScoreReportCommittee:       NewScoreCommitteeController[*model.ScoreReportCommittee](db),
	}
}
