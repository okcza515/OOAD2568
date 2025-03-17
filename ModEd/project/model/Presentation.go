package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Presentation struct {
	gorm.Model
	PresentationId              uuid.UUID                    `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	SeniorProjectId             uuid.UUID                    `gorm:"type:text;not null;index"`
	PresentationType            PresentationType             `gorm:"type:varchar(50);not null"`
	Date                        time.Time                    `gorm:"type:date;not null"`
	ScoresPresentationCommittee []ScorePresentationCommittee `gorm:"foreignKey:PresentationId"`
	ScoresPresentationAdvisor   []ScorePresentationAdvisor   `gorm:"foreignKey:PresentationId"`
}
