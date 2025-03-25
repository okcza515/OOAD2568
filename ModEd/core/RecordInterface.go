package core

type RecordInterface interface {
	GetID() uint
	ToString() string

	Validate() error

	ToCSVRow() string
	FromCSV(raw string) error

	ToJSON() string
	FromJSON(raw string) error
}
