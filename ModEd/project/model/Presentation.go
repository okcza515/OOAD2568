package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Presentation struct {
	gorm.Model
	PresentationId              uuid.UUID                    `gorm:"type:uuid;primaryKey"`
	SeniorProjectId             uuid.UUID                    `gorm:"not null;index"`
	PresentationType            PresentationType             `gorm:"type:varchar(50);not null"`
	Date                        time.Time                    `gorm:"type:date;not null"`
	ScoresPresentationCommittee []ScorePresentationCommittee `gorm:"foreignKey:PresentationId"`
	ScoresPresentationAdvisor   []ScorePresentationAdvisor   `gorm:"foreignKey:PresentationId"`
}
