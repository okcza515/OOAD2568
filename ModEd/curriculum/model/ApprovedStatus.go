// MEP-1009 Student Internship
package model

type ApprovedStatus string

const (
	INTERN_APP_NOT_START   ApprovedStatus = "Not Started"
	INTERN_APP_IN_PROGRESS ApprovedStatus = "In Progress"
	INTERN_APP_APPROVED    ApprovedStatus = "approved"
	INTERN_APP_REJECT      ApprovedStatus = "rejected"
)
