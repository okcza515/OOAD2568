//MEP-1008
package model

import (
	common "ModEd/common/model"

	gorm "gorm.io/gorm"
)

type StudentRequest struct {
	gorm.Model
	StudentCode  string            `gorm:"not null" json:"student_code"`
	Student      common.Student    `gorm:"foreignKey:StudentCode;references:StudentCode" json:"student"`
	InstructorId uint              `gorm:"not null" json:"instructor_id"`
	Instructor   common.Instructor `gorm:"foreignKey:InstructorId;references:ID" json:"-"`
	RequestType  string            `gorm:"not null" json:"request_type"`
	CreatedAt    string            `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    string            `gorm:"autoUpdateTime" json:"updated_at"`
	Review       string            `gorm:"default:null" json:"review"`
	Comment      string            `gorm:"default:null" json:"comment"`
}
