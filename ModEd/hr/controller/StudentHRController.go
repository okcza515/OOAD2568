package controller

import (
	commonModel "ModEd/common/model"
	"ModEd/hr/model"

	"gorm.io/gorm"
)

type StudentHRController struct {
	db *gorm.DB
}

// CreateStudentHRController creates a new instance of StudentHRController
// and automigrates the StudentInfo model.
func CreateStudentHRController(db *gorm.DB) *StudentHRController {
	db.AutoMigrate(&model.StudentInfo{})
	return &StudentHRController{db: db}
}

// GetAll returns all StudentInfo records.
func (c *StudentHRController) GetAll() ([]model.StudentInfo, error) {
	var infos []model.StudentInfo
	err := c.db.Find(&infos).Error
	return infos, err
}

// GetById retrieves a student's HR information by SID.
func (c *StudentHRController) GetById(sid string) (*model.StudentInfo, error) {
	var studentInfo model.StudentInfo
	if err := c.db.Where("student_code = ?", sid).First(&studentInfo).Error; err != nil {
		return nil, err
	}
	return &studentInfo, nil
}

// Insert inserts a new StudentInfo record.
func (c *StudentHRController) Insert(info *model.StudentInfo) error {
	return c.db.Create(info).Error
}

// Update updates an existing StudentInfo record.
func (c *StudentHRController) Update(info *model.StudentInfo) error {
	return c.db.Save(info).Error
}

// Delete deletes a student's HR information by SID.
func (c *StudentHRController) Delete(sid string) error {
	return c.db.Where("student_code = ?", sid).Delete(&model.StudentInfo{}).Error
}

// UpdateStatus updates the status of a student by SID.
func (c *StudentHRController) UpdateStatus(sid string, status commonModel.StudentStatus) error {
	// First retrieve the student record
	var studentInfo model.StudentInfo
	if err := c.db.Where("student_code = ?", sid).First(&studentInfo).Error; err != nil {
		return err
	}

	// Update the status field
	studentInfo.Status = &status

	// Save the updated record
	return c.db.Save(&studentInfo).Error
}

// Upsert inserts or updates a StudentInfo record.
func (c *StudentHRController) Upsert(info *model.StudentInfo) error {
	return c.db.FirstOrCreate(info).Error
}