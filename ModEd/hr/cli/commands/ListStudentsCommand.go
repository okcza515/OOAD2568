package commands

import (
	"flag"
	"fmt"

	"ModEd/hr/controller" // adjust the import paths as needed
	"ModEd/hr/util"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go list student
// no required fields !!
// Run executes the list command, using flags to parse arguments.
func listStudents(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	fs.Parse(args)

	// Using the databasePath flag defined in the main module.
	// Adjust the variable reference or pass the flag value as needed.
	db := util.OpenDatabase(*util.DatabasePath)

	hrFacade := controller.NewHRFacade(db)
	studentInfos, err := hrFacade.GetAllStudents()
	if err != nil {
		return fmt.Errorf("Error listing students: %v\n", err)
	}

	fmt.Println("Human Resource Student Info:")
	for _, s := range studentInfos {
		fmt.Printf("SID: %s | Name: %s %s | Gender: %s | CitizenID: %s | Phone: %s | Status: %s | Email: %s\n",
			s.StudentCode, s.FirstName, s.LastName, s.Gender, s.CitizenID, s.PhoneNumber, util.StatusToString(*s.Status), s.Email)
	}
	return nil
}
