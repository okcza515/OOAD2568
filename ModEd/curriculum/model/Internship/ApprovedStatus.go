package model

type ApprovedStatus string

const (
	WAIT  ApprovedStatus = "WAIT"
	APPROVED ApprovedStatus = "APPROVED"
	REJECT   ApprovedStatus = "REJECT"
)
