package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go export -field="value"
// required field : path, format !!
// format : csv or json !!

func (c *ExportStudentsCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("export", flag.ExitOnError)
	filePath := fs.String("path", "", "File path to export data")
	format := fs.String("format", "", "Export format (csv or json)")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"path", "format"}); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	fileInfo, err := os.Stat(*filePath)
	if err == nil && fileInfo.IsDir() {
		// Append the default file name based on the format
		switch *format {
		case "csv":
			*filePath = fmt.Sprintf("%s/studentinfo.csv", *filePath)
		case "json":
			*filePath = fmt.Sprintf("%s/studentinfo.json", *filePath)
		default:
			return fmt.Errorf("invalid format. Supported formats are 'csv' and 'json'")
		}
	}

	db := util.OpenDatabase(*util.DatabasePath)

	hrFacade := controller.NewHRFacade(db)
	studentInfos, err := hrFacade.GetAllStudents()
	if err != nil {
		return fmt.Errorf("error fetching students: %v", err)
	}

	var exportErr error
	switch *format {
	case "csv":
		// Use gocsv directly for CSV serialization
		file, err := os.OpenFile(*filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}
		defer file.Close()

		if err := gocsv.MarshalFile(&studentInfos, file); err != nil {
			exportErr = err
		}
	case "json":
		// Use encoding/json directly for JSON serialization
		file, err := os.OpenFile(*filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(studentInfos); err != nil {
			exportErr = err
		}
	default:
		return fmt.Errorf("invalid format. Supported formats are 'csv' and 'json'")
	}

	if exportErr != nil {
		return fmt.Errorf("error exporting data: %v", exportErr)
	}

	fmt.Println("Student info exported successfully!")
	return nil
}
