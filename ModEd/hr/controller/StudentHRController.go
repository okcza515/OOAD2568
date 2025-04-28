package controller

import (
	commonController "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/hr/model"
	hrModel "ModEd/hr/model"
	"fmt"

	"gorm.io/gorm"
)

type StudentHRController struct {
	db *gorm.DB
}

// CreateStudentHRController creates a new instance of StudentHRController
// and automigrates the StudentInfo model.
func createStudentHRController(db *gorm.DB) *StudentHRController {
	db.AutoMigrate(&model.StudentInfo{})
	return &StudentHRController{db: db}
}

// GetAll returns all StudentInfo records.
func (c *StudentHRController) getAll() ([]*model.StudentInfo, error) {
	var infos []*model.StudentInfo
	err := c.db.Find(&infos).Error
	return infos, err
}

// GetById retrieves a student's HR information by SID.
func (c *StudentHRController) getById(sid string) (*model.StudentInfo, error) {
	var studentInfo model.StudentInfo
	if err := c.db.Where("student_code = ?", sid).First(&studentInfo).Error; err != nil {
		return nil, err
	}
	return &studentInfo, nil
}

// Insert inserts a new StudentInfo record.
func (c *StudentHRController) insert(info *model.StudentInfo) error {
	return c.db.Create(info).Error
}

// Update updates an existing StudentInfo record.
func (c *StudentHRController) update(info *model.StudentInfo) error {
	return c.db.Model(&model.StudentInfo{}).
		Where("student_code = ?", info.StudentCode).
		Updates(info).Error
}

// Delete deletes a student's HR information by SID.
func (c *StudentHRController) delete(sid string) error {
	return c.db.Where("student_code = ?", sid).Delete(&model.StudentInfo{}).Error
}

// UpdateStatus updates the status of a student by SID.
func (c *StudentHRController) updateStatus(sid string, status commonModel.StudentStatus) error {
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

func AddStudent(tx *gorm.DB,
	studentCode, firstName, lastName, email, gender, citizenID, phone string,
) error {
	// 1) common record
	common := &commonModel.Student{
		StudentCode: studentCode,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
	}
	if err := commonController.CreateStudentController(tx).Create(common); err != nil {
		return fmt.Errorf("common.Create failed: %w", err)
	}

	// 2) migrate to HR
	if err := MigrateStudentsToHR(tx); err != nil {
		return fmt.Errorf("MigrateStudentsToHR failed: %w", err)
	}

	// 3) build HR info & upsert
	builder := hrModel.NewStudentInfoBuilder().
		WithStudentCode(studentCode).
		WithFirstName(firstName).
		WithLastName(lastName).
		WithGender(gender).
		WithCitizenID(citizenID).
		WithPhoneNumber(phone)

	hrInfo, err := builder.Build()
	if err != nil {
		return fmt.Errorf("build HR info failed: %w", err)
	}

	// Upsert so we handle both insert & update
	if err := NewHRFacade(tx).UpsertStudent(hrInfo); err != nil {
		return fmt.Errorf("HR.UpsertStudent failed: %w", err)
	}
	return nil
}
