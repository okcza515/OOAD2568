package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetUserInputUint(msg string) uint {
	var input uint
	fmt.Print(msg)
	fmt.Scanln(&input)
	return input
}

func GetInputDatabasePath(defaultDBPath string) string {
	database := defaultDBPath

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
