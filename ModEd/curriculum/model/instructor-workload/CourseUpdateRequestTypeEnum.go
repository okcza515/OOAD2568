package model

type CourseUpdateRequestTypeEnum string

const (
	NAME         CourseUpdateRequestTypeEnum = "Name"
	PREREQUISITE CourseUpdateRequestTypeEnum = "Prerequisite"
)

func IsValidCourseUpdateRequestTypeEnum(value string) bool {
	switch CourseUpdateRequestTypeEnum(value) {
	case NAME, PREREQUISITE:
		return true
	default:
		return false
	}
}
