package controller

import (
	"ModEd/hr/model"
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type StudentHRController struct {
	db *gorm.DB
}

// NewStudentHRController creates a new instance of StudentHRController
// and automigrates the StudentInfo model.
func NewStudentHRController(db *gorm.DB) *StudentHRController {
	db.AutoMigrate(&model.StudentInfo{})
	return &StudentHRController{db: db}
}

// ListAllStudentInfos returns all StudentInfo records.
func (c *StudentHRController) ListAllStudentInfos() ([]model.StudentInfo, error) {
	var infos []model.StudentInfo
	err := c.db.Find(&infos).Error
	return infos, err
}

// RetrieveStudentInfo retrieves a student's HR information by SID.
func (c *StudentHRController) RetrieveStudentInfo(sid string) (*model.StudentInfo, error) {
	var studentInfo model.StudentInfo
	if err := c.db.Where("s_id = ?", sid).First(&studentInfo).Error; err != nil {
		return nil, err
	}
	return &studentInfo, nil
}

// InsertStudentInfo inserts a new StudentInfo record.
func (c *StudentHRController) InsertStudentInfo(info *model.StudentInfo) error {
	return c.db.Create(info).Error
}

// UpdateStudentInfo updates an existing StudentInfo record.
func (c *StudentHRController) UpdateStudentInfo(info *model.StudentInfo) error {
	return c.db.Save(info).Error
}

// DeleteStudentInfo deletes a student's HR information by SID.
func (c *StudentHRController) DeleteStudentInfo(sid string) error {
	return c.db.Where("s_id = ?", sid).Delete(&model.StudentInfo{}).Error
}

func (c *StudentHRController) UpdateStudentStatus(sid string, status commonModel.StudentStatus) error {
    // First retrieve the student record
    var studentInfo model.StudentInfo
    if err := c.db.Where("s_id = ?", sid).First(&studentInfo).Error; err != nil {
        return err
    }
    
    // Update the status field
    studentInfo.Status = status
    
    // Save the updated record
    return c.db.Save(&studentInfo).Error
}