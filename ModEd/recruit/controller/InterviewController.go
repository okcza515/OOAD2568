package controller

import (
	"ModEd/recruit/model"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type InterviewController struct {
	DB *gorm.DB
}

func CreateInterviewController(db *gorm.DB) *InterviewController {
	controller := &InterviewController{DB: db}

	if err := db.AutoMigrate(&model.Interview{}); err != nil {
		fmt.Println("❌ Failed to migrate interviews table:", err)
	}

	return controller
}

func (c *InterviewController) CreateInterview(interview *model.Interview) error {
	return c.DB.Create(interview).Error
}

func (c *InterviewController) DeleteInterview(id uuid.UUID) error {
	return c.DB.Delete(&model.Interview{}, id).Error
}

func GetApplicationStatus(db *gorm.DB, applicantID uuid.UUID) (string, error) {
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

func GetInterviewDetails(db *gorm.DB, applicantID uuid.UUID) (*model.Interview, error) {
	var interview model.Interview

	err := db.Where("applicant_id = ?", applicantID).First(&interview).Error
	if err != nil {
		return nil, errors.New("ไม่พบข้อมูลสัมภาษณ์สำหรับผู้สมัครที่ให้มา")
	}
	return &interview, nil
}
