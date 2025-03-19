package model

import (
	"github.com/google/uuid"
)

type ProjectEvaluation struct {
	EvaluationId   uuid.UUID                 `gorm:"type:string;default:uuid_generate_v4();primaryKey"`
	TaskId         uuid.UUID                 `gorm:"type:string;not null;index"`
	GroupId        uuid.UUID                 `gorm:"type:string;not null;index"`
	AssignmentType ProjectEvaluationTypeEnum `gorm:"type:float6;not null"`
	Score          float64                   `gorm:"type:decimal(5,2);not null"`
	Comment        string                    `gorm:"type:text;not null"`
}
