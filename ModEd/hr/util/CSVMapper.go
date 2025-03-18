package util

import (
    "os"

    "github.com/gocarina/gocsv"
)

type CSVMapper[T any] struct {
    Path string
}

func (mapper *CSVMapper[T]) Map() []*T {
    csvFile, err := os.OpenFile(mapper.Path, os.O_RDWR|os.O_CREATE, os.ModePerm)
    if err != nil {
        panic(err)
    }
    defer csvFile.Close()

    var result []*T
    if err := gocsv.UnmarshalFile(csvFile, &result); err != nil {
        panic(err)
    }
    return result
}