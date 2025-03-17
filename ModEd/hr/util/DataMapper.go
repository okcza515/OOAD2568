package util

import (
	"errors"
	"strings"
)

type DataMapper[T any] interface {
	Map() []*T
}

func CreateMapper[T any](path string) (DataMapper[T], error) {
	if strings.HasSuffix(path, ".csv") {
		return &CSVMapper[T]{Path: path}, nil
	} else if strings.HasSuffix(path, ".json") {
		return &JSONMapper[T]{Path: path}, nil
	}
	return nil, errors.New("not supported file extension")
}
