package model

import (
	"time"

	"gorm.io/gorm"
)

type Presentation struct {
	gorm.Model
	SeniorProjectID             uint                         `gorm:"not null;index"`
	PresentationType            PresentationType             `gorm:"type:varchar(50);not null"`
	Date                        time.Time                    `gorm:"not null"`
	ScoresPresentationCommittee []ScorePresentationCommittee `gorm:"foreignKey:PresentationID"`
	ScoresPresentationAdvisor   []ScorePresentationAdvisor   `gorm:"foreignKey:PresentationID"`
}
