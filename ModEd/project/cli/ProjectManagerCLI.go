package main

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	db := utils.OpenDatabase("project.db")

	reportController := controller.NewReportController(db)
	progressController := controller.NewProgressController(db)
	assessmentController := controller.NewAssessmentController(db)
	assessmentCriteriaController := controller.NewAssessmentCriteriaController(db)
	assessmentCriteriaLinkController := controller.NewAssessmentCriteriaLinkController(db)

	utils.PrintTitle("Senior Project CLI")

	builder := utils.NewMenuBuilder(&utils.MenuItem{
		Title: "Main Menu",
		Children: []*utils.MenuItem{
			{
				Title: "Senior Project Setup",
				Children: []*utils.MenuItem{
					{
						Title: "Assessment Criteria Manager",
						Children: []*utils.MenuItem{
							{
								Title: "Define Assessment Criteria",
								Action: func(io *utils.MenuIO) {
									io.Println("Defining Assessment Criteria...")
									io.Print("Enter Criteria Name (-1 to cancel): ")
									criteriaName, err := io.ReadInput()
									if err != nil || criteriaName == "-1" {
										io.Println("Cancelled.")
										return
									}
									err = assessmentCriteriaController.InsertAssessmentCriteria(criteriaName)
									if err != nil {
										io.Println(fmt.Sprintf("Error adding new criteria: %v", err))
									} else {
										io.Println("Criteria added successfully!")
									}
								},
							},
							{
								Title: "List All Assessment Criteria",
								Action: func(io *utils.MenuIO) {
									io.Println("Listing All Assessment Criteria...")
									criteriaList, err := assessmentCriteriaController.ListAllAssessmentCriterias()
									if err != nil {
										io.Println(fmt.Sprintf("Error listing criteria: %v", err))
										return
									}
									if len(criteriaList) == 0 {
										io.Println("No assessment criteria found.")
										return
									}
									for _, c := range criteriaList {
										io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
									}
								},
							},
							{
								Title: "Update Assessment Criteria Name",
								Action: func(io *utils.MenuIO) {
									io.Println("All Criteria:")
									criteriaList, err := assessmentCriteriaController.ListAllAssessmentCriterias()
									if err != nil {
										io.Println(fmt.Sprintf("Error retrieving criteria list: %v", err))
										return
									}
									for _, c := range criteriaList {
										io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
									}

									io.Print("Enter Criteria ID to update (-1 to cancel): ")
									input, _ := io.ReadInput()
									if input == "-1" {
										io.Println("Cancelled.")
										return
									}
									id, err := strconv.Atoi(input)
									if err != nil {
										io.Println("Invalid ID.")
										return
									}

									criteria, err := assessmentCriteriaController.RetrieveAssessmentCriteria(uint(id))
									if err != nil || criteria == nil {
										io.Println("Criteria not found.")
										return
									}

									io.Print("Enter New Criteria Name (-1 to cancel): ")
									newName, _ := io.ReadInput()
									if newName == "-1" {
										io.Println("Cancelled.")
										return
									}

									criteria.CriteriaName = newName
									err = assessmentCriteriaController.UpdateAssessmentCriteria(uint(id), criteria)
									if err != nil {
										io.Println(fmt.Sprintf("Error updating criteria: %v", err))
									} else {
										io.Println("Criteria updated successfully!")
									}
								},
							},
							{
								Title: "Delete Assessment Criteria",
								Action: func(io *utils.MenuIO) {
									io.Println("All Criteria:")
									criteriaList, err := assessmentCriteriaController.ListAllAssessmentCriterias()
									if err != nil {
										io.Println(fmt.Sprintf("Error listing criteria: %v", err))
										return
									}
									for _, c := range criteriaList {
										io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
									}

									io.Print("Enter Criteria ID to delete (-1 to cancel): ")
									input, _ := io.ReadInput()
									if input == "-1" {
										io.Println("Cancelled.")
										return
									}
									id, err := strconv.Atoi(input)
									if err != nil {
										io.Println("Invalid ID.")
										return
									}

									err = assessmentCriteriaController.DeleteAssessmentCriteria(uint(id))
									if err != nil {
										io.Println(fmt.Sprintf("Error deleting criteria: %v", err))
									} else {
										io.Println("Criteria deleted successfully!")
									}
								},
							},
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

					{
						Title: "Assessment Manager",
						Action: func(io *utils.MenuIO) {
							io.Println("Assessment Manager, choose what you want to do...")
						},
						Children: []*utils.MenuItem{
							{
								Title: "List All Criteria Linked to Assessment",
								Action: func(io *utils.MenuIO) {
									io.Print("Enter Senior Project ID (-1 to cancel): ")
									input, err := io.ReadInput()
									if err != nil {
										io.Println(fmt.Sprintf("Error reading input: %v", err))
										return
									}
									if input == "-1" {
										return
									}

									seniorProjectID, err := strconv.ParseUint(input, 10, 32)
									if err != nil {
										io.Println(fmt.Sprintf("Invalid project ID: %v", err))
										return
									}

									links, err := assessmentCriteriaLinkController.ListProjectAssessmentCriteriaLinks(uint(seniorProjectID))
									if err != nil {
										io.Println(fmt.Sprintf("Error retrieving links: %v", err))
										return
									}
									if len(links) == 0 {
										io.Println("No criteria linked to this assessment.")
										return
									}

									io.Println("Linked Criteria:")
									for _, link := range links {
										criteria, err := assessmentCriteriaController.RetrieveAssessmentCriteria(link.AssessmentCriteriaId)
										if err != nil {
											log.Printf("Error retrieving criteria (ID %d): %v", link.AssessmentCriteriaId, err)
											continue
										}
										if criteria == nil {
											log.Printf("Criteria ID %d not found", link.AssessmentCriteriaId)
											continue
										}
										io.Println(fmt.Sprintf("Criteria ID: %v, Name: %v", criteria.ID, criteria.CriteriaName))
									}
								},
							},

							{
								Title: "Link Criteria to Assessment",
								Action: func(io *utils.MenuIO) {
									io.Println("Linking Criteria to Assessment...")
									io.Print("Enter Senior Project ID (-1 to cancel): ")
									input, err := io.ReadInput()
									if err != nil {
										io.Println(fmt.Sprintf("Error reading input: %v", err))
										return
									}
									if input == "-1" {
										return
									}

									seniorProjectID, err := strconv.ParseUint(input, 10, 32)
									if err != nil {
										io.Println(fmt.Sprintf("Invalid project ID: %v", err))
										return
									}

									links, err := assessmentCriteriaLinkController.ListProjectAssessmentCriteriaLinks(uint(seniorProjectID))
									if err != nil {
										io.Println(fmt.Sprintf("Error retrieving linked criteria: %v", err))
										return
									}

									io.Println("Currently linked criteria:")
									for _, link := range links {
										criteria, err := assessmentCriteriaController.RetrieveAssessmentCriteria(link.AssessmentCriteriaId)
										if err != nil || criteria == nil {
											log.Printf("Error retrieving criteria (ID %d): %v", link.AssessmentCriteriaId, err)
											continue
										}
										io.Println(fmt.Sprintf("ID: %v, Name: %v", criteria.ID, criteria.CriteriaName))
									}

									// List all available criteria
									io.Println("\nAvailable Assessment Criteria:")
									allCriteria, err := assessmentCriteriaController.ListAllAssessmentCriterias()
									if err != nil {
										io.Println(fmt.Sprintf("Error listing all criteria: %v", err))
										return
									}
									for _, c := range allCriteria {
										io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
									}

									io.Print("Enter Criteria ID to link (-1 to cancel): ")
									criteriaInput, err := io.ReadInput()
									if err != nil {
										io.Println(fmt.Sprintf("Error reading input: %v", err))
										return
									}
									if criteriaInput == "-1" {
										return
									}

									criteriaID, err := strconv.ParseUint(criteriaInput, 10, 32)
									if err != nil {
										io.Println(fmt.Sprintf("Invalid criteria ID: %v", err))
										return
									}

									// Check if it exists
									criteria, err := assessmentCriteriaController.RetrieveAssessmentCriteria(uint(criteriaID))
									if err != nil || criteria == nil {
										io.Println(fmt.Sprintf("Criteria ID %v not found.", criteriaID))
										return
									}

									// Check if already linked
									for _, link := range links {
										if link.AssessmentCriteriaId == uint(criteriaID) {
											io.Println("This criteria is already linked to the project.")
											return
										}
									}

									// Do the link
									_, err = assessmentCriteriaLinkController.InsertAssessmentCriteriaLink(uint(assessment), uint(criteriaID))
									if err != nil {
										io.Println(fmt.Sprintf("Error linking criteria: %v", err))
										return
									}
									io.Println("Criteria successfully linked!")
								},
							},

							{
								Title: "Update Assessment Criteria Link",
								Action: func(io *utils.MenuIO) {
									io.Println("Updating Assessment Criteria...")
									io.Print("Enter Senior Project ID (-1 to cancel): ")

									input, err := io.ReadInput()
									if err != nil {
										io.Println(fmt.Sprintf("Error reading input: %v", err))
										return
									}
									if input == "-1" {
										return
									}

									seniorProjectID, err := strconv.ParseUint(input, 10, 32)
									if err != nil {
										io.Println(fmt.Sprintf("Invalid senior project ID: %v", err))
										return
									}

									assessment, err := assessmentController.RetrieveAssessment(uint(seniorProjectID))
									if err != nil {
										if assessment == nil {
											assessment, err = assessmentController.InsertAssessment(uint(seniorProjectID))
											if err != nil {
												log.Printf("Error inserting assessment: %v", err)
												return
											}
										} else {
											log.Printf("Error retrieving assessment: %v", err)
											return
										}
									}

									mappers, err := assessmentCriteriaLinkController.ListProjectAssessmentCriteriaLinks(uint(seniorProjectID))
									if err != nil {
										log.Printf("Error listing assessment criteria: %v", err)
										return
									}

									io.Println("Current Linked Criteria:")
									for _, mapper := range mappers {
										assessmentCriteria, err := assessmentCriteriaController.RetrieveAssessmentCriteria(mapper.AssessmentCriteriaId)
										if err != nil {
											log.Printf("Error retrieving assessmentCriteria (ID %d): %v", mapper.AssessmentCriteriaId, err)
											continue
										}
										if assessmentCriteria == nil {
											log.Printf("No assessment criteria found with ID %d", mapper.AssessmentCriteriaId)
											continue
										}
										io.Println(fmt.Sprintf("Criteria ID: %v, Criteria Name: %v", assessmentCriteria.ID, assessmentCriteria.CriteriaName))
									}

									io.Print("Enter Criteria ID to update (-1 to cancel): ")
									input, err = io.ReadInput()
									if err != nil {
										io.Println(fmt.Sprintf("Error reading input: %v", err))
										return
									}
									if input == "-1" {
										return
									}

									criteriaID, err := strconv.ParseUint(input, 10, 32)
									if err != nil {
										io.Println(fmt.Sprintf("Invalid criteria ID: %v", err))
										return
									}

									assessmentCriteriaLink, err := assessmentCriteriaLinkController.RetrieveAssessmentCriteriaLink(assessment.ID, uint(criteriaID))
									if err != nil {
										io.Println(fmt.Sprintf("Error retrieving assessmentCriteriaLink: %v", err))
										return
									}
									if assessmentCriteriaLink == nil {
										io.Println("AssessmentCriteriaLink not found.")
										return
									}

									io.Print("Enter new Criteria ID (-1 to cancel): ")
									input, err = io.ReadInput()
									if err != nil {
										io.Println(fmt.Sprintf("Error reading input: %v", err))
										return
									}
									if input == "-1" {
										return
									}

									newCriteriaID, err := strconv.ParseUint(input, 10, 32)
									if err != nil {
										io.Println(fmt.Sprintf("Invalid new criteria ID: %v", err))
										return
									}

									// Check for duplication
									for _, mapper := range mappers {
										if mapper.AssessmentCriteriaId == uint(newCriteriaID) {
											io.Println("This criteria is already linked to the assessment.")
											return
										}
									}

									assessmentCriteriaLink.AssessmentCriteriaId = uint(newCriteriaID)
									err = assessmentCriteriaLinkController.UpdateAssessmentCriteriaLink(assessmentCriteriaLink.ID, assessmentCriteriaLink)
									if err != nil {
										io.Println(fmt.Sprintf("Error updating assessmentCriteriaLink: %v", err))
									} else {
										io.Println("Assessment criteria link updated successfully.")
									}
								},
							},

							{
								Title: "Delete Criteria Link from Assessment",
								Action: func(io *utils.MenuIO) {
									io.Println("Deleting Criteria Link from Assessment...")
									io.Print("Enter Senior Project ID (-1 to cancel): ")
									input, err := io.ReadInput()
									if err != nil {
										io.Println(fmt.Sprintf("Error reading input: %v", err))
										return
									}
									if input == "-1" {
										return
									}

									seniorProjectID, err := strconv.ParseUint(input, 10, 32)
									if err != nil {
										io.Println(fmt.Sprintf("Invalid project ID: %v", err))
										return
									}

									assessment, err := assessmentController.RetrieveAssessment(uint(seniorProjectID))
									if err != nil || assessment == nil {
										io.Println("Assessment not found.")
										return
									}

									io.Println("Current Linked Criteria:")
									mappers, err := assessmentCriteriaLinkController.ListProjectAssessmentCriteriaLinks(uint(seniorProjectID))
									if err != nil {
										io.Println(fmt.Sprintf("Error retrieving criteria links: %v", err))
										return
									}
									if len(mappers) == 0 {
										io.Println("No criteria linked to this assessment.")
										return
									}

									for _, mapper := range mappers {
										criteria, err := assessmentCriteriaController.RetrieveAssessmentCriteria(mapper.AssessmentCriteriaId)
										if err != nil {
											log.Printf("Error retrieving criteria (ID %d): %v", mapper.AssessmentCriteriaId, err)
											continue
										}
										if criteria == nil {
											log.Printf("AssessmentCriteria with ID %d not found", mapper.AssessmentCriteriaId)
											continue
										}
										io.Println(fmt.Sprintf("Criteria ID: %v, Name: %v", criteria.ID, criteria.CriteriaName))
									}

									io.Print("Enter Criteria ID to unlink (-1 to cancel): ")
									input, err = io.ReadInput()
									if err != nil {
										io.Println(fmt.Sprintf("Error reading input: %v", err))
										return
									}
									if input == "-1" {
										return
									}

									criteriaID, err := strconv.ParseUint(input, 10, 32)
									if err != nil {
										io.Println(fmt.Sprintf("Invalid criteria ID: %v", err))
										return
									}

									err = assessmentCriteriaLinkController.DeleteAssessmentCriteriaLink(uint(assessment.ID), uint(criteriaID))
									if err != nil {
										io.Println(fmt.Sprintf("Error unlinking criteria: %v", err))
									} else {
										io.Println("Criteria unlinked successfully.")
									}
								},
							}},
					},
				},
			},
		},
	}, nil, nil)

	builder.Show()
}
