package model

import (
	"ModEd/core"
)

type RegularStudent struct {
	core.BaseModel
	Student
}

func (rs RegularStudent) Validate() error {
	return rs.Student.Validate()
}
