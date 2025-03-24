package seed

import (
	controller "ModEd/curriculum/controller/curriculum"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

func CreateCourseSeed(db *gorm.DB, path string) (courses []*model.Course, err error) {

	courseController := controller.NewCourseController(db)
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file deserializer")
	}

	if err := deserializer.Deserialize(&courses); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize courses")
	}

	for _, course := range courses {
		_, err := courseController.CreateCourse(course)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create course")
		}
	}
	fmt.Println("Create Course Seed Successfully")
	return courses, nil
}
