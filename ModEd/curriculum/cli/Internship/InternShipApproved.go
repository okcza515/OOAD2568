package Internship

import (
	controller "ModEd/curriculum/controller/Internship"
	"bufio"
	"fmt"
	"os"
	"strings"

	model "ModEd/curriculum/model/Internship"
)

func InternShipApproved(Approved *controller.ApprovedController) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter StudentCode: ")
	scanner.Scan()
	studentCode := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Advisor Approval Status (APPROVED/REJECT): ")
	scanner.Scan()
	advisorStatus := strings.ToUpper(strings.TrimSpace(scanner.Text()))

	fmt.Print("Enter Company Approval Status (APPROVED/REJECT): ")
	scanner.Scan()
	companyStatus := strings.ToUpper(strings.TrimSpace(scanner.Text()))

	err := Approved.UpdateApprovalStatuses(studentCode, model.ApprovedStatus(advisorStatus), model.ApprovedStatus(companyStatus))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Approval statuses updated successfully!")
		if advisorStatus == string(model.APPROVED) && companyStatus == string(model.APPROVED) {
			fmt.Println("Intern status updated to ACTIVE.")
		}
	}
}
