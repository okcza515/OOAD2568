package model

type StudentRequestTypeEnum string

const (
	SickLeave     StudentRequestTypeEnum = "SickLeave"
	PersonalLeave StudentRequestTypeEnum = "PersonalLeave"
	AddCourse     StudentRequestTypeEnum = "ExtraCourseEnrollment"
)

func IsValidRequestType(requestType string) bool {
	switch StudentRequestTypeEnum(requestType) {
	case SickLeave, PersonalLeave, AddCourse:
		return true
	default:
		return false
	}
}
