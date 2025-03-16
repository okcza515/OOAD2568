package model

import (
  "ModEd/common/model"
  "gorm.io/gorm"
)

type InternStudent struct {
  gorm.Model
  model.Student
  InternStatus InternStatus  `csv:"-"`
}