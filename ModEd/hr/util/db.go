package util

import (
	"fmt"
	"os"

	controller "ModEd/hr/controller/Migration"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
