package model

type StudentRequestStatusEnum string

const (
	PENDING  StudentRequestStatusEnum = "Pending"
	APPROVED StudentRequestStatusEnum = "Approved"
	REJECTED StudentRequestStatusEnum = "Rejected"
)

func IsValidRequestStatus(requestType string) bool {
	switch StudentRequestStatusEnum(requestType) {
	case PENDING, APPROVED, REJECTED:
		return true
	default:
		return false
	}
}
