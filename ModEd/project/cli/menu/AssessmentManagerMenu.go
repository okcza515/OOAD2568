package menu

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"log"
	"strconv"
)

func BuildAssessmentManagerMenu(
	assessmentCriteriaController *controller.AssessmentCriteriaController,
	assessmentController *controller.AssessmentController,
	assessmentCriteriaLinkController *controller.AssessmentCriteriaLinkController,
	scoreAdvisorController *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor],
	scoreCommitteeController *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee],
) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Assessment Manager",
		Children: []*utils.MenuItem{
			{
				Title: "List All Criteria Linked to Assessment",
				Action: listCriteriaLinkedToAssessment(
					assessmentCriteriaController,
					assessmentCriteriaLinkController,
				),
			},
			{
				Title: "Link Criteria to Assessment",
				Action: linkCriteriaToAssessment(
					assessmentCriteriaController,
					assessmentController,
					assessmentCriteriaLinkController,
				),
			},
			{
				Title: "Update Assessment Criteria Link",
				Action: updateCriteriaLink(
					assessmentCriteriaController,
					assessmentController,
					assessmentCriteriaLinkController,
				),
			},
			{
				Title: "Delete Criteria Link from Assessment",
				Action: deleteCriteriaLink(
					assessmentController,
					assessmentCriteriaLinkController,
					assessmentCriteriaController,
				),
			},
			BuildAssessmentScoreManagerMenu(
				assessmentController,
				assessmentCriteriaLinkController,
				assessmentCriteriaController,
				scoreAdvisorController,
				scoreCommitteeController,
			),
			BuildAssessmentCriteriaManagerMenu(assessmentCriteriaController),
		},
	}
}

func listCriteriaLinkedToAssessment(criteriaCtrl *controller.AssessmentCriteriaController, linkCtrl *controller.AssessmentCriteriaLinkController) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
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

		links, err := linkCtrl.ListProjectAssessmentCriteriaLinks(uint(seniorProjectID))
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
			criteria, err := criteriaCtrl.RetrieveByID(link.AssessmentCriteriaId)
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
	}
}

func linkCriteriaToAssessment(criteriaCtrl *controller.AssessmentCriteriaController, assessmentCtrl *controller.AssessmentController, linkCtrl *controller.AssessmentCriteriaLinkController) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
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

		links, err := linkCtrl.ListProjectAssessmentCriteriaLinks(uint(seniorProjectID))
		if err != nil {
			io.Println(fmt.Sprintf("Error retrieving linked criteria: %v", err))
			return
		}

		io.Println("Currently linked criteria:")
		for _, link := range links {
			criteria, err := criteriaCtrl.RetrieveByID(link.AssessmentCriteriaId)
			if err != nil || criteria == nil {
				log.Printf("Error retrieving criteria (ID %d): %v", link.AssessmentCriteriaId, err)
				continue
			}
			io.Println(fmt.Sprintf("ID: %v, Name: %v", criteria.ID, criteria.CriteriaName))
		}

		io.Println("\nAvailable Assessment Criteria:")
		allCriteria, err := criteriaCtrl.List(map[string]interface{}{})
		if err != nil {
			io.Println(fmt.Sprintf("Error listing all criteria: %v", err))
			return
		}
		for _, c := range allCriteria {
			io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
		}

		io.Print("Enter Criteria ID to link (-1 to cancel): ")
		criteriaId, err := io.ReadInputID()
		if err != nil {
			return
		}

		for _, link := range links {
			if link.AssessmentCriteriaId == criteriaId {
				io.Println("This criteria is already linked to the project.")
				return
			}
		}

		assessment, err := assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(seniorProjectID))
		if err != nil {
			io.Println(fmt.Sprintf("Error retrieving assessment: %v", err))
			return
		}

		_, err = linkCtrl.InsertAssessmentCriteriaLink(uint(assessment.ID), uint(criteriaId))
		if err != nil {
			io.Println(fmt.Sprintf("Error linking criteria: %v", err))
			return
		}
		io.Println("Criteria successfully linked!")
	}
}

func updateCriteriaLink(criteriaCtrl *controller.AssessmentCriteriaController, assessmentCtrl *controller.AssessmentController, linkCtrl *controller.AssessmentCriteriaLinkController) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
		io.Println("Updating Assessment Criteria...")
		io.Print("Enter Senior Project ID (-1 to cancel): ")

		seniorProjectID, err := io.ReadInputID()
		if err != nil {
			return
		}

		assessment, err := assessmentCtrl.RetrieveByID(seniorProjectID)
		if err != nil {
			if assessment == nil {
				assessment, err = assessmentCtrl.InsertAssessment(seniorProjectID)
				if err != nil {
					log.Printf("Error inserting assessment: %v", err)
					return
				}
			} else {
				log.Printf("Error retrieving assessment: %v", err)
				return
			}
		}

		mappers, err := linkCtrl.ListProjectAssessmentCriteriaLinks(seniorProjectID)
		if err != nil {
			log.Printf("Error listing assessment criteria: %v", err)
			return
		}

		io.Println("Current Linked Criteria:")
		for _, mapper := range mappers {
			assessmentCriteria, err := criteriaCtrl.RetrieveByID(mapper.AssessmentCriteriaId)
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
		criteriaId, err := io.ReadInputID()
		if err != nil {
			return
		}

		assessmentCriteriaLink, err := linkCtrl.RetrieveAssessmentCriteriaLink(assessment.ID, criteriaId)
		if err != nil {
			io.Println(fmt.Sprintf("Error retrieving assessmentCriteriaLink: %v", err))
			return
		}
		if assessmentCriteriaLink == nil {
			io.Println("AssessmentCriteriaLink not found.")
			return
		}

		io.Print("Enter new Criteria ID (-1 to cancel): ")
		newCriteriaID, err := io.ReadInputID()
		if err != nil {
			return
		}

		for _, mapper := range mappers {
			if mapper.AssessmentCriteriaId == newCriteriaID {
				io.Println("This criteria is already linked to the assessment.")
				return
			}
		}

		assessmentCriteriaLink.AssessmentCriteriaId = newCriteriaID
		err = linkCtrl.UpdateByID(assessmentCriteriaLink)
		if err != nil {
			io.Println(fmt.Sprintf("Error updating assessmentCriteriaLink: %v", err))
		} else {
			io.Println("Assessment criteria link updated successfully.")
		}
	}
}

func deleteCriteriaLink(assessmentCtrl *controller.AssessmentController, linkCtrl *controller.AssessmentCriteriaLinkController, criteriaCtrl *controller.AssessmentCriteriaController) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
		io.Println("Deleting Criteria Link from Assessment...")
		io.Print("Enter Senior Project ID (-1 to cancel): ")
		seniorProjectID, err := io.ReadInputID()
		if err != nil {
			return
		}

		assessment, err := assessmentCtrl.RetrieveByID(seniorProjectID)
		if err != nil || assessment == nil {
			io.Println("Assessment not found.")
			return
		}

		io.Println("Current Linked Criteria:")
		mappers, err := linkCtrl.ListProjectAssessmentCriteriaLinks(seniorProjectID)
		if err != nil {
			io.Println(fmt.Sprintf("Error retrieving criteria links: %v", err))
			return
		}
		if len(mappers) == 0 {
			io.Println("No criteria linked to this assessment.")
			return
		}

		for _, mapper := range mappers {
			criteria, err := criteriaCtrl.RetrieveByID(mapper.AssessmentCriteriaId)
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
		criteriaId, err := io.ReadInputID()
		if err != nil {
			return
		}

		err = linkCtrl.DeleteAssessmentCriteriaLink(assessment.ID, criteriaId)
		if err != nil {
			io.Println(fmt.Sprintf("Error unlinking criteria: %v", err))
		} else {
			io.Println("Criteria unlinked successfully.")
		}
	}
}
