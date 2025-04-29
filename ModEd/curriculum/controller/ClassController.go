// MEP-1002
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"ModEd/core"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
)

type ClassController struct {
	db   *gorm.DB
	core *core.BaseController[*model.Class]
}

type ClassControllerInterface interface {
	CreateClass(class *model.Class) (classId uint, err error)
	GetClass(classId uint, preload ...string) (class *model.Class, err error)
	GetClasses(preload ...string) (classes []*model.Class, err error)
	UpdateClass(updatedClass *model.Class) (class *model.Class, err error)
	DeleteClass(classId uint) (class *model.Class, err error)
	CreateSeedClass(path string) (classes []*model.Class, err error)
}

func NewClassController(db *gorm.DB) *ClassController {
	return &ClassController{
		db:   db,
		core: core.NewBaseController[*model.Class](db),
	}
}

func (c *ClassController) CreateClass(class *model.Class) (classId uint, err error) {
	if err := c.core.Insert(class); err != nil {
		return 0, err
	}
	return class.ClassId, nil
}

func (c *ClassController) GetClass(classId uint, preload ...string) (class *model.Class, err error) {
	class, err = c.core.RetrieveByID(classId, preload...)
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (c *ClassController) GetClasses(preload ...string) (classes []*model.Class, err error) {
	classes, err = c.core.List(nil, preload...)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (c *ClassController) UpdateClass(updatedClass *model.Class) (class *model.Class, err error) {
	class, err = c.core.RetrieveByID(updatedClass.ClassId)
	if err != nil {
		return nil, err
	}

	class.CourseId = updatedClass.CourseId
	class.Schedule = updatedClass.Schedule
	class.Section = updatedClass.Section

	if err := c.core.UpdateByID(class); err != nil {
		return nil, err
	}
	return class, nil
}

func (c *ClassController) DeleteClass(classId uint) (class *model.Class, err error) {
	class, err = c.core.RetrieveByID(classId)
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByID(classId); err != nil {
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
