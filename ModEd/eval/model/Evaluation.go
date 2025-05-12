// 65070503445
// MEP-1006

package model

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
	"time"
)

type Evaluation struct {
	core.BaseModel
	StudentCode    string                 `gorm:"column:student_code" csv:"student_code"`
	Student        commonModel.Student    `gorm:"foreignKey:StudentCode;references:StudentCode" csv:"-"`
	InstructorCode string                 `gorm:"column:instructor_code" csv:"instructor_code"`
	Instructor     commonModel.Instructor `gorm:"foreignKey:InstructorCode;references:InstructorCode" csv:"-"`
	AssignmentId   uint                   `gorm:"column:assignment_id" csv:"assignment_id"`
	Assignment     Assignment             `gorm:"foreignKey:AssignmentId;references:AssignmentId" csv:"-"`
	Score          uint                   `gorm:"column:score;default:0" csv:"score"`
	Comment        string                 `gorm:"column:comment;type:varchar(255);" csv:"comment"`
	EvaluatedAt    time.Time              `gorm:"column:evaluated_at" csv:"evaluated_at"`
}
