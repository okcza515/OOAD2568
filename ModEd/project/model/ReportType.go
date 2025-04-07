package model

import (
	"database/sql/driver"
	"fmt"
)

type ReportType string

const (
	ReportTypeIdea     ReportType = "Idea"
	ReportTypeProposal ReportType = "Proposal"
	ReportTypeProgress ReportType = "Progress"
	ReportTypeMidterm  ReportType = "Midterm"
	ReportTypeFinal    ReportType = "Final"
)

func ValidReportTypes() []ReportType {
	return []ReportType{
		ReportTypeIdea,
		ReportTypeProposal,
		ReportTypeProgress,
		ReportTypeMidterm,
		ReportTypeFinal,
	}
}
func (rt ReportType) IsValid() bool {
	for _, validType := range ValidReportTypes() {
		if rt == validType {
			return true
		}
	}
	return false
}

func (rt ReportType) Value() (driver.Value, error) {
	return string(rt), nil
}

func (rt *ReportType) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan ReportType: %v", value)
	}
	*rt = ReportType(str)
	return nil
}
