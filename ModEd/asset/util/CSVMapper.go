package util

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type CSVMapper struct {
	Path string
}

func (mapper *CSVMapper) Map(data interface{}) error {
	pwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	fmt.Println(pwd)

	filePath, err := os.OpenFile(mapper.Path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	fmt.Println("skibidi: ", filePath)

	if err != nil {
		panic(err)
	}
	defer filePath.Close()

	//result := new([]interface{})
	if err := gocsv.UnmarshalFile(filePath, &data); err != nil {
		panic(err)
	}
	return nil
}
