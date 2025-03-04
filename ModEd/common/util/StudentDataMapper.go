package util

import (
	"ModEd/common/model"
	"errors"
)

type StudentDataMapper interface {
	Map() []*model.Student
}

func CreateMapper(path string) (StudentDataMapper, error) {
	length := len(path)
	if path[length-4:length] == ".csv" {
		mapper := &StudentCSVMapper{Path: path}
		return mapper, nil
	} else if path[length-4:length] == ".json" {
		mapper := &StudentJSONMapper{Path: path}
		return mapper, nil
	} else {
		return nil, errors.New("not supported file extension")
	}
}
