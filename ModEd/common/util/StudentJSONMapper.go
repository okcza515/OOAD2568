package util

import (
	"ModEd/common/model"
)

type StudentJSONMapper struct {
	Path string
}

func (mapper *StudentJSONMapper) Map() []*model.Student {
	result := []*model.Student{}
	return result
}
