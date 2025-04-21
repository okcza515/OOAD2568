//MEP-1009 Student Internship
package Internship

import (
	controller "ModEd/curriculum/controller"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InternshipReview(Review *controller.ReviewController) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter StudentCode: ")
	scanner.Scan()
	studentCode := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Supervisor Score: ")
	scanner.Scan()
	SupervisorScoreInput := strings.TrimSpace(scanner.Text())
	SupervisorScore, err := strconv.Atoi(SupervisorScoreInput)
	if err != nil {
		fmt.Printf("Error: Invalid score input. Please enter a valid number.\n")
		return
	}

	fmt.Print("Enter Mentor Score: ")
	scanner.Scan()
	MentorScoreInput := strings.TrimSpace(scanner.Text())
	MentorScore, err := strconv.Atoi(MentorScoreInput)
	if err != nil {
		fmt.Printf("Error: Invalid score input. Please enter a valid number.\n")
		return
	}
	Review.UpdateReviewScore(studentCode, SupervisorScore, MentorScore)

}
