package model

type ReportType int

const (
	Idea ReportType = iota
	Proposal
	Progress
	Midterm
	Final
)

var ReportTypeLabel = map[ReportType]string{
	Idea:     "Idea",
	Proposal: "Proposal",
	Progress: "Progress",
	Midterm:  "Midterm",
	Final:    "Final",
}

func (rt ReportType) String() string {
	return ReportTypeLabel[rt]
}

func ValidReportTypes() []ReportType {
	return []ReportType{Idea, Proposal, Progress, Midterm, Final}
}

func (rt ReportType) IsValid() bool {
	_, exists := ReportTypeLabel[rt]
	return exists
}
