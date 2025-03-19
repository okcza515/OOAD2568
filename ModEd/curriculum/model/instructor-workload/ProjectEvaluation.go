package model

type ProjectEvaluation struct {
	EvaluationId   int                       `gorm:"primaryKey;autoIncrement"`
	TaskId         int                       `gorm:"not null;index"`
	GroupId        int                       `gorm:"not null;index"`
	AssignmentType ProjectEvaluationTypeEnum `gorm:"type:float6;not null"`
	Score          float64                   `gorm:"type:decimal(5,2);not null"`
	Comment        string                    `gorm:"type:text;not null"`
}
