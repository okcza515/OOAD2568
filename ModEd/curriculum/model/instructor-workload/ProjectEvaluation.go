package model

import "gorm.io/gorm"

type ProjectEvaluation struct {
	gorm.Model
	GroupId        int                       `gorm:"not null;index"`
	AssignmentType ProjectEvaluationTypeEnum `gorm:"type:enum('Assignment', 'Proposal', 'Report', 'Presentation');not null"`
	Score          float64                   `gorm:"type:decimal(5,2);not null"`
	Comment        string                    `gorm:"type:text;not null"`
	Audit
}
