package model

type ReportType string

const (
	ReportTypeIdea      ReportType = "idea"
	ReportTypeProposal  ReportType = "proposal"
	ReportTypeProgress  ReportType = "progress"
	ReportTypeMidterm   ReportType = "midterm"
	ReportTypeFinal     ReportType = "final"
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