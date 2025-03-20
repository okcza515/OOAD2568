// MEP-1003 Student Recruitment
package model

import (
	"github.com/google/uuid"
)

// ApplicationRound defines the details of an application round.
type ApplicationRound struct {
	RoundID   uuid.UUID `gorm:"primaryKey"`
	RoundName string
}
