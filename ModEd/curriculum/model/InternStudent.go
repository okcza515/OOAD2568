package model

import (
	"ModEd/common/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InternStudent struct {
  gorm.Model
  model.Student
	InternID uuid.UUID
  InternStatus InternStatus  `csv:"-"`
}