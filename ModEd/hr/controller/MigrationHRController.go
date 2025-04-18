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