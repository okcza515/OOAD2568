package testing

import (
	"ModEd/project/utils"
	"fmt"
	"os"
	"testing"
)

func TestImportCsv(t *testing.T) {
	// Create a temporary CSV file
	tmpFile, err := os.CreateTemp("", "testdata-*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	content := `Name,Department Alice,Computer Science Bob,Information Technology
`
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	if records := utils.ImportCsv(tmpFile.Name()); records != nil {
		fmt.Println(records)
	}
}
