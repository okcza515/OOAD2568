package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
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
