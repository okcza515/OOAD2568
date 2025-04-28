package controller

import (
	"ModEd/hr/util"
	"fmt"

	"gorm.io/gorm"
)

func UpdateStudentStatusBusinessLogic(tx *gorm.DB, studentID, statusStr string) error {
	// Convert status string to the appropriate enum or type using our util layer.
	newStatus, err := util.StatusFromString(statusStr)
	if err != nil {
		return fmt.Errorf("invalid status input: %v", err)
	}

	// Use the provided transaction/db connection to create the HRFacade.
	hrFacade := NewHRFacade(tx)

	// Call the business logic to update the student status.
	return hrFacade.UpdateStudentStatus(studentID, newStatus)
}
