package controller

import (
	"fmt"

	common "ModEd/common/model"
	"ModEd/hr/model"
	hr "ModEd/hr/model"

	"gorm.io/gorm"
)

// MigrateStudentsToHR migrates student data from the common Student table
// to the HR StudentInfo table.
func MigrateStudentsToHR(db *gorm.DB) error {
	// Automigrate the HR StudentInfo model.
	if err := db.AutoMigrate(&hr.StudentInfo{}); err != nil {
		return fmt.Errorf("failed to automigrate HR StudentInfo: %w", err)
	}

	// Retrieve student data from the common Student table.
	var students []common.Student
	if err := db.Find(&students).Error; err != nil {
		return fmt.Errorf("failed to retrieve common students: %w", err)
	}

	// Migrate data to HR StudentInfo.
	for _, s := range students {
		migrateStudent := model.NewStudentInfo(s.StudentCode, "", "", "").SetStudent(s)
			

		// Use FirstOrCreate to avoid duplicate unique errors.
		if err := db.Where("student_code = ?", s.StudentCode).FirstOrCreate(&migrateStudent).Error; err != nil {
			return fmt.Errorf("failed to migrate student %s: %w", s.StudentCode, err)
		}
		// fmt.Printf("Migrated student %s successfully\n", s.SID)
	}

	return nil
}
