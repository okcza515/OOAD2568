package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
	"log"
	"strconv"
)

type AssessmentCriteriaLinkHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewAssessmentCriteriaLinkHandler(instanceStorer *controller.InstanceStorer) *AssessmentCriteriaLinkHandler {
	return &AssessmentCriteriaLinkHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

// link criteria to an assessment
func (h *AssessmentCriteriaLinkHandler) LinkCriteriaToAssessment(io *core.MenuIO) {
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

	criteriaList, err := h.instanceStorer.AssessmentCriteria.List(map[string]interface{}{})
	if err != nil {
		io.Println(fmt.Sprintf("Error listing criteria: %v", err))
		return
	}

	for _, c := range criteriaList {
		io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
	}

	io.Print("Enter Criteria ID to link (-1 to cancel): ")
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

	assessment, err := h.instanceStorer.Assessment.RetrieveAssessmentBySeniorProjectId(uint(seniorProjectID))
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving assessment: %v", err))
		return
	}

	err = h.instanceStorer.AssessmentCriteriaLink.Insert(&model.AssessmentCriteriaLink{
		AssessmentId:         assessment.ID,
		AssessmentCriteriaId: uint(criteriaID),
	})
	if err != nil {
		io.Println(fmt.Sprintf("Error linking criteria: %v", err))
		return
	}

	io.Println("Criteria linked successfully.")
}

// UpdateLink updates a link of criteria to an assessment, meaning to change criteria link to the assessment
func (h *AssessmentCriteriaLinkHandler) UpdateLink(io *core.MenuIO) {
	io.Print("Enter Senior Project ID (-1 to cancel): ")
	seniorProjectID, err := io.ReadInputID()
	if err != nil {
		io.Println(fmt.Sprintf("Error reading input: %v", err))
		return
	}
	if seniorProjectID == 0 {
		return
	}

	io.Print("Enter Criteria ID to update (-1 to cancel): ")
	criteriaID, err := io.ReadInputID()
	if err != nil {
		io.Println(fmt.Sprintf("Error reading input: %v", err))
		return
	}
	if criteriaID == 0 {
		return
	}

	io.Print("Enter New Criteria ID (-1 to cancel): ")
	newCriteriaID, err := io.ReadInputID()
	if err != nil {
		io.Println(fmt.Sprintf("Error reading input: %v", err))
		return
	}
	if newCriteriaID == 0 {
		return
	}

	link, err := h.instanceStorer.AssessmentCriteriaLink.RetrieveAssessmentCriteriaLink(uint(seniorProjectID), uint(criteriaID))
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving link: %v", err))
		return
	}

	link.AssessmentCriteriaId = uint(newCriteriaID)
	err = h.instanceStorer.AssessmentCriteriaLink.UpdateByID(link)
	if err != nil {
		io.Println(fmt.Sprintf("Error updating link: %v", err))
		return
	}

	io.Println("Link updated successfully.")

}

func (h *AssessmentCriteriaLinkHandler) ListCriteriaLinkedToAssessment(io *core.MenuIO) {
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

	links, err := h.instanceStorer.AssessmentCriteriaLink.ListProjectAssessmentCriteriaLinks(uint(seniorProjectID))
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
		criteria, err := h.instanceStorer.AssessmentCriteria.RetrieveByID(link.AssessmentCriteriaId)
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
