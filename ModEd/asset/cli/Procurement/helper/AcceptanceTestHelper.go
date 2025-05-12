package helper

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"fmt"
)

func ListAllAcceptanceRequests(facade *controller.ProcurementControllerFacade) {
	acceptanceRequests, err := facade.Acceptance.ListAllApprovals()
	if err != nil {
		fmt.Println("Failed to fetch acceptance requests:", err)
		return
	}

	if len(acceptanceRequests) == 0 {
		fmt.Println("No acceptance requests found.")
		return
	}

	fmt.Println("Acceptance Request List:")
	for _, req := range acceptanceRequests {
		approverID := "Not Assigned"
		if req.ApproverID != nil {
			approverID = fmt.Sprintf("%d", *req.ApproverID)
		}
		createdAt := "-"
		if !req.CreatedAt.IsZero() {
			createdAt = req.CreatedAt.Format("2006-01-02 15:04:05")
		}
		approvalTime := "-"
		if req.ApprovalTime != nil {
			approvalTime = req.ApprovalTime.Format("2006-01-02 15:04:05")
		}
		fmt.Printf("ID: %d | ProcurementID: %d | ApproverID: %s | Status: %s | CreatedAt: %s | ApprovalTime: %s\n",
			req.AcceptanceApprovalID, req.ProcurementID, approverID, req.Status, createdAt, approvalTime)
	}
}

func PrintQuotationDetailsByAcceptance(facade *controller.ProcurementControllerFacade, acceptanceID uint) {
	fmt.Printf("Searching for Quotation Details for Acceptance Request ID: %d\n", acceptanceID)

	details, err := facade.Acceptance.GetQuotationDetailsByAcceptance(acceptanceID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(details) == 0 {
		fmt.Printf("No quotation details found for Acceptance Request ID: %d\n", acceptanceID)
		return
	}

	var allCriteria []model.AcceptanceCriteria
	if err := facade.GetDB().Find(&allCriteria).Error; err != nil {
		fmt.Println("Error loading criteria:", err)
		return
	}

	criteriaMap := make(map[uint][]model.AcceptanceCriteria)
	for _, criteria := range allCriteria {
		criteriaMap[criteria.CategoryID] = append(criteriaMap[criteria.CategoryID], criteria)
	}

	fmt.Printf("\nQuotation Details for Acceptance Request ID: %d\n", acceptanceID)
	totalPrice := 0.0
	for _, detail := range details {
		fmt.Printf("QuotationDetailID: %d\n", detail.QuotationDetailID)
		fmt.Printf("InstrumentLabel: %s\n", detail.InstrumentLabel)
		if detail.Description != nil {
			fmt.Printf("Description: %s\n", *detail.Description)
		} else {
			fmt.Println("Description: (none)")
		}
		fmt.Printf("CategoryID: %d\n", detail.CategoryID)
		fmt.Printf("Quantity: %d\n", detail.Quantity)
		fmt.Printf("Offered Price: %.2f\n", detail.OfferedPrice)

		if criteria, exists := criteriaMap[detail.CategoryID]; exists {
			fmt.Println("Criteria:")
			for _, crit := range criteria {
				fmt.Printf(" - %s: %s\n", crit.CriteriaName, crit.Description)
			}
		} else {
			fmt.Println("No criteria found for this category.")
		}

		fmt.Println("------")
		totalPrice += detail.OfferedPrice * float64(detail.Quantity)
	}

	fmt.Printf("\nTotal Estimated Cost for Acceptance Request ID %d: %.2f\n", acceptanceID, totalPrice)
}
