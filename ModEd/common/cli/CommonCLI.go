package main

import (
	"ModEd/common/cli/menu"
	"ModEd/core/authentication"
	"ModEd/core/cli"
	"flag"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var (
		database string
		path     string
	)
	flag.StringVar(&database, "database", "data/ModEd.bin", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "", "Path to CSV or JSON.")
	flag.Parse()

	args := flag.Args()
	fmt.Printf("args: %v\n", args)

	db := ConnectDB()

	authenticationCLI := authentication.NewAuthenticationCLI()
	authenticationCLI.SetDB(db)
	authenticationCLI.SetAllowedRoles([]string{"admin"})
	authenticationCLI.ExecuteItem(args)

	manager := cli.NewCLIMenuManager()
	commonCLIMenu := menu.NewCommonMenuState(manager)

	manager.SetState(commonCLIMenu)
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
