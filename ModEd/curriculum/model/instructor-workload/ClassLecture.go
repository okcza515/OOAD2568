package model

import (
	common "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ClassLecture struct {
	gorm.Model
	ClassId          uint                  `gorm:"not null" json:"class_id"`
	Class            curriculumModel.Class `gorm:"foreignKey:ClassID;references:ClassId" json:"-"`
	LectureName      string                `gorm:"not null" json:"lecture_name"`
	InstructorId     uint                  `gorm:"not null" json:"instructor_id"`
	Instructor       common.Instructor     `gorm:"foreignKey:InstructorId;references:ID" json:"-"`
	ClassMaterialIds []uint                `gorm:"-" json:"class_material_id,omitempty"`
	ClassMaterials   []ClassMaterial       `gorm:"foreignKey:ClassLectureId" json:"material,omitempty"`

	StartTime string `gorm:"not null" json:"start_time"`
	EndTime   string `gorm:"not null" json:"end_time"`
}
