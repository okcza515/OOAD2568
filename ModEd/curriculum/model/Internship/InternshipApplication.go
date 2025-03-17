package model

import (
	"time"

	"github.com/google/uuid"
)

type InternshipApplication struct {
	InternshipApplicationId uuid.UUID     `gorm:"primaryKey;unique"`
	Company                 string        `gorm:"not null"`
	Mentor                  string        `gorm:"not null"`
	Advisor                 string        `gorm:"not null"`
	Student                 InternStudent `gorm:"foreignKey:InternID"`
	TurninDate              time.Time     `gorm:"not null"`
}
