// MEP-1002
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
)

type ClassController struct {
	db *gorm.DB
}

func NewClassController(db *gorm.DB) *ClassController {
	return &ClassController{db: db}
}

func (c *ClassController) CreateClass(class *model.Class) (classId uint, err error) {
	if err := c.db.Create(&class).Error; err != nil {
		return 0, err
	}
	return class.ClassId, nil
}

func (c *ClassController) GetClass(classId uint) (class *model.Class, err error) {
	class = &model.Class{}
	if err := c.db.First(class, classId).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (c *ClassController) GetClasses() (classes []*model.Class, err error) {
	if err := c.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (c *ClassController) UpdateClass(updatedClass *model.Class) (class *model.Class, err error) {
	class = &model.Class{}
	if err := c.db.First(class, updatedClass.ClassId).Error; err != nil {
		return nil, err
	}
	class.CourseId = updatedClass.CourseId
	class.ClassId = updatedClass.ClassId
	class.Schedule = updatedClass.Schedule
	class.Section = updatedClass.Section

	if err := c.db.Updates(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (c *ClassController) DeleteClass(classId uint) (class *model.Class, err error) {
	class = &model.Class{}
	if err := c.db.First(class, classId).Error; err != nil {
		return nil, err
	}
	if err := c.db.Delete(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (c *ClassController) CreateSeedClass(path string) (classes []*model.Class, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file deserializer")
	}

	if err := deserializer.Deserialize(&classes); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize classes")
	}

	for _, class := range classes {
		_, err := c.CreateClass(class)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create class")
		}
	}
	fmt.Println("Create Class Seed Successfully")
	return classes, nil
}
