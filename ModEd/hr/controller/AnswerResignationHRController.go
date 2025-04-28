package controller

import "gorm.io/gorm"

func HandleResignationStatus(tx *gorm.DB, id string, status string, reason string) error {
	hrFacade := NewHRFacade(tx) // Assuming nil is acceptable for the transaction in this context

	return hrFacade.UpdateResignationStudentStatus(id, status, reason)
}
