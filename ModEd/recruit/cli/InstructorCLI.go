package cli

import (
	"ModEd/recruit/controller"
	"bufio"
	"fmt"
	"os"
)

func InstructorCLI(instructorCtrl *controller.InstructorController) {
	var instructorID uint
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&instructorID)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==== Instructor Menu ====")
		fmt.Println("1. View Interview Details")
		fmt.Println("2. Evaluate an Applicant")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		var choice int
		scanner.Scan()
		_, err := fmt.Sscan(scanner.Text(), &choice)
		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		switch choice {
		case 1:
			ViewInterviewDetails(instructorCtrl, instructorID)
		case 2:
			EvaluateApplicant(instructorCtrl)
		case 3:
			fmt.Println("Exiting...")
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
	var interviewID uint
	var score float64

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Interview ID: ")
	scanner.Scan()
	_, err := fmt.Sscan(scanner.Text(), &interviewID)
	if err != nil {
		fmt.Println("Invalid Interview ID.")
		return
	}

	fmt.Print("Enter Interview Score: ")
	scanner.Scan()
	_, err = fmt.Sscan(scanner.Text(), &score)
	if err != nil {
		fmt.Println("Invalid score.")
		return
	}

	err = instructorCtrl.EvaluateApplicant(interviewID, score)
	if err != nil {
		fmt.Println("Error updating interview score:", err)
	} else {
		fmt.Println("Score updated successfully!")
	}
}
