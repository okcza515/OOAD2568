package controller

import (
	"ModEd/hr/model"

	"gorm.io/gorm"
)

type InstructorHRController struct {
	db *gorm.DB
}

// NewInstructorHRController creates a new instance of InstructorHRController
// and automigrates the InstructorInfo model.
func NewInstructorHRController(db *gorm.DB) *InstructorHRController {
	db.AutoMigrate(&model.InstructorInfo{})
	return &InstructorHRController{db: db}
}

// ListAllInstructorInfos returns all InstructorInfo records.
func (c *InstructorHRController) ListAllInstructorInfos() ([]model.InstructorInfo, error) {
	var infos []model.InstructorInfo
	err := c.db.Find(&infos).Error
	return infos, err
}

// RetrieveInstructorInfo retrieves an instructor's HR information by ID.
func (c *InstructorHRController) RetrieveInstructorInfo(id string) (*model.InstructorInfo, error) {
	var instructorInfo model.InstructorInfo
	if err := c.db.Where("instructor_id = ?", id).First(&instructorInfo).Error; err != nil {
		return nil, err
	}
	return &instructorInfo, nil
}

// InsertInstructorInfo inserts a new InstructorInfo record.
func (c *InstructorHRController) InsertInstructorInfo(info *model.InstructorInfo) error {
	return c.db.Create(info).Error
}

// UpdateInstructorInfo updates an existing InstructorInfo record.
func (c *InstructorHRController) UpdateInstructorInfo(info *model.InstructorInfo) error {
	return c.db.Save(info).Error
}

// DeleteInstructorInfo deletes an instructor's HR information by ID.
func (c *InstructorHRController) DeleteInstructorInfo(id string) error {
	return c.db.Where("instructor_id = ?", id).Delete(&model.InstructorInfo{}).Error
}
