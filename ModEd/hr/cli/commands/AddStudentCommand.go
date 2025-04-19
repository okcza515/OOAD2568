package commands

import (
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"os"
)

func (c *AddStudentCommand) Run(args []string) {
    fs := flag.NewFlagSet("add", flag.ExitOnError)
    studentID := fs.String("id", "", "Student ID")
    firstName := fs.String("fname", "", "First Name")
    lastName := fs.String("lname", "", "Last Name")
    gender := fs.String("gender", "", "Gender")
    citizenID := fs.String("citizenID", "", "Citizen ID")
    phoneNumber := fs.String("phone", "", "Phone Number")
    fs.Parse(args)

    util.ValidateFlags(fs, []string{"id", "fname", "lname"})

    db := util.OpenDatabase(*util.DatabasePath)

	hrFacade := controller.NewHRFacade(db)
	
    newStudent := hrModel.NewStudentInfoBuilder().
        WithStudentCode(*studentID).
        WithFirstName(*firstName).
        WithLastName(*lastName).
        WithGender(*gender).
        WithCitizenID(*citizenID).
        WithPhoneNumber(*phoneNumber).
        Build()

		if err := hrFacade.InsertStudent(newStudent); err != nil {
			fmt.Printf("Failed to add student info: %v\n", err)
			os.Exit(1)
		}

    fmt.Println("Student added successfully!")
}