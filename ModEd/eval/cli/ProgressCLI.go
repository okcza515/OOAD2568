package cli

import (
	"fmt"
	"time"

	commonModel "ModEd/common/model"
	evalModel "ModEd/eval/model"
)

type Progress struct {
	StudentCode  commonModel.Student
	Title        evalModel.Assignment
	AssignmentId evalModel.Assignment
	Status       evalModel.Assignment
	LastUpdate   time.Time
	TotalSubmit  uint
}

func RunProgressInputCLI() {
	var studentCode, firstName, lastName, title, status string
	var totalSubmit uint

	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)

	fmt.Print("Enter First Name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Enter Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Enter Assignment Title: ")
	fmt.Scanln(&title)

	fmt.Print("Enter Assignment Status: ")
	fmt.Scanln(&status)

	fmt.Print("Enter Total Submit Count: ")
	fmt.Scanln(&totalSubmit)

	progress := Progress{
		StudentCode: commonModel.Student{
			StudentCode: studentCode,
			FirstName:   firstName,
			LastName:    lastName,
		},
		Title:       evalModel.Assignment{Title: title},
		Status:      evalModel.Assignment{Status: status},
		LastUpdate:  time.Now(),
		TotalSubmit: totalSubmit,
	}

	fmt.Println("=== Progress Info ===")
	fmt.Println("Student Code:", progress.StudentCode.StudentCode)
	fmt.Printf("Student: %s %s\n", progress.StudentCode.FirstName, progress.StudentCode.LastName)
	fmt.Println("Title:", progress.Title.Title)
	fmt.Println("Status:", progress.Status.Status)
	fmt.Println("Submit Count:", progress.TotalSubmit)
	fmt.Println("Last Update:", progress.LastUpdate.Format("2006-01-02 15:04:05"))
}
