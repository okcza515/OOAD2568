package model

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
