package controller

type BaseReport interface {
	GetFilteredInterviews(condition map[string]interface{}) ([]interface{}, error)
	DisplayReport(filteredReport []interface{})
}
