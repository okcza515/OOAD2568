package core

import (
	"encoding/json"
	"os"
)

type JSONMapper[T any] struct {
	Path string
}

func (mapper *JSONMapper[T]) Deserialize() []*T {
	result := []*T{}
	return result
}

func (mapper *JSONMapper[T]) Serialize(data []*T) error {
	file, err := os.OpenFile(mapper.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}
