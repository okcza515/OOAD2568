package controller

import (
	commonController "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/core"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
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

func (c *StudentHRController) GetAllStudents() ([]*model.StudentInfo, error) {
	studentInfos, err := c.getAll()
	if err != nil {
		return nil, fmt.Errorf("error fetching students: %v", err)
	}
	return studentInfos, nil
}

func (c *StudentHRController) AddStudent(
	studentCode string, firstName string, lastName string, email string, startDate string, birthdate string, program string, department string, status string,
	gender string, citizenID string, phoneNumber string,
) error {
	startDateParsed, err := time.Parse("02-01-2006", startDate)
	if err != nil {
		return fmt.Errorf("failed to parse start date: %w", err)
	}

	birthDateParsed, err := time.Parse("02-01-2006", birthdate)
	if err != nil {
		return fmt.Errorf("failed to parse birth date: %w", err)
	}

	programParsed, err := util.ProgramTypeFromString(program)
	if err != nil {
		return fmt.Errorf("failed to parse program: %w", err)
	}

	statusParse, err := util.StatusFromString(status)
	if err != nil {
		return fmt.Errorf("failed to parse status: %w", err)
	}

	tm := &util.TransactionManager{DB: c.db}

	err = tm.Execute(func(tx *gorm.DB) error {
		common := &commonModel.Student{
			StudentCode: studentCode,
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			StartDate:   startDateParsed,
			BirthDate:   birthDateParsed,
			Program:     programParsed,
			Department:  department,
			Status:      &statusParse,
		}

		if err := commonModel.ManualAddStudent(tx, common); err != nil {
			return fmt.Errorf("insert failed in common model: %w", err)
		}

		if migrateErr := MigrateStudentsToHR(tx); migrateErr != nil {
			return fmt.Errorf("migrateStudentsToHR failed: %w", migrateErr)
		}

		hrInfo := model.NewStudentInfo(*common, gender, citizenID, phoneNumber)

		studentController := NewStudentHRController(tx)
		if updateErr := studentController.update(hrInfo); updateErr != nil {
			return fmt.Errorf("failed to update HR student info: %w", updateErr)
		}
		return nil
	})
	return err
}

func (c *StudentHRController) DeleteStudent(studentID string) error {
	tm := &util.TransactionManager{DB: c.db}

	err := tm.Execute(func(tx *gorm.DB) error {
		commonStudentController := commonController.CreateStudentController(tx)
		if err := commonStudentController.DeleteByCode(studentID); err != nil {
			return fmt.Errorf("failed to delete student from common data: %w", err)
		}

		studentController := NewStudentHRController(tx)
		if err := studentController.delete(studentID); err != nil {
			return fmt.Errorf("failed to delete student HR info: %w", err)
		}

		return nil
	})
	return err
}

func (c *StudentHRController) UpdateStudentInfo(
	studentCode string, firstName string, lastName string, email string,
	gender string, citizenID string, phoneNumber string,
) error {
	tm := &util.TransactionManager{DB: c.db}
	err := tm.Execute(func(tx *gorm.DB) error {
		// Retrieve the existing HR info using StudentHRController.
		studentController := NewStudentHRController(tx)
		studentInfo, err := studentController.getById(studentCode)
		if err != nil {
			return fmt.Errorf("error retrieving student with ID %s: %v", studentCode, err)
		}

		updatedStudent := model.NewUpdatedStudentInfo(studentInfo, firstName, lastName, gender, citizenID, phoneNumber, email)

		// 1) Update common student data.
		studentData := map[string]any{
			"FirstName":  updatedStudent.FirstName,
			"LastName":   updatedStudent.LastName,
			"Email":      updatedStudent.Email,
			"Program":    updatedStudent.Program,
			"Department": updatedStudent.Department,
			"Status":     updatedStudent.Status,
		}

		commonStudentController := commonController.CreateStudentController(tx)
		if err := commonStudentController.Update(studentCode, studentData); err != nil {
			return fmt.Errorf("failed to update common student data: %v", err)
		}

		// 2) Migrate students to HR.
		if err := MigrateStudentsToHR(tx); err != nil {
			return fmt.Errorf("failed to migrate student to HR module: %v", err)
		}

		// 3) Update HR-specific student info using the controller directly.
		studentHRData := model.NewUpdatedStudentInfo(studentInfo, firstName, lastName, gender, citizenID, phoneNumber, email)
		if err := studentController.update(studentHRData); err != nil {
			return fmt.Errorf("failed to update student HR info: %v", err)
		}
		return nil
	})
	return err
}

func (c *StudentHRController) ImportStudents(filepath string) error {
	tm := &util.TransactionManager{DB: c.db}

	err := tm.Execute(func(tx *gorm.DB) error {
		hrMapper, err := core.CreateMapper[model.StudentInfo](filepath)
		if err != nil {
			return fmt.Errorf("failed to create HR mapper: %w", err)
		}

		hrRecords := hrMapper.Deserialize()
		hrRecordsMap := make(map[string]model.StudentInfo)
		for _, hrRec := range hrRecords {
			if _, exists := hrRecordsMap[hrRec.StudentCode]; exists {
				return fmt.Errorf("duplicate student code found in import file: %s", hrRec.StudentCode)
			}
			if hrRec != nil {
				hrRecordsMap[hrRec.StudentCode] = *hrRec
			} else {
				continue
			}
		}

		studentController := NewStudentHRController(tx)
		for studentCode, hrRec := range hrRecordsMap {
			studentInfo, err := studentController.getById(studentCode)
			if err != nil {
				return fmt.Errorf("error retrieving student with ID %s: %w", studentCode, err)
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

			if err := studentController.update(importStudent); err != nil {
				return fmt.Errorf("failed to update student %s: %w", importStudent.StudentCode, err)
			}
		}
		return nil
	})
	return err
}

func (c *StudentHRController) ExportStudents(tx *gorm.DB, filePath string, format string) error {
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
	studentInfos, err := c.getAll()
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

func (c *StudentHRController) MigrateStudentRecords() error {
	var commonStudents []commonModel.Student
	if err := c.db.Find(&commonStudents).Error; err != nil {
		return fmt.Errorf("failed to retrieve common students: %w", err)
	}

	for _, s := range commonStudents {
		studentInfo := model.StudentInfo{
			Student:     s,  // Embed the common student data
			Gender:      "", // Initialize HR fields as empty
			CitizenID:   "",
			PhoneNumber: "",
		}

		if err := c.db.Where("student_code = ?", s.StudentCode).
			FirstOrCreate(&studentInfo).Error; err != nil {
			return fmt.Errorf("failed to migrate student %s: %w", s.StudentCode, err)
		}
	}

	return nil
}
