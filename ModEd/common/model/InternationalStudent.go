package model

import (
	"ModEd/core"
)

type InternationalStudent struct {
	core.BaseModel
	Student
}

func (is InternationalStudent) Validate() error {
	return is.Student.Validate()
}
