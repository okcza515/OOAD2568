package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

// usage : go run hr/cli/HumanResourceCLI.go export -field="value"
// required field : path, format !!
// format : csv or json !!

func (c *ExportStudentsCommand) Run(args []string) {
	fs := flag.NewFlagSet("export", flag.ExitOnError)
	filePath := fs.String("path", "", "File path to export data")
	format := fs.String("format", "", "Export format (csv or json)")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"path", "format"}); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		fs.Usage()
		os.Exit(1)
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
			fmt.Println("Invalid format. Supported formats are 'csv' and 'json'.")
			os.Exit(1)
		}
	}

	db := util.OpenDatabase(*util.DatabasePath)

	hrFacade := controller.NewHRFacade(db)
	studentInfos, err := hrFacade.GetAllStudents()
	if err != nil {
		fmt.Printf("Error fetching students: %v\n", err)
		os.Exit(1)
	}

	var exportErr error
	switch *format {
	case "csv":
		// Use gocsv directly for CSV serialization
		file, err := os.OpenFile(*filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		if err := gocsv.MarshalFile(&studentInfos, file); err != nil {
			exportErr = err
		}
	case "json":
		// Use encoding/json directly for JSON serialization
		file, err := os.OpenFile(*filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(studentInfos); err != nil {
			exportErr = err
		}
	default:
		fmt.Println("Invalid format. Supported formats are 'csv' and 'json'.")
		os.Exit(1)
	}

	if exportErr != nil {
		fmt.Printf("Error exporting data: %v\n", exportErr)
		os.Exit(1)
	}

	fmt.Println("Student info exported successfully!")
}
