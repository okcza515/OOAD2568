package util

import (
	"ModEd/common/model"
	"os"

	"github.com/gocarina/gocsv"
)

type CSVMapper struct {
	Path string
}

func (mapper *CSVMapper) Map() []*model.Student {
	studentFile, err := os.OpenFile(mapper.Path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer studentFile.Close()

	result := []*model.Student{}
	if err := gocsv.UnmarshalFile(studentFile, &result); err != nil {
		panic(err)
	}
	return result
}
