package controller

type BaseReport interface {
	GetReport(model interface{}) ([]interface{}, error)
	FilterReport(report []interface{}) ([]interface{}, error)
	DisplayReport(filteredReport []interface{})
}
