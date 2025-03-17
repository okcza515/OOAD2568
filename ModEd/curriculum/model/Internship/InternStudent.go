package model

import (
	"ModEd/common/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InternStudent struct {
  gorm.Model
  Student 			model.Student	`json:"Student" gorm:"foreignKey:StudentID;references:SID"`
	InternID 			uuid.UUID
  InternStatus 	InternStatus  `csv:"-"`
}