package controller

import (
	commonController "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/hr/model"
	"ModEd/hr/util"
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
	studentCode string, firstName string, lastName string, gender string, citizenID string, phone string, email string,
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
	hrInfo := model.NewStudentInfo(studentCode, gender, citizenID, phone)

	// // Upsert so we handle both insert & update
	if err := NewHRFacade(tx).UpsertStudent(hrInfo); err != nil {
		return fmt.Errorf("HR.UpsertStudent failed: %w", err)
	}
	return nil
}

func DeleteStudent(tx *gorm.DB, studentID string) error {
	// Delete student from common data.
	studentController := commonController.CreateStudentController(tx)
	if err := studentController.DeleteByCode(studentID); err != nil {
		return fmt.Errorf("failed to delete student from common data: %w", err)
	}

	hrFacade := NewHRFacade(tx)

	return hrFacade.DeleteStudent(studentID)
}

func UpdateStudentInfo(tx *gorm.DB, studentID, firstName, lastName, gender, citizenID, phoneNumber, email string) error {
	// Wrap the business logic in a transaction.
	return tx.Transaction(func(tx *gorm.DB) error {
		// Use HRFacade to get the existing HR info.
		hrFacade := NewHRFacade(tx)
		studentInfo, err := hrFacade.GetStudentById(studentID)
		if err != nil {
			return fmt.Errorf("error retrieving student with ID %s: %v", studentID, err)
		}

		// Update student info using non-empty values.
		updatedStudent := studentInfo.
			SetFirstName(util.IfNotEmpty(firstName, studentInfo.FirstName)).
			SetLastName(util.IfNotEmpty(lastName, studentInfo.LastName)).
			SetGender(util.IfNotEmpty(gender, studentInfo.Gender)).
			SetCitizenID(util.IfNotEmpty(citizenID, studentInfo.CitizenID)).
			SetPhoneNumber(util.IfNotEmpty(phoneNumber, studentInfo.PhoneNumber)).
			SetEmail(util.IfNotEmpty(email, studentInfo.Email))

		// 1) Update common student data.
		studentData := map[string]any{
			"FirstName": updatedStudent.FirstName,
			"LastName":  updatedStudent.LastName,
			"Email":     updatedStudent.Email,
		}
		studentController := commonController.CreateStudentController(tx)
		if err := studentController.Update(studentID, studentData); err != nil {
			return fmt.Errorf("failed to update common student data: %v", err)
		}

		// 2) Migrate students to HR.
		if err := MigrateStudentsToHR(tx); err != nil {
			return fmt.Errorf("failed to migrate student to HR module: %v", err)
		}

		// 3) Upsert HR-specific student info.
		if err := hrFacade.UpsertStudent(updatedStudent); err != nil {
			return fmt.Errorf("failed to update student HR info: %v", err)
		}
		return nil
	})
}
