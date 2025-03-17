package model

type ScoreAssessmentCommittee struct {
	ScoreAssessment
	Comittee Comittee `gorm:"Comittee"`
}

type ScoreCommitteeRepository interface {
	Create(score *ScoreAssessmentCommittee) error
	Update(scoreAssessment *ScoreAssessmentCommittee) error
	Delete(scoreAssessment *ScoreAssessmentCommittee) error
	GetByAssessmentId(assesmentId string) (*ScoreAssessmentCommittee, error)
}