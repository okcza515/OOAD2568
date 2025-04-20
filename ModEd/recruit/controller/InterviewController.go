// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"
	"errors"

	"gorm.io/gorm"
)

type InterviewController struct {
	*core.BaseController
	DB *gorm.DB
}

func CreateInterviewController(db *gorm.DB) *InterviewController {
	return &InterviewController{
		DB:             db,
		BaseController: core.NewBaseController("Interview", db),
	}
}

func (c *InterviewController) CreateInterview(interview *model.Interview) error {
	return c.Insert(interview)
}

func (c *InterviewController) DeleteInterview(id uint) error {
	return c.DeleteByID(id)
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

	// คืนค่า interview_status จากฐานข้อมูล
	return interview.InterviewStatus, nil
}

func GetInterviewDetails(db *gorm.DB, applicantID uint) (*model.Interview, error) {
	var interview model.Interview

	err := db.Where("applicant_id = ?", applicantID).First(&interview).Error
	if err != nil {
		return nil, errors.New("ไม่พบข้อมูลสัมภาษณ์สำหรับผู้สมัครที่ให้มา")
	}
	return &interview, nil
}
