// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/asset/util"
	"ModEd/recruit/controller"
	"bufio"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func InstructorCLI(instructorCtrl *controller.InstructorController) {

	util.ClearScreen()

	var instructorID uint
	fmt.Print("Enter Instructor ID: ")
	fmt.Scan(&instructorID)

	for {
		util.ClearScreen()
		fmt.Println("\n==== Instructor Menu ====")
		fmt.Println("1. View Interview Details")
		fmt.Println("2. Evaluate an Applicant")
		fmt.Println("3. back")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ViewInterviewDetails(instructorCtrl, instructorID)
		case 2:
			EvaluateApplicant(instructorCtrl)
		case 3:
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func ViewInterviewDetails(instructorCtrl *controller.InstructorController, instructorID uint) {
	interviews, err := instructorCtrl.GetInterviewsByInstructor(instructorID)
	if err != nil {
		fmt.Println("Error retrieving interviews:", err)
		return
	}

	fmt.Println("\n==== Interview Schedule ====")
	for _, interview := range interviews {
		fmt.Printf("ID: %d | Applicant ID: %d | Date: %s | Score: ",
			interview.ID, interview.ApplicantID, interview.ScheduledAppointment)

		// Check if InterviewScore is nil before dereferencing
		if interview.InterviewScore != nil {
			fmt.Println(*interview.InterviewScore) // Dereference pointer
		} else {
			fmt.Println("Not Assigned") // Handle nil case
		}
	}
}

func EvaluateApplicant(instructorCtrl *controller.InstructorController) {
	var interviewID uuid.UUID
	var score float64

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Interview ID: ")
	scanner.Scan()
	interviewID, _ = uuid.Parse(scanner.Text())

	fmt.Print("Enter Interview Score: ")
	fmt.Scan(&score)

	err := instructorCtrl.EvaluateApplicant(interviewID, score)
	if err != nil {
		fmt.Println("Error updating interview score:", err)
	} else {
		fmt.Println("Score updated successfully!")
	}
}
