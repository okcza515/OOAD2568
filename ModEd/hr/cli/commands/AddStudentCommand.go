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
	email := fs.String("email", "", "Email")
	gender := fs.String("gender", "", "Gender")
	citizenID := fs.String("citizenID", "", "Citizen ID")
	phoneNumber := fs.String("phone", "", "Phone Number")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "fname", "lname", "email", "gender", "citizenID", "phoneNumber"}); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	idLengthValidator := util.NewLengthFlagValidator(fs, "id", 11)
	regexValidator, err := util.NewRegexFlagValidator(fs, "id", "^[0-9]{11}$")
	if err != nil {
		fs.Usage()
		return fmt.Errorf("failed to create regex validator for id: %v", err)
	}

	idLengthValidator.SetNext(regexValidator)
	if err := idLengthValidator.Validate(); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error for id: %v", err)
	}

	emailRegexValidator, err := util.NewRegexFlagValidator(fs, "email", "^[a-zA-Z0-9._%%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
	if err != nil {
		fs.Usage()
		return fmt.Errorf("failed to create regex validator for email: %v", err)
	}
	if err := emailRegexValidator.Validate(); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error for email: %v", err)
	}

	db := util.OpenDatabase(*util.DatabasePath)

	// Create a TransactionManager instance.
	tm := &util.TransactionManager{DB: db}

	err = tm.Execute(func(tx *gorm.DB) error {
		return controller.AddStudent(
			tx,
			*studentID,
			*firstName,
			*lastName,
			*email,
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
