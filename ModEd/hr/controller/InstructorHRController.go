package controller

import (
	"ModEd/hr/model"

	"gorm.io/gorm"
)

type InstructorHRController struct {
	db *gorm.DB
}

// CreateInstructorHRController creates a new instance of InstructorHRController
// and automigrates the InstructorInfo model.
func createInstructorHRController(db *gorm.DB) *InstructorHRController {
	db.AutoMigrate(&model.InstructorInfo{})
	return &InstructorHRController{db: db}
}

// GetAll returns all InstructorInfo records.
func (c *InstructorHRController) getAll() ([]*model.InstructorInfo, error) {
	var infos []*model.InstructorInfo
	err := c.db.Find(&infos).Error
	return infos, err
}

// GetById retrieves an instructor's HR information by ID.
func (c *InstructorHRController) getById(id string) (*model.InstructorInfo, error) {
	var instructorInfo model.InstructorInfo
	if err := c.db.Where("instructor_id = ?", id).First(&instructorInfo).Error; err != nil {
		return nil, err
	}
	return &instructorInfo, nil
}

// Insert inserts a new InstructorInfo record.
func (c *InstructorHRController) insert(info *model.InstructorInfo) error {
	return c.db.Create(info).Error
}

// Update updates an existing InstructorInfo record.
// No Primary Key for InstructorInfo
func (c *InstructorHRController) update(info *model.InstructorInfo) error {
	return c.db.Model(&model.InstructorInfo{}).
		Where("id = ?", info.ID).
		Updates(info).Error
}

// Delete deletes an instructor's HR information by ID.
func (c *InstructorHRController) delete(id string) error {
	return c.db.Where("instructor_id = ?", id).Delete(&model.InstructorInfo{}).Error
}

// Upsert inserts or updates a InstructorInfo record.
func (c *InstructorHRController) upsert(info *model.InstructorInfo) error {
	var existingInfo model.InstructorInfo
	if err := c.db.Where("instructor_id = ?", info.ID).First(&existingInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Insert(info)
		}
		return err
	}
	return c.Update(info)
}
