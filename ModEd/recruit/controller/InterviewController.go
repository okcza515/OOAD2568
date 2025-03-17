package controller

import (
	"ModEd/recruit/model"
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
