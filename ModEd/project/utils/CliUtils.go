package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase(database string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	return db
}

func PrintTitle(title string) {
	fmt.Println("==========================================================")
	fmt.Printf("                   🎓 %s                  \n", title)
	fmt.Println("==========================================================")
	fmt.Println("Welcome to the Senior Project Management (MEP-1005) System!")
	fmt.Println("Use the menu below to navigate through the system.")
	fmt.Println()
}
