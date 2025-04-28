package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

func (c *AddStudentCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	firstName := fs.String("fname", "", "First Name")
	lastName := fs.String("lname", "", "Last Name")
	gender := fs.String("gender", "", "Gender")
	citizenID := fs.String("citizenID", "", "Citizen ID")
	phoneNumber := fs.String("phone", "", "Phone Number")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "fname", "lname"}); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	db := util.OpenDatabase(*util.DatabasePath)

	// Create a TransactionManager instance.
	tm := &util.TransactionManager{DB: db}

	err := tm.Execute(func(tx *gorm.DB) error {
		return controller.AddStudent(
			tx,
			*studentID,
			*firstName,
			*lastName,
			*gender,
			*citizenID,
			*phoneNumber,
		)
	})

	if err != nil {
		return fmt.Errorf("transaction failed: %v", err)
	}

	fmt.Println("Student added and HR info updated successfully!")

	return nil
}
