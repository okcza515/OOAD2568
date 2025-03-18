package util

import (
	"encoding/json"
	"os"
)

type JSONMapper[T any] struct {
	Path string
}

func (mapper *JSONMapper[T]) Map() []*T {
	file, err := os.ReadFile(mapper.Path)
	if err != nil {
		panic(err)
	}

	var result []*T
	if err := json.Unmarshal(file, &result); err != nil {
		panic(err)
	}
	return result
}
