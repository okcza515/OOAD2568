package commands

import (
	commonController "ModEd/common/controller"
	"ModEd/core"
	"ModEd/hr/controller"
	"ModEd/hr/model"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"

	"errors"
	"flag"
	"fmt"
	"os"

	"gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go import student -path=<path>
// required field : path !!

func importStudents(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("import", flag.ExitOnError)
	filePath := fs.String("path", "", "Path to CSV or JSON for HR student info (only studentid and HR fields).")
	fs.Parse(args)

	err := util.NewValidationChain(fs).
		Required("path").
		Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	if _, err := os.Stat(*filePath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("*** Error: File %s does not exist", *filePath)
	}

	hrMapper, err := core.CreateMapper[model.StudentInfo](*filePath)
	if err != nil {
		return fmt.Errorf("failed to create HR mapper: %v", err)
	}

	hrRecords := hrMapper.Deserialize()
	hrRecordsMap := make(map[string]model.StudentInfo)
	for _, hrRec := range hrRecords {
		if _, exists := hrRecordsMap[hrRec.StudentCode]; exists {
			return fmt.Errorf("duplicate student code found: %s", hrRec.StudentCode)
		}
		hrRecordsMap[hrRec.StudentCode] = *hrRec
	}

	db := util.OpenDatabase(*util.DatabasePath)

	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(db)

		for _, hrRec := range hrRecordsMap {
			commonStudentController := commonController.CreateStudentController(db)
			commonStudent, err := commonStudentController.GetByCode(hrRec.StudentCode)
			if err != nil {
				fmt.Printf("failed to retrieve student %s from common data: %v", hrRec.StudentCode, err)
				continue
			}

			builder := hrModel.NewStudentInfoBuilder()
			req, err := builder.WithStudentCode(hrRec.StudentCode).
				WithStudent(*commonStudent).
				WithGender(hrRec.Gender).
				WithCitizenID(hrRec.CitizenID).
				WithPhoneNumber(hrRec.PhoneNumber).
				Build()

			if err := hrFacade.UpsertStudent(req); err != nil {
				return fmt.Errorf("failed to upsert student %s: %v", req.StudentCode, err)
			}
		}
		fmt.Println("Students imported successfully!")
		return nil
	})
}
