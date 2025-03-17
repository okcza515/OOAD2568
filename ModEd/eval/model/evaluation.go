//Chanawat Limpanatewin 65070503445

package model

import (
	"ModEd/common/model"

	"time"

	"gorm.io/gorm"
)

type Evaluation struct {
	gorm.Model
	SID          model.Student    `gorm:"foreignKey:SID"`
	InstructorId model.Instructor `gorm:"foreignKey:InstructorId"`
	Score        float64          `gorm:"not null"`
	Feedback     string
	EvaluatedAt  time.Time
	LastUpdate   time.Time
}
