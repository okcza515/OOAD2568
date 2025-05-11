// MEP-1003 Student Recruitment
package model

type ApplicationStatus string

const (
	Pending        ApplicationStatus = "Pending"
	InterviewStage ApplicationStatus = "Interview"
	Evaluated      ApplicationStatus = "Evaluated"
	Accepted       ApplicationStatus = "Accepted"
	Rejected       ApplicationStatus = "Rejected"
	Confirmed      ApplicationStatus = "Confirmed"
	Withdrawn      ApplicationStatus = "Withdrawn"
	Student        ApplicationStatus = "Student"
)
