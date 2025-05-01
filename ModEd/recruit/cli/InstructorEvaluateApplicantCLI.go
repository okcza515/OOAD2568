package cli

import (
	recruitUtil "ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func EvaluateApplicant(instructorEvaluateApplicantService InstructorEvaluateApplicantService, ApplicantReportService ApplicantReportService, instructorID uint) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Application Report ID: ")
	scanner.Scan()
	reportID := scanner.Text()
	convReportID, err := strconv.ParseUint(reportID, 10, 32)
	if err != nil {
		fmt.Println("Invalid Application Report ID. Please enter a valid number.")
		return
	}
	applicationReportID := uint(convReportID)

	hasPermission, err := instructorEvaluateApplicantService.HasPermissionToEvaluate(instructorID, applicationReportID)
	if err != nil {
		fmt.Println("Error checking permission:", err)
		recruitUtil.WaitForEnter()
		return
	}
	if !hasPermission {
		fmt.Println("You do not have permission to evaluate this application.")
		recruitUtil.WaitForEnter()
		return
	}

	report, err := ApplicantReportService.GetFullApplicationReportByApplicationID(applicationReportID)
	if err != nil {
		fmt.Println("Failed to fetch application report:", err)
		recruitUtil.WaitForEnter()
		return
	}

	err = instructorEvaluateApplicantService.EvaluateApplicant(applicationReportID, report.ApplicationRound.RoundName)
	if err != nil {
		fmt.Println("Error updating interview score:", err)
	} else {
		fmt.Println("Score updated successfully!")
	}
	recruitUtil.WaitForEnter()
}
