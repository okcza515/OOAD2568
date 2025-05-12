// MEP-1003 Student Recruitment
package controller

type BaseReport interface {
	GetFilteredReport(condition map[string]interface{}) ([]interface{}, error)
	DisplayReport(filteredReport []interface{})
}
