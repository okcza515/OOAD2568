// MEP-1003 Student Recruitment
package model

type ApplicationStatus string

const (
	Pending        ApplicationStatus = "Pending"
	InterviewStage ApplicationStatus = "Interview"
	Accepted       ApplicationStatus = "Accepted"
	Rejected       ApplicationStatus = "Rejected"
)
