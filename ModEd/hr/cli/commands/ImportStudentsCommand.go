package commands

import (
	commonController "ModEd/common/controller"
	"ModEd/hr/controller"
	"ModEd/hr/model"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"

	"errors"
	"flag"
	"fmt"
	"os"
)

// usage : go run hr/cli/HumanResourceCLI.go import -path=<path>
// required field : path !!

func (c *ImportStudentsCommand) Run(args []string) {
	fs := flag.NewFlagSet("import", flag.ExitOnError)
	filePath := fs.String("path", "", "Path to CSV or JSON for HR student info (only studentid and HR fields).")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"filePath"}); err != nil {
        fmt.Printf("Validation error: %v\n", err)
        fs.Usage()
        os.Exit(1)
    }


	if _, err := os.Stat(*filePath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("*** Error: File %s does not exist.\n", *filePath)
		os.Exit(1)
	}
	hrMapper, err := util.CreateMapper[model.StudentInfo](*filePath)
	if err != nil {
		fmt.Printf("Failed to create HR mapper: %v\n", err)
		os.Exit(1)
	}

	hrRecords := hrMapper.Map()

	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	for _, hrRec := range hrRecords {
		commonStudentController := commonController.CreateStudentController(db)
		commonStudent, err := commonStudentController.GetByStudentId(hrRec.StudentCode)
		if err != nil {
			fmt.Printf("Failed to retrieve student %s from common data: %v\n", hrRec.StudentCode, err)
			continue
		}

		newStudent := hrModel.NewStudentInfoBuilder().
			WithStudent(*commonStudent).
			WithGender(hrRec.Gender).
			Build()

		if err := hrFacade.UpsertStudent(newStudent); err != nil {
			fmt.Printf("Failed to upsert student %s: %v\n", newStudent.StudentCode, err)
			continue
		}
	}
	fmt.Println("Students imported successfully!")
}
