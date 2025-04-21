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

func importInstructor(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("import", flag.ExitOnError)
	filePath := fs.String("path", "", "Path to CSV or JSON for HR instructor info (only instructorid and HR fields).")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"path"}); err != nil {
		fs.Usage()
		return fmt.Errorf("Validation error: %v\n", err)
	}

	if _, err := os.Stat(*filePath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("*** Error: File %s does not exist.\n", *filePath)
	}

	hrMapper, err := core.CreateMapper[model.InstructorInfo](*filePath)
	if err != nil {
		return fmt.Errorf("Failed to create HR mapper: %v\n", err)
	}

	hrRecords := hrMapper.Deserialize()
	hrRecordsMap := make(map[string]model.InstructorInfo)
	for _, hrRec := range hrRecords {
		idStr := fmt.Sprintf("%d", hrRec.ID)
		if _, exists := hrRecordsMap[idStr]; exists {
			return fmt.Errorf("Duplicate instructor code found: %s\n", idStr)
		}
		hrRecordsMap[idStr] = *hrRec
	}

	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	for _, hrRec := range hrRecordsMap {
		commonInstructorController := commonController.CreateInstructorController(db)
		commonInstructor, err := commonInstructorController.GetByInstructorId(fmt.Sprintf("%d", hrRec.ID)) // Convert uint to string
		if err != nil {
			fmt.Printf("Failed to retrieve instructor %d from common data: %v\n", hrRec.ID, err) // Use %d for uint in the error message
			continue
		}

		newInstructor := hrModel.NewInstructorInfoBuilder().
			WithInstructor(*commonInstructor).
			WithGender(hrRec.Gender).
			WithCitizenID(hrRec.CitizenID).
			WithPhoneNumber(hrRec.PhoneNumber).
			WithSalary(hrRec.Salary).
			WithAcademicPosition(hrRec.AcademicPosition).
			WithDepartmentPosition(hrRec.DepartmentPosition).
			Build()
		fmt.Printf("Importing instructor")

		if err := hrFacade.UpsertInstructor(newInstructor); err != nil {
			return fmt.Errorf("Failed to upsert instructor %d: %v\n", newInstructor.ID, err)
		}
	}
	fmt.Println("Instructors imported successfully!")
	return nil
}
