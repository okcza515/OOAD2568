// MEP-1008
package model

import (
	"ModEd/core"
	"fmt"

	"gorm.io/gorm"
)

type ProjectEvaluation struct {
	gorm.Model
	GroupId        int     `gorm:"not null;index"`
	AssignmentId   int     `gorm:"not null;index"`
	AssignmentType string  `gorm:"type:varchar(20);not null"`
	Score          float64 `gorm:"type:decimal(5,2);not null"`
	Comment        string  `gorm:"type:text;not null"`
	*core.SerializableRecord
}

func (p *ProjectEvaluation) GetID() uint {
	return p.ID
}

func (p *ProjectEvaluation) ToString() string {
	return fmt.Sprintf("%+v", p)
}

func (p *ProjectEvaluation) Validate() error {
	if p.GroupId == 0 {
		return fmt.Errorf("Group ID cannot be zero")
	}
	if p.AssignmentId == 0 {
		return fmt.Errorf("Assignment ID cannot be zero")
	}
	if p.AssignmentType == "" {
		return fmt.Errorf("Assignment Type cannot be empty")
	}
	if p.Score < 0 || p.Score > 100 {
		return fmt.Errorf("Score must be between 0 and 100")
	}
	if p.Comment == "" {
		return fmt.Errorf("Comment cannot be empty")
	}
	return nil
}
