package cli

import (
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func AdminScheduleInterviewCLI(service AdminScheduleInterviewService) {
	var instructorID string
	var con_int_instrucID uint
	var applicationReportID uint

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Instructor ID: ")
	scanner.Scan()
	instructorID = scanner.Text()
	convInstructorID, err := strconv.ParseUint(instructorID, 10, 32)
	if err != nil {
		fmt.Println("Invalid Instructor ID. Please enter a valid number.")
		return
	}
	con_int_instrucID = uint(convInstructorID)

	fmt.Print("Enter Application Report ID: ")
	scanner.Scan()
	reportID := scanner.Text()
	convReportID, err := strconv.ParseUint(reportID, 10, 32)
	if err != nil {
		fmt.Println("Invalid Application Report ID. Please enter a valid number.")
		return
	}
	applicationReportID = uint(convReportID)

	fmt.Print("Enter Scheduled Appointment (YYYY-MM-DD HH:MM:SS): ")
	scanner.Scan()
	scheduledTime := scanner.Text()

	err = service.ScheduleInterview(con_int_instrucID, applicationReportID, scheduledTime)
	if err != nil {
		fmt.Println("Error scheduling interview:", err)
	} else {
		fmt.Println("Interview scheduled successfully!")
	}

	util.WaitForEnter()
}
