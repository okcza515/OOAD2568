// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"
	"errors"

	"gorm.io/gorm"
)

type InterviewController struct {
	Base *core.BaseController[*model.Interview]
	DB   *gorm.DB
}

func CreateInterviewController(db *gorm.DB) *InterviewController {
	return &InterviewController{
		DB:   db,
		Base: core.NewBaseController[*model.Interview](db),
	}
}

func (c *InterviewController) CreateInterview(interview *model.Interview) error {
	return c.Base.Insert(interview)
}

func (c *InterviewController) DeleteInterview(id uint) error {
	return c.Base.DeleteByID(id)
}

func GetApplicationStatus(db *gorm.DB, applicantID uint) (string, error) {
	var interview model.Interview

	err := db.Where("applicant_id = ?", applicantID).First(&interview).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("ไม่พบข้อมูลสัมภาษณ์ของผู้สมัคร")
		}
		return "", err
	}

	return string(interview.InterviewStatus), nil
}

func GetInterviewDetails(db *gorm.DB, applicantID uint) (*model.Interview, error) {
	var interview model.Interview

	err := db.Where("applicant_id = ?", applicantID).First(&interview).Error
	if err != nil {
		return nil, errors.New("ไม่พบข้อมูลสัมภาษณ์สำหรับผู้สมัครที่ให้มา")
	}
	return &interview, nil
}
