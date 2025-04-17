package utils

import "fmt"

func GetUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func GetInputDatabasePath(module string, defaultDBPath string) string {
	database := defaultDBPath

	fmt.Printf("Welcome to ModEd %s CLI!\n", module)

	fmt.Println("Please enter the path of SQLite Database (press Enter to use default):")
	fmt.Printf("Default: %s\n", defaultDBPath)
	fmt.Scanln(&database)

	if database == "" {
		database = defaultDBPath
	}

	fmt.Println("Using database path:", database)

	return database
}

func GetInputDataPath(module string, defaultDataPath string) string {
	dataPath := defaultDataPath

	fmt.Printf("Please enter the path of CSV or JSON for %s (press Enter to use default):\n", module)
	fmt.Printf("Default: %s\n", defaultDataPath)
	fmt.Scanln(&dataPath)

	if dataPath == "" {
		dataPath = defaultDataPath
	}

	fmt.Println("Using data path:", dataPath)

	return dataPath
}
