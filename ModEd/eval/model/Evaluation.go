// 65070503445
// MEP-1006

package model

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
	"fmt"
	"time"
)

type Evaluation struct {
	core.BaseModel
	StudentCode    string
	Student        commonModel.Student `gorm:"foreignKey:StudentCode;references:StudentCode"`
	InstructorCode string
	Instructor     commonModel.Instructor `gorm:"foreignKey:InstructorCode;references:InstructorCode"`
	AssessmentId   uint                   //assignment
	Assessment     Assessment             `gorm:"foreignKey:AssessmentId;references:AssessmentId"`
	Score          uint                   `gorm:"column:Score;default:0"`
	Comment        string                 `gorm:"type:varchar(255);"`
	EvaluatedAt    time.Time
}

func (e Evaluation) GetID() uint {
	return e.ID
}

func (e Evaluation) ToString() string {
	return fmt.Sprintf("Evaluation{ID: %d, StudentCode: %s, InstructorCode: %s, AssessmentId: %d, Score: %d, Comment: %s, EvaluatedAt: %v}",
		e.ID, e.StudentCode, e.InstructorCode, e.AssessmentId, e.Score, e.Comment, e.EvaluatedAt)
}

func (e Evaluation) Validate() error {
	if e.StudentCode == "" {
		return fmt.Errorf("StudentCode is required")
	}
	if e.InstructorCode == "" {
		return fmt.Errorf("InstructorCode is required")
	}
	if e.AssessmentId == 0 {
		return fmt.Errorf("AssessmentId is required")
	}
	return nil
}
