package model

type StudentRequestTypeEnum string

const (
	SickLeave     StudentRequestTypeEnum = "SickLeave"
	PersonalLeave StudentRequestTypeEnum = "PersonalLeave"
	AddCourse     StudentRequestTypeEnum = "ExtraCourseEnrollment"
)
