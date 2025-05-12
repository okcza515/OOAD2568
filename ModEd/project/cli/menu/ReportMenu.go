package menu

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
)

func BuildReportMenu(reportController *controller.ReportController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Reports Management",
		Children: []*utils.MenuItem{
			{
				Title: "View All Reports",
				Action: func(io *core.MenuIO) {
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
				Action: func(io *core.MenuIO) {
					io.Println("Adding New Report...")

					io.Print("Enter Senior Project ID: ")
					projectID, err := io.ReadInputID()
					if err != nil {
						return
					}

					io.Print("Enter Report Type (Idea, Proposal, Progress, Midterm, Final): ")
					reportTypeInput, err := io.ReadInput()
					if err != nil {
						io.Println(fmt.Sprintf("Error reading input: %v", err))
						return
					}

					io.Print("Enter Due Date (YYYY-MM-DD): ")
					dueDate, err := io.ReadInputTime()
					if err != nil {
						io.Println(fmt.Sprintf("Invalid Due Date format: %v", err))
						return
					}

					err = reportController.AddNewReport(projectID, reportTypeInput, dueDate)
					if err != nil {
						io.Println(fmt.Sprintf("Error adding new report: %v", err))
					} else {
						io.Println("Report added successfully!")
					}
				},
			},
			{
				Title: "View Report by ID",
				Action: func(io *core.MenuIO) {
					io.Println("Viewing Report by ID...")
					io.Print("Enter Report ID: ")

					reportID, err := io.ReadInputID()
					if err != nil {
						return
					}

					report, err := reportController.RetrieveByID(reportID)
					if err != nil {
						io.Println(fmt.Sprintf("Error retrieving report: %v", err))
						return
					}
					io.Println(report.ToString())
				},
			},
			{
				Title: "Update Report",
				Action: func(io *core.MenuIO) {
					io.Println("Updating Report...")
					io.Print("Enter Report ID to update: ")

					reportID, err := io.ReadInputID()
					if err != nil {
						return
					}

					io.Print("Enter new Due Date (YYYY-MM-DD): ")
					newDueDate, err := io.ReadInputTime()
					if err != nil {
						return
					}

					err = reportController.UpdateReport(reportID, newDueDate)
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						fmt.Println("Report updated successfully!")
					}

				},
			},
			{
				Title: "Delete Report",
				Action: func(io *core.MenuIO) {
					io.Println("Deleting Report...")
					io.Print("Enter Report ID to delete: ")

					reportID, err := io.ReadInputID()
					if err != nil {
						return
					}

					err = reportController.DeleteByID(reportID)
					if err != nil {
						io.Println(fmt.Sprintf("Error deleting report: %v", err))
					} else {
						io.Println("Report deleted successfully!")
					}
				},
			},
			{
				Title: "Submit Report",
				Action: func(io *core.MenuIO) {
					io.Println("Submitting Report...")
					io.Print("Enter Report ID to submit: ")

					reportID, err := io.ReadInputID()
					if err != nil {
						return
					}

					err = reportController.SubmitReport(reportID)
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
