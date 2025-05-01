package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

type ListStudentCommand struct{}

func (cmd *ListStudentCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("list-student", flag.ExitOnError)
	fs.Parse(args)

	studentInfos, err := controller.GetAllStudents(tx)
	if err != nil {
		return fmt.Errorf("error listing students: %v", err)
	}

	fmt.Println("Human Resource Student Info:")
	for _, s := range studentInfos {
		fmt.Printf("SID: %s | Name: %s %s | Gender: %s | CitizenID: %s | Phone: %s | Status: %s | Email: %s\n",
			s.StudentCode, s.FirstName, s.LastName, s.Gender, s.CitizenID, s.PhoneNumber, util.StatusToString(*s.Status), s.Email)
	}
	return nil
}
