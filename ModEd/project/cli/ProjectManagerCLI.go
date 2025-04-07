package main

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
	"time"
)

func main() {
	db := utils.OpenDatabase("project.db")

	reportController := controller.NewReportController(db)
	progressController := controller.NewProgressController(db)

	utils.PrintTitle("Senior Project CLI")

	builder := utils.NewMenuBuilder(&utils.MenuItem{
		Title: "Main Menu",
		Children: []*utils.MenuItem{
			{
				Title: "Senior Project Setup",
				Children: []*utils.MenuItem{
					{
						Title: "Define Assessment Criteria",
						Action: func(io *utils.MenuIO) {
							io.Println("Defining Assessment Criteria...")
							// Add logic to define criteria เพิ่มแล้วลบด้วย
						},
					},
					{
						Title: "Create Senior Project",
						Action: func(io *utils.MenuIO) {
							io.Println("Creating Senior Project...")
							// Add logic to create senior project เพิ่มแล้วลบด้วย
						},
					},
					{
						Title: "Assign Groups, Advisors, and Committees",
						Action: func(io *utils.MenuIO) {
							io.Println("Assigning Groups, Advisors, and Committees...")
							// Add logic to assign groups, advisors, and committees เพิ่มแล้วลบด้วย
						},
					},
				},
			},
			{
				Title: "Project Execution and Monitoring",
				Children: []*utils.MenuItem{
					{
						Title: "View Schedule",
						Action: func(io *utils.MenuIO) {
							io.Println("Viewing Schedule...")

							formattedReports, err := reportController.GetFormattedReportList()
							if err != nil {
								io.Println(fmt.Sprintf("Error retrieving reports: %v", err))
								return
							}
							if len(formattedReports) == 0 {
								io.Println("No reports found.")
								return
							}
							io.Println("Schedule (Based on Due Dates):")
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
					{
						Title: "Track Progress",
						Children: []*utils.MenuItem{
							{
								Title: "View All Progress",
								Action: func(io *utils.MenuIO) {
									io.Println("Viewing All Progress...")

									formattedList, err := progressController.GetFormattedProgressList()
									if err != nil {
										io.Println(fmt.Sprintf("Error retrieving progress: %v", err))
										return
									}

									if len(formattedList) == 0 {
										io.Println("No progress found.")
										return
									}

									io.Println("Progress List:")
									for _, progress := range formattedList {
										io.Println(progress)
									}
								},
							},
							{
								Title: "Mark Progress as Completed",
								Action: func(io *utils.MenuIO) {
									io.Println("Marking Progress as Completed...")
									io.Print("Enter Progress ID: ")

									input, err := io.ReadInput()
									if err != nil {
										io.Println(fmt.Sprintf("Error reading input: %v", err))
										return
									}

									progressID, err := strconv.ParseUint(input, 10, 32)
									if err != nil {
										io.Println(fmt.Sprintf("Invalid Progress ID: %v", err))
										return
									}

									err = progressController.MarkAsCompleted(uint(progressID))
									if err != nil {
										io.Println(fmt.Sprintf("Error marking progress as completed: %v", err))
									} else {
										io.Println("Progress marked as completed successfully!")
									}
								},
							},
						},
					},
				},
			},
			{
				Title: "Evaluation & Assessment",
				Children: []*utils.MenuItem{
					{
						Title: "Evaluate Presentation",
						Action: func(io *utils.MenuIO) {
							io.Println("Evaluating Presentation...")
							io.Print("Enter Evaluation ID: ")

							input, err := io.ReadInput()
							if err != nil {
								io.Println(fmt.Sprintf("Error reading input: %v", err))
								return
							}

							evaluationID, err := strconv.ParseUint(input, 10, 32)
							if err != nil {
								io.Println(fmt.Sprintf("Invalid Evaluation ID: %v", err))
								return
							}

							// Add logic to evaluate presentation เพิ่มแล้วลบด้วย
							io.Println(fmt.Sprintf("Presentation with ID %d evaluated successfully!", evaluationID))
						},
					},
					{
						Title: "Compile Final Scores",
						Action: func(io *utils.MenuIO) {
							io.Println("Compiling Final Scores...")
							// Add logic to compile final scores เพิ่มแล้วลบด้วย
						},
					},
				},
			},
		},
	}, nil, nil)

	builder.Show()
}
