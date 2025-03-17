package model

import "gorm.io/gorm"

type CourseUpdateStrategy interface {
	Update() error
}

type CourseUpdate struct {
	Connector *gorm.DB
}

func (c *CourseUpdateRequest) SetCourseUpdateStrategy(updateStrategy CourseUpdateStrategy) {
	c.UpdateStrategy = updateStrategy
}

func (c *CourseUpdateRequest) ExecuteUpdate() error {
	if c == nil {
		return nil
	}
	return c.UpdateStrategy.Update()
}
