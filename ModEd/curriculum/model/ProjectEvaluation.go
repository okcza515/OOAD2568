//MEP-1008
package model

import "gorm.io/gorm"

type ProjectEvaluation struct {
	gorm.Model
	GroupId        int     `gorm:"not null;index"`
	AssignmentId   int     `gorm:"not null;index"`
	AssignmentType string  `gorm:"type:varchar(20);not null"`
	Score          float64 `gorm:"type:decimal(5,2);not null"`
	Comment        string  `gorm:"type:text;not null"`
}
