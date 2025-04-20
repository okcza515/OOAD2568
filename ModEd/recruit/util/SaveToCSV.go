// MEP-1003 Student Recruitment
package util

import (
	"ModEd/asset/util"
	"os"

	"github.com/gocarina/gocsv"
)

type CustomCSVMapper struct {
	util.CSVMapper
}

func (mapper *CustomCSVMapper) Save(data interface{}) error {
	file, err := os.OpenFile(mapper.Path, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gocsv.MarshalFile(data, file); err != nil {
		return err
	}
	return nil
}
