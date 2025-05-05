// MEP-1002
package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CourseController struct {
	db   *gorm.DB
	core *core.BaseController[*model.Course]
}

type CourseControllerInterface interface {
	CreateCourse(course *model.Course) (courseId uint, err error)
	GetCourse(courseId uint, preload ...string) (course *model.Course, err error)
	GetCourses(preload ...string) (courses []*model.Course, err error)
	UpdateCourse(updatedCourse *model.Course) (*model.Course, error)
	DeleteCourse(courseId uint) (course *model.Course, err error)
	CreateSeedCourse(path string) (courses []*model.Course, err error)
}

func NewCourseController(db *gorm.DB) CourseControllerInterface {
	return &CourseController{
		db:   db,
		core: core.NewBaseController[*model.Course](db),
	}
}

func (c *CourseController) CreateCourse(course *model.Course) (courseId uint, err error) {
	if err := c.core.Insert(course); err != nil {
		return 0, err
	}
	return course.CourseId, nil
}

func (c *CourseController) GetCourse(courseId uint, preload ...string) (course *model.Course, err error) {
	course, err = c.core.RetrieveByCondition(map[string]interface{}{"course_id": courseId}, preload...)
	if err != nil {
		return nil, err
	}
	return course, nil
}

// TODO: Consider adding condtions or preload params
func (c *CourseController) GetCourses(preload ...string) (courses []*model.Course, err error) {
	courses, err = c.core.List(nil, preload...)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *CourseController) UpdateCourse(updatedCourse *model.Course) (course *model.Course, err error) {
	course, err = c.core.RetrieveByCondition(map[string]interface{}{"course_id": updatedCourse.CourseId})
	if err != nil {
		return nil, err
	}

	// Update fields
	course.Name = updatedCourse.Name
	course.Description = updatedCourse.Description
	course.Optional = updatedCourse.Optional
	course.CourseStatus = updatedCourse.CourseStatus
	// course.Prerequisite = updatedCourse.Prerequisite
	// course.ClassList = updatedCourse.ClassList

	if err := c.core.UpdateByCondition(map[string]interface{}{"course_id": updatedCourse.CourseId}, course); err != nil {
		return nil, err
	}
	return course, nil
}

func (c *CourseController) DeleteCourse(courseId uint) (course *model.Course, err error) {
	course, err = c.core.RetrieveByCondition(map[string]interface{}{"course_id": courseId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"course_id": courseId}); err != nil {
		return nil, err
	}
	return course, nil
}

func (c *CourseController) CreateSeedCourse(path string) (courses []*model.Course, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file deserializer")
	}

	if err := deserializer.Deserialize(&courses); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize courses")
	}

	for _, course := range courses {
		_, err := c.CreateCourse(course)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create course")
		}
	}
	fmt.Println("Create Course Seed Successfully")
	return courses, nil

}
