package commands

import (
    "flag"
    "fmt"
    "os"
    "ModEd/hr/controller"
    hrUtil "ModEd/hr/util"
    

)
func (c *UpdateStudentCommand) Run(args []string) {
    fs := flag.NewFlagSet("update", flag.ExitOnError)
    studentID := fs.String("id", "", "Student ID to update")
    firstName := fs.String("fname", "", "New First Name value")
    lastName := fs.String("lname", "", "New Last Name value")
    gender := fs.String("gender", "", "New Gender value")
    citizenID := fs.String("citizenID", "", "New Citizen ID value")
    phoneNumber := fs.String("phone", "", "New Phone Number value")
    emailStudent := fs.String("email", "", "New Email value")
    fs.Parse(args)

    hrUtil.ValidateFlags(fs, []string{"id"})

    db := hrUtil.OpenDatabase(*hrUtil.DatabasePath)

	hrFacade := controller.NewHRFacade(db)
	studentInfo, err := hrFacade.GetStudentById(*studentID)
    if err != nil {
        fmt.Printf("Error retrieving student with ID %s: %v\n", *studentID, err)
        os.Exit(1)
    }

    if *firstName != "" {
        studentInfo.FirstName = *firstName
    }
    if *lastName != "" {
        studentInfo.LastName = *lastName
    }
    if *gender != "" {
        studentInfo.Gender = *gender
    }
    if *citizenID != "" {
        studentInfo.CitizenID = *citizenID
    }
    if *phoneNumber != "" {
        studentInfo.PhoneNumber = *phoneNumber
    }
    if *emailStudent != "" {
        studentInfo.Email = *emailStudent
    }

	if err := hrFacade.UpdateStudent(studentInfo); err != nil {
        fmt.Printf("Failed to update student info: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Student updated successfully!")
}