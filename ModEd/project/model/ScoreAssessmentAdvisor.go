package model

type ScoreAssessmentAdvisor struct {
	ScoreAssessment
	AdvisorId uuid.UUID `gorm:"type :text; not null; index"`
}
