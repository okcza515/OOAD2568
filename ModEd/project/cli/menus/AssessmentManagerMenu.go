package main

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"log"
	"strconv"
)

func main() {
	db := utils.OpenDatabase("project.db")
	db.Exec("PRAGMA foreign_keys = ON;")

	assessmentCriteriaController := controller.NewAssessmentCriteriaController(db)
	assessmentController := controller.NewAssessmentController(db)
	assessmentCriteriaLinkController := controller.NewAssessmentCriteriaLinkController(db)
	scoreAdvisorController := controller.NewScoreAdvisorController(db)
	scoreCommitteeController := controller.NewScoreCommitteeController(db)

	utils.PrintTitle("Assessment Management CLI")

	builder := utils.NewMenuBuilder(&utils.MenuItem{
		Title: "Assessment Manager", // Changed root menu to Assessment Manager
		Children: []*utils.MenuItem{
			// Original Assessment Manager actions
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

					criteria, err := assessmentCriteriaController.RetrieveAssessmentCriteria(uint(criteriaID))
					if err != nil || criteria == nil {
						io.Println(fmt.Sprintf("Criteria ID %v not found.", criteriaID))
						return
					}

					for _, link := range links {
						if link.AssessmentCriteriaId == uint(criteriaID) {
							io.Println("This criteria is already linked to the project.")
							return
						}
					}

					assessment, err := assessmentController.RetrieveAssessmentBySeniorProjectId(uint(seniorProjectID))
					if err != nil {
						io.Println(fmt.Sprintf("Error retrieving assessment: %v", err))
						return
					}

					_, err = assessmentCriteriaLinkController.InsertAssessmentCriteriaLink(uint(assessment.ID), uint(criteriaID))
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
			},
			// Added Assessment Score Manager as a child
			{
				Title: "Assessment Score Manager",
				Children: []*utils.MenuItem{
					{
						Title: "Senior Project Selection",
						Action: func(io *utils.MenuIO) {
							io.Print("Enter Senior Project ID to view scores of criterias: ")
							input, _ := io.ReadInput()
							projectID, _ := strconv.Atoi(input)

							assessment, err := assessmentController.RetrieveAssessmentBySeniorProjectId(uint(projectID))
							if err != nil {
								io.Println(fmt.Sprintf("Error retrieving assessment: %v", err))
								return
							}

							links, err := assessmentCriteriaLinkController.ListProjectAssessmentCriteriaLinks(uint(projectID))
							if err != nil || len(links) == 0 {
								io.Println("No assessment criteria linked to this project.")
								return
							}

							for _, link := range links {
								criteria, err := assessmentCriteriaController.RetrieveAssessmentCriteria(link.AssessmentCriteriaId)
								if err != nil {
									continue
								}
								io.Println(fmt.Sprintf("Criteria ID: %d | Name: %s", criteria.ID, criteria.CriteriaName))

								advisorScore, err := scoreAdvisorController.RetrieveAdvisorScoreByCondition(
									"assessment", "assessment_criteria_link_id = ?", link.ID,
								)
								if err == nil {
									if score, ok := advisorScore.(*model.ScoreAssessmentAdvisor); ok {
										io.Println(fmt.Sprintf("  Advisor Score: %.2f, By Advisor ID: %d", score.Score, score.AdvisorId))
									} else {
										io.Println("  Advisor Score: -")
									}
								} else {
									io.Println("  Advisor Score: -")
								}

								committeeScores, err := scoreCommitteeController.ListCommitteeScoresByCondition(
									"assessment", "assessment_criteria_link_id = ?", link.ID,
								)
								if err != nil {
									io.Println("  Committee Score: -")
									return
								}

								scoreList, ok := committeeScores.(*[]model.ScoreAssessmentCommittee)
								if !ok {
									io.Println("  Committee Score: -")
									return
								}

								if len(*scoreList) == 0 {
									io.Println("  Committee Score: -")
								} else {
									for _, cs := range *scoreList {
										if cs.AssessmentCriteriaLinkId == link.ID {
											io.Println(fmt.Sprintf("  Committee Score: %.2f, By Committee ID: %d", cs.Score, cs.CommitteeId))
										}
									}
								}
							}

							rootMenu := &utils.MenuItem{
								Title: fmt.Sprintf("Manage Assessment Score of Senior Project %v", input),
								Children: []*utils.MenuItem{
									{
										Title: "Fill Input",
										Action: func(io *utils.MenuIO) {
											io.Print("Enter Criteria ID to score: ")
											criteriaIdStr, _ := io.ReadInput()
											criteriaId, _ := strconv.Atoi(criteriaIdStr)

											link, err := assessmentCriteriaLinkController.RetrieveAssessmentCriteriaLink(assessment.ID, uint(criteriaId))
											if err != nil {
												io.Println("Criteria not found")
												return
											}
											cid := link.ID

											io.Print("Enter scorer type (advisor/committee): ")
											scorer, _ := io.ReadInput()

											io.Print("Enter scorer ID (advisorId/committeeId): ")
											scorerIdStr, _ := io.ReadInput()
											scorerIdVal, err := strconv.ParseUint(scorerIdStr, 10, 64)
											if err != nil {
												io.Println("Invalid ID Input")
												return
											}

											io.Print("Enter score (0.0 - 100.0): ")
											scoreStr, _ := io.ReadInput()
											scoreVal, err := strconv.ParseFloat(scoreStr, 64)
											if err != nil {
												io.Println("Invalid score.")
												return
											}

											if scorer == "advisor" {
												score := model.ScoreAssessmentAdvisor{
													AssessmentCriteriaLinkId: uint(cid),
													AdvisorId:                uint(scorerIdVal),
													Score:                    scoreVal,
												}
												if err := scoreAdvisorController.InsertAdvisorScore(&score); err != nil {
													io.Println(fmt.Sprintf("Failed to insert advisor score: %v", err))
												} else {
													io.Println("Advisor score submitted.")
												}
											} else if scorer == "committee" {
												score := model.ScoreAssessmentCommittee{
													AssessmentCriteriaLinkId: uint(cid),
													CommitteeId:              uint(scorerIdVal),
													Score:                    scoreVal,
												}
												if err := scoreCommitteeController.InsertCommitteeScore(&score); err != nil {
													io.Println(fmt.Sprintf("Failed to insert committee score: %v", err))
												} else {
													io.Println("Committee score submitted.")
												}
											} else {
												io.Println("Invalid scorer type.")
											}
										},
									},
								},
							}

							menuBuilder := utils.NewMenuBuilder(rootMenu, nil, nil)
							menuBuilder.Show()
						},
					},
				},
			},
			// Added Assessment Criteria Manager as a child
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
		},
	}, nil, nil)

	builder.Show()
}
