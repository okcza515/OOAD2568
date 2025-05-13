package main

import (
	"ModEd/common/cli/menu"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CommonCLI struct {
	db   *gorm.DB
	path string
}

func NewCommonCLI(db *gorm.DB, path string) *CommonCLI {
	return &CommonCLI{
		db:   db,
		path: path,
	}
}

func (c *CommonCLI) Run() {
	fmt.Println("Welcome to the Common CLI!")
	fmt.Println("This CLI provides a simple interface to manage your data.")
	fmt.Println()

	handlerContext := menu.NewCommonHandlerContext(c.db)

	for {
		handlerContext.ShowMenu()
		fmt.Print("\nEnter your choice: ")
		var choice string
		fmt.Scanln(&choice)

		if choice == "exit" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		if err := handlerContext.HandleInput(choice); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Println()
	}
}

func initDB() *gorm.DB {
	connector, err := gorm.Open(sqlite.Open("data/ModEd.bin"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return connector
}

func main() {
	db := initDB()
	cli := NewCommonCLI(db, "data")
	cli.Run()
}