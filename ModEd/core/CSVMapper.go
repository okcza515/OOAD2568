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

func (mapper *CSVMapper[T]) Serialize(data []*T) error {
	file, err := os.OpenFile(mapper.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(&data, file)
	if err != nil {
		return err
	}
	return nil
}
