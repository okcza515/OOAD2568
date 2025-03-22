package controller

import (
	"gorm.io/gorm"

	modelCurriculum "ModEd/curriculum/model"
)

type IClassController interface {
	CreateClass(class modelCurriculum.Class) (classId uint, err error)
	GetClass(classId uint) (class *modelCurriculum.Class, err error)
	GetClasses() (classes []*modelCurriculum.Class, err error)
	UpdateClass(updatedClass modelCurriculum.Class) (class *modelCurriculum.Class, err error)
	DeleteClass(classId uint) (class *modelCurriculum.Class, err error)
}

type ClassController struct {
	db *gorm.DB
}

func NewClassController(db *gorm.DB) IClassController {
	return &ClassController{db: db}
}

func (c *ClassController) CreateClass(class modelCurriculum.Class) (classId uint, err error) {
	if err := c.db.Create(&class).Error; err != nil {
		return 0, err
	}
	return class.ID, nil
}

func (c *ClassController) GetClass(classId uint) (class *modelCurriculum.Class, err error) {
	class = &modelCurriculum.Class{}
	if err := c.db.First(class, classId).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (c *ClassController) GetClasses() (classes []*modelCurriculum.Class, err error) {
	if err := c.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (c *ClassController) UpdateClass(updatedClass modelCurriculum.Class) (class *modelCurriculum.Class, err error) {
	class = &modelCurriculum.Class{}
	if err := c.db.First(class, updatedClass.ID).Error; err != nil {
		return nil, err
	}
	class.CourseId = updatedClass.CourseId
	class.ClassId = updatedClass.ClassId
	class.Schedule = updatedClass.Schedule
	class.Section = updatedClass.Section
	if err := c.db.Save(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (c *ClassController) DeleteClass(classId uint) (class *modelCurriculum.Class, err error) {
	class = &modelCurriculum.Class{}
	if err := c.db.First(class, classId).Error; err != nil {
		return nil, err
	}
	if err := c.db.Delete(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}
