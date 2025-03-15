package util

import (
	"ModEd/common/model"
)

type JSONMapper struct {
	Path string
}

func (mapper *JSONMapper) Map() []*model.Student {
	result := []*model.Student{}
	return result
}
