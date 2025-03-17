package model

type ScoreAssessmentAdvisor struct {
	ScoreAssessment
	Advisor Advisor `gorm:"Advisor"`
}

type ScoreAdvisorRepository interface {
	Create(score *ScoreAssessmentAdvisor) error
	Update(scoreAssessment *ScoreAssessmentAdvisor) error
	Delete(scoreAssessment *ScoreAssessmentAdvisor) error
	GetByAssessmentId(assesmentId string) (*ScoreAssessmentAdvisor, error)
}