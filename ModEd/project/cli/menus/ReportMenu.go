package menus

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
	"time"
)

func BuildReportMenu(reportController *controller.ReportController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Reports Management",
		Children: []*utils.MenuItem{
			{
				Title: "View Reports",
				Action: func(io *utils.MenuIO) {
					io.Println("Viewing Report...")

					formattedReports, err := reportController.GetFormattedReportList()
					if err != nil {
						io.Println(fmt.Sprintf("Error retrieving reports: %v", err))
						return
					}
					if len(formattedReports) == 0 {
						io.Println("No reports found.")
						return
					}
					io.Println("Report (Based on Due Dates):")
					for _, report := range formattedReports {
						io.Println(report)
					}
				},
			},
			{
				Title: "Add New Report",
				Action: func(io *utils.MenuIO) {
					io.Println("Adding New Report...")

					io.Print("Enter Senior Project ID: ")
					projectIDInput, err := io.ReadInput()
					if err != nil {
						io.Println(fmt.Sprintf("Error reading input: %v", err))
						return
					}
					projectID, err := strconv.Atoi(projectIDInput)
					if err != nil {
						io.Println(fmt.Sprintf("Invalid Senior Project ID: %v", err))
						return
					}

					io.Print("Enter Report Type (Idea, Proposal, Progress, Midterm, Final): ")
					reportTypeInput, err := io.ReadInput()
					if err != nil {
						io.Println(fmt.Sprintf("Error reading input: %v", err))
						return
					}

					io.Print("Enter Due Date (YYYY-MM-DD): ")
					dueDateInput, err := io.ReadInput()
					if err != nil {
						io.Println(fmt.Sprintf("Error reading input: %v", err))
						return
					}
					dueDate, err := time.Parse("2006-01-02", dueDateInput)
					if err != nil {
						io.Println(fmt.Sprintf("Invalid Due Date format: %v", err))
						return
					}

					err = reportController.AddNewReport(uint(projectID), reportTypeInput, dueDate)
					if err != nil {
						io.Println(fmt.Sprintf("Error adding new report: %v", err))
					} else {
						io.Println("Report added successfully!")
					}
				},
			},
			{
				Title: "Submit Report",
				Action: func(io *utils.MenuIO) {
					io.Println("Submitting Report...")
					io.Print("Enter Report ID to submit: ")

					input, err := io.ReadInput()
					if err != nil {
						io.Println(fmt.Sprintf("Error reading input: %v", err))
						return
					}

					reportID, err := strconv.ParseUint(input, 10, 32)
					if err != nil {
						io.Println(fmt.Sprintf("Invalid Report ID: %v", err))
						return
					}

					err = reportController.SubmitReport(uint(reportID))
					if err != nil {
						io.Println(fmt.Sprintf("Error submitting report: %v", err))
					} else {
						io.Println("Report submitted successfully!")
					}
				},
			},
		},
	}
}
