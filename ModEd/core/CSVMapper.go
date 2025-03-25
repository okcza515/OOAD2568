package core

import (
	"os"

	"github.com/gocarina/gocsv"
)

type CSVMapper[T any] struct {
	Path string
}

func (mapper *CSVMapper[T]) Deserialize() []*T {
	studentFile, err := os.OpenFile(mapper.Path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer studentFile.Close()

	result := []*T{}
	if err := gocsv.UnmarshalFile(studentFile, &result); err != nil {
		panic(err)
	}
	return result
}
