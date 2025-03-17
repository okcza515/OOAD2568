package util

import (
	"errors"
)

type DataMapper interface {
	Map(data interface{}) error
}

func CreateMapper(path string) (DataMapper, error) {
	length := len(path)
	if path[length-4:length] == ".csv" {
		mapper := &CSVMapper{Path: path}
		return mapper, nil
	} else if path[length-5:length] == ".json" {
		mapper := &JSONMapper{Path: path}
		return mapper, nil
	} else {
		return nil, errors.New("not supported file extension")
	}
}
