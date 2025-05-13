package main

import (
	"ModEd/common/cli/menu"
	"ModEd/core/cli"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CommonCLI struct {
	db          *gorm.DB
	path        string
	menuManager *cli.CLIMenuStateManager
}

func NewCommonCLI(db *gorm.DB, path string) *CommonCLI {
	menuManager := menu.InitMenuManager(db)

	cli := &CommonCLI{
		db:          db,
		path:        path,
		menuManager: menuManager,
	}
	cli.initializeMenus()
	return cli
}

func (c *CommonCLI) initializeMenus() {
	modelMenu := c.menuManager.GetState("main")
	if modelMenu != nil {
		c.menuManager.SetState(modelMenu)
	}
}

func (c *CommonCLI) Run() {
	fmt.Println("Welcome to the Common CLI!")
	fmt.Println("This CLI provides a simple interface to manage your data.")
	fmt.Println()

	for {
		c.menuManager.Render()
		fmt.Print("\nEnter your choice: ")
		var choice string
		fmt.Scanln(&choice)

		c.menuManager.UserInput = choice
		if err := c.menuManager.HandleUserInput(); err != nil {
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