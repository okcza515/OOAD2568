package model

type StudentRequestStatusEnum string

const (
	REQUEST_PENDING  StudentRequestStatusEnum = "Pending"
	REQUEST_APPROVED StudentRequestStatusEnum = "Approved"
	REQUEST_REJECTED StudentRequestStatusEnum = "Rejected"
)
