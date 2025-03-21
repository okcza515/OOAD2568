package main

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/gorm"
)

// TODO: folder migration might be remove, waiting for WIT to do something with it
func main() {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "test.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}
	migrationController := controller.NewMigrationController(db)
	if err = migrationController.MigrateToDB(); err != nil {
		panic(err)
	}
	fmt.Println("Migration completed")
}
