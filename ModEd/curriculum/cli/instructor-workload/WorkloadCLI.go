package main

import (
	controller "ModEd/curriculum/controller/instructor-workload"
	"flag"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var (
		database string
		// path     string
	)

	flag.StringVar(&database, "database", "", "Path of SQLite Database.")
	// flag.StringVar(&path, "path", "", "Path to CSV or JSON for student registration.")
	flag.Parse()

	connector, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
	// 	fmt.Printf("*** Error: %s does not exist.\n", path)
	// 	return
	// }

	migrationController := controller.NewMigrationController(connector)
	err = migrationController.MigrateToDB()
	if err != nil {
		panic("err: migration failed")
	}

	fmt.Println("Migration successful")
}
