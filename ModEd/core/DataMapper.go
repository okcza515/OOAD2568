package core

import (
	"errors"
)

type DataMapper[T any] interface {
	Deserialize() []*T
	Serialize(data []*T) error
}

func CreateMapper[T any](path string) (DataMapper[T], error) {
	length := len(path)
	if path[length-4:length] == ".csv" {
		mapper := &CSVMapper[T]{Path: path}
		return mapper, nil
	} else if path[length-5:length] == ".json" {
		mapper := &JSONMapper[T]{Path: path}
		return mapper, nil
	} else {
		return nil, errors.New("not supported file extension")
	}
}
