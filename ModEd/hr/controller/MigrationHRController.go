package controller

import (
	"fmt"

	common "ModEd/common/model"
	hr "ModEd/hr/model"

	"gorm.io/gorm"
)

// MigrateStudentsToHR migrates student data from the common Student table
// to the HR StudentInfo table.
func MigrateStudentsToHR(db *gorm.DB) error {
	// Automigrate the HR StudentInfo model.
	if err := db.AutoMigrate(&hr.StudentInfo{}); err != nil {
		return fmt.Errorf("Failed to automigrate HR StudentInfo: %w", err)
	}

	// Retrieve student data from the common Student table.
	var students []common.Student
	if err := db.Find(&students).Error; err != nil {
		return fmt.Errorf("Failed to retrieve common students: %w", err)
	}

	// Migrate data to HR StudentInfo.
	for _, s := range students {
		studentInfo := hr.StudentInfo{
			Student:     s,
			Gender:      "", // default value; update as needed
			CitizenID:   "", // default value; update as needed
			PhoneNumber: "", // default value; update as needed
		}

		// Use FirstOrCreate to avoid duplicate unique errors.
		if err := db.Where("student_code = ?", s.StudentCode).FirstOrCreate(&studentInfo).Error; err != nil {
			return fmt.Errorf("Failed to migrate student %s: %w", s.StudentCode, err)
		}
		// fmt.Printf("Migrated student %s successfully\n", s.SID)
	}

	fmt.Println("Migration completed successfully.")

	return nil
}

func SynchronizeStudents(db *gorm.DB) error {
	// Retrieve all HR student records.
	var hrStudents []hr.StudentInfo
	if err := db.Find(&hrStudents).Error; err != nil {
		return fmt.Errorf("failed to retrieve HR students: %w", err)
	}

	for _, hrRec := range hrStudents {
		sid := hrRec.Student.StudentCode

		// Retrieve the common student record.
		var commonStudent common.Student
		err := db.Where("student_code = ?", sid).First(&commonStudent).Error
		if err != nil {
			// If common student not found, create it using HR data
			commonStudent = syncCreateCommonFromHR(hrRec.Student)
			if err := db.Create(&commonStudent).Error; err != nil {
				return fmt.Errorf("failed to create common student %s: %w", sid, err)
			}
		} else {
			// If common student exists, update it if fields differ
			if syncUpdateCommonFromHR(&commonStudent, hrRec.Student) {
				if err := db.Save(&commonStudent).Error; err != nil {
					return fmt.Errorf("failed to update common student %s: %w", sid, err)
				}
			}
		}

		// Push any changes from the common student record back to the HR record
		if syncUpdateHRFromCommon(&hrRec, commonStudent) {
			if err := db.Save(&hrRec).Error; err != nil {
				return fmt.Errorf("failed to update HR student %s: %w", sid, err)
			}
		}
	}
	fmt.Println("Two-way synchronization completed successfully.")
	return nil
}

// syncCreateCommonFromHR creates a new common.Student from HR student data
func syncCreateCommonFromHR(hrStudent common.Student) common.Student {
	return common.Student{
		StudentCode:       hrStudent.StudentCode,
		FirstName: hrStudent.FirstName,
		LastName:  hrStudent.LastName,
		Email:     hrStudent.Email,
		StartDate: hrStudent.StartDate,
		BirthDate: hrStudent.BirthDate,
		Program:   hrStudent.Program,
		Status:    hrStudent.Status,
		// Add new fields here when they're added to the model
	}
}

// syncUpdateCommonFromHR updates common.Student from HR data and returns true if updated
func syncUpdateCommonFromHR(common *common.Student, hr common.Student) bool {
	updated := false

	// Define field mappings for synchronization
	fieldMappings := []struct {
		commonField *string
		hrField     string
	}{
		{&common.FirstName, hr.FirstName},
		{&common.LastName, hr.LastName},
		{&common.Email, hr.Email},
		// Add new string fields here
	}

	// Update string fields
	for _, mapping := range fieldMappings {
		if *mapping.commonField != mapping.hrField {
			*mapping.commonField = mapping.hrField
			updated = true
		}
	}

	if common.Program != hr.Program {
		common.Program = hr.Program
		updated = true
	}

	if common.Status != hr.Status {
		common.Status = hr.Status
		updated = true
	}

	if !common.StartDate.Equal(hr.StartDate) {
		common.StartDate = hr.StartDate
		updated = true
	}
	if !common.BirthDate.Equal(hr.BirthDate) {
		common.BirthDate = hr.BirthDate
		updated = true
	}

	return updated
}

// syncUpdateHRFromCommon updates HR.StudentInfo from common data and returns true if updated
func syncUpdateHRFromCommon(hr *hr.StudentInfo, common common.Student) bool {
	updated := false

	// Define field mappings for synchronization
	fieldMappings := []struct {
		hrField     *string
		commonField string
	}{
		{&hr.Student.FirstName, common.FirstName},
		{&hr.Student.LastName, common.LastName},
		{&hr.Student.Email, common.Email},
		// Add new string fields here
	}

	// Update string fields
	for _, mapping := range fieldMappings {
		if *mapping.hrField != mapping.commonField {
			*mapping.hrField = mapping.commonField
			updated = true
		}
	}

	if hr.Student.Program != common.Program {
		hr.Student.Program = common.Program
		updated = true
	}

	if hr.Student.Status != common.Status {
		hr.Student.Status = common.Status
		updated = true
	}

	if !hr.Student.StartDate.Equal(common.StartDate) {
		hr.Student.StartDate = common.StartDate
		updated = true
	}
	if !hr.Student.BirthDate.Equal(common.BirthDate) {
		hr.Student.BirthDate = common.BirthDate
		updated = true
	}

	return updated
}
