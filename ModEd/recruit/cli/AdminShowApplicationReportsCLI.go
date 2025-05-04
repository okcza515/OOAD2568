package cli

import (
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AdminShowApplicationReportsCLI(service AdminShowApplicationReportsService) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Applicant ID (number): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	applicantID, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Invalid applicant ID. Please enter a valid number.")
		util.WaitForEnter()
		return
	}

	report, err := service.GetApplicationReport(uint(applicantID))
	if err != nil {
		fmt.Println("Error retrieving report:", err)
	} else {
		fmt.Println("===== Applicant Report =====")
		fmt.Println(report)
	}
	util.WaitForEnter()
}
