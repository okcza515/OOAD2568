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

func InternshipReport(Report *controller.ReportController) {

	fmt.Print("Enter StudentCode: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	studentCode := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Score: ")
	scanner.Scan()
	scoreInput := strings.TrimSpace(scanner.Text())
	score, err := strconv.Atoi(scoreInput)
	if err != nil {
		fmt.Printf("Error: Invalid score input. Please enter a valid number.\n")
		return
	}
	Report.UpdateReportScore(studentCode, score)
}
