package core

import "github.com/gocarina/gocsv"

type SerializableRecord struct {
}

func (record *SerializableRecord) ToCSVRow() string {
	records := [1]*SerializableRecord{record}
	serialized, _ := gocsv.MarshalString(&records)
	return serialized
}

func (record *SerializableRecord) FromCSV(raw string) error {
	records := [1]*SerializableRecord{record}
	err := gocsv.UnmarshalString(raw, records)
	return err
}

func (record *SerializableRecord) ToJSON() string {
	return ""
}

func (record *SerializableRecord) FromJSON(raw string) error {
	return nil
}
