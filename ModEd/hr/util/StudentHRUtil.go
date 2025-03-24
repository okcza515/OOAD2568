package util

import (
	"fmt"
	"os"

	commonModel "ModEd/common/model"
	"ModEd/hr/controller"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StatusToString(status commonModel.StudentStatus) string {
	switch status {
	case commonModel.ACTIVE:
		return "ACTIVE"
	case commonModel.GRADUATED:
		return "GRADUATED"
	case commonModel.DROP:
		return "DROP"
	default:
		return "ACTIVE"
	}
}

func StatusFromString(status string) (commonModel.StudentStatus, error) {
	switch status {
	case "ACTIVE":
		return commonModel.ACTIVE, nil
	case "GRADUATED":
		return commonModel.GRADUATED, nil
	case "DROP":
		return commonModel.DROP, nil
	default:
		return commonModel.ACTIVE, fmt.Errorf("invalid status: %s (must be ACTIVE, GRADUATED, or DROP)", status)
	}
}

func OpenDatabase(database string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := controller.MigrateStudentsToHR(db); err != nil {
		fmt.Printf("Migration Failed: %v\n", err)
		os.Exit(1)
	}
	return db
}
