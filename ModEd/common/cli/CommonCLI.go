package cli

import (
	"ModEd/common/cli/menu"
	"fmt"

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

	mainMenu := menu.NewCommonCLIMenu(c.db, c.path)
	mainMenu.Run()
}

func main() {
	db := initDB()
	cli := NewCommonCLI(db, "data")
	cli.Run()
}

func ConnectDB() *gorm.DB {
	connector, err := gorm.Open(sqlite.Open("data/ModEd.bin"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return connector
}

func confirmAction(prompt string) bool {
	var response string
	fmt.Print(prompt)
	fmt.Scan(&response)
	return response == "y" || response == "Y"
}

func initDB() *gorm.DB {
	connector, err := gorm.Open(sqlite.Open("data/ModEd.bin"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return connector
}
