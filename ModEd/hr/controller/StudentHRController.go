package controller

import (
	commonController "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/core"
	"ModEd/hr/model"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
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
	// common := &commonModel.Student{
	// 	StudentCode: studentCode,
	// 	FirstName:   firstName,
	// 	LastName:    lastName,
	// 	Email:       email,
	// }
	// if err := commonController.CreateStudentController(tx).Create(common); err != nil {
	// 	return fmt.Errorf("common.Create failed: %w", err)
	// }

	// 2) migrate to HR
	if err := MigrateStudentsToHR(tx); err != nil {
		return fmt.Errorf("MigrateStudentsToHR failed: %w", err)
	}

	// 3) build HR info & insert
	hrInfo := model.NewStudentInfo(studentCode, gender, citizenID, phone)

	// Insert the new HR record using the StudentHRController directly.
	if err := createStudentHRController(tx).insert(hrInfo); err != nil {
		return fmt.Errorf("failed to insert HR student info: %w", err)
	}
	return nil
}

func DeleteStudent(tx *gorm.DB, studentID string) error {
	// Delete student from common data.
	studentController := commonController.CreateStudentController(tx)
	if err := studentController.DeleteByCode(studentID); err != nil {
		return fmt.Errorf("failed to delete student from common data: %w", err)
	}

	return createStudentHRController(tx).delete(studentID)
}

func UpdateStudentInfo(tx *gorm.DB, studentID, firstName, lastName, gender, citizenID, phoneNumber, email string) error {
	// Wrap the business logic in a transaction.
	return tx.Transaction(func(tx *gorm.DB) error {
		// Retrieve the existing HR info using StudentHRController.
		controller := createStudentHRController(tx)
		studentInfo, err := controller.getById(studentID)
		if err != nil {
			return fmt.Errorf("error retrieving student with ID %s: %v", studentID, err)
		}

		updatedStudent := model.NewUpdatedStudentInfo(studentInfo, firstName, lastName, gender, citizenID, phoneNumber, email)

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

		// 3) Update HR-specific student info using the controller directly.
		if err := controller.update(updatedStudent); err != nil {
			return fmt.Errorf("failed to update student HR info: %v", err)
		}
		return nil
	})
}

func ImportStudents(tx *gorm.DB, filepath string) error {

	hrMapper, err := core.CreateMapper[model.StudentInfo](filepath)
	if err != nil {
		return fmt.Errorf("failed to create HR mapper: %v", err)
	}

	hrRecords := hrMapper.Deserialize()
	hrRecordsMap := make(map[string]model.StudentInfo)
	for _, hrRec := range hrRecords {
		if _, exists := hrRecordsMap[hrRec.StudentCode]; exists {
			return fmt.Errorf("duplicate student code found: %s", hrRec.StudentCode)
		}
		hrRecordsMap[hrRec.StudentCode] = *hrRec
	}

	controller := createStudentHRController(tx)
	for _, hrRec := range hrRecordsMap {
		studentInfo, err := controller.getById(hrRec.StudentCode)
		if err != nil {
			return fmt.Errorf("error retrieving student with ID %s: %v", hrRec.StudentCode, err)
		}

		importStudent := model.NewUpdatedStudentInfo(
			studentInfo,
			studentInfo.FirstName,
			studentInfo.LastName,
			hrRec.Gender,
			hrRec.CitizenID,
			hrRec.PhoneNumber,
			studentInfo.Email,
		)

		if err := controller.update(importStudent); err != nil {
			return fmt.Errorf("failed to upsert student %s: %v", importStudent.StudentCode, err)
		}
	}
	return nil
}

func ExportStudents(tx *gorm.DB, filePath string, format string) error {

	fileInfo, err := os.Stat(filePath)
	if err == nil && fileInfo.IsDir() {
		switch format {
		case "csv":
			filePath = fmt.Sprintf("%s/studentinfo.csv", filePath)
		case "json":
			filePath = fmt.Sprintf("%s/studentinfo.json", filePath)
		default:
			return fmt.Errorf("invalid format. Supported formats are 'csv' and 'json'")
		}
	}

	// Fetch all student records
	controller := createStudentHRController(tx)
	studentInfos, err := controller.getAll()
	if err != nil {
		return fmt.Errorf("error fetching students: %v", err)
	}

	// Handle export based on the format
	switch format {
	case "csv":
		// Use gocsv for CSV serialization
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}
		defer file.Close()

		if err := gocsv.MarshalFile(&studentInfos, file); err != nil {
			return fmt.Errorf("error exporting to CSV: %v", err)
		}
	case "json":
		// Use encoding/json for JSON serialization
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(studentInfos); err != nil {
			return fmt.Errorf("error exporting to JSON: %v", err)
		}
	default:
		return fmt.Errorf("invalid format. Supported formats are 'csv' and 'json'")
	}

	return nil
}
