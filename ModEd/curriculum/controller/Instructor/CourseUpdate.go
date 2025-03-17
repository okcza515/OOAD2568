package controller

import (
	"ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type CourseUpdate struct {
	Connector *gorm.DB
}

func (courseUpdate *CourseUpdate) CourseUpdate(courseUpdateRequest *model.CourseUpdateRequest) error {
	var updater model.CourseUpdateStrategy

	switch courseUpdateRequest.RequestType {
	case model.NAME:
		//
	case model.PREREQUISITE:
		//
	default:
		return fmt.Errorf("invalid request type")
	}

	courseUpdateRequest.SetCourseUpdateStrategy(updater)
	return courseUpdateRequest.ExecuteUpdate()
}
