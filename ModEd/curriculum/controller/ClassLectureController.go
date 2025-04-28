// MEP-1008

package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ClassLectureService interface {

	// AddClassLecture(lecture *model.ClassLecture) error
	Insert(data model.ClassLecture) error
	//GetClassLecturesByClassId(classId uint) ([]model.ClassLecture, error)
	RetrieveByID(id uint, preloads ...string) (*model.ClassLecture, error)
	//DeleteClassLecture(lectureId uint) error
	DeleteByID(id uint) error
	//UpdateClassLecture(lecture *model.ClassLecture) error
	UpdateByID(data model.ClassLecture) error
}

type ClassLectureController struct {
	*core.BaseController[*model.ClassLecture]
	Connector *gorm.DB
}

func CreateClassLectureController(db *gorm.DB) *ClassLectureController {
	return &ClassLectureController{
		BaseController: core.NewBaseController[*model.ClassLecture](db),
		Connector:      db,
	}
}

// func (c *ClassLectureController) AddClassLecture(lecture *model.ClassLecture) error {
// 	if err := c.Connector.Create(lecture).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *ClassLectureController) GetClassLecturesByClassId(classId uint) ([]model.ClassLecture, error) {
// 	var lectures []model.ClassLecture
// 	if err := c.Connector.Where("class_id = ?", classId).Find(&lectures).Error; err != nil {
// 		return nil, err
// 	}
// 	return lectures, nil
// }

// func (c *ClassLectureController) DeleteClassLecture(lectureId uint) error {
// 	if err := c.Connector.Delete(&model.ClassLecture{}, lectureId).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *ClassLectureController) UpdateClassLecture(lecture *model.ClassLecture) error {
// 	if err := c.Connector.Save(lecture).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
