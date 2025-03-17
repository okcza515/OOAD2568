package model

type ScoreAssessmentCommittee struct {
	ScoreAssessment
	ComitteeId uuid.UUID `gorm:"type:text; not null; index`
}
