package model

type ApplicationStatus string

const (
	Pending        ApplicationStatus = "Pending"
	UnderReview    ApplicationStatus = "Under Review"
	InterviewStage ApplicationStatus = "Interview"
	Accepted       ApplicationStatus = "Accepted"
	Rejected       ApplicationStatus = "Rejected"
)
