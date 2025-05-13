package utils

import (
	"ModEd/core"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type CSVImporter struct {
	io *core.MenuIO
}

func NewCSVImporter(io *core.MenuIO) *CSVImporter {
	return &CSVImporter{io: io}
}

func (ci *CSVImporter) ReadFile() ([][]string, error) {
	ci.io.Println("Current directory:" + ci.getCurrentDir())
	filePath := ci.promptFilePath()
	if filePath == "" {
		return nil, nil // Cancelled
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("error resolving path: %v", err)
	}

	if !ci.fileExists(absPath) {
		return nil, fmt.Errorf("file does not exist")
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	return csv.NewReader(file).ReadAll()
}

func (ci *CSVImporter) promptFilePath() string {
	ci.io.Print("Enter CSV file path (-1 to cancel): ")
	filePath, err := ci.io.ReadInput()
	if err != nil || filePath == "-1" {
		ci.io.Println("Cancelled.")
		return ""
	}
	return filePath
}

func (ci *CSVImporter) getCurrentDir() string {
	cwd, _ := os.Getwd()
	return cwd
}

func (ci *CSVImporter) fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
