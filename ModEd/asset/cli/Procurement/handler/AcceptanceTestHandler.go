package handler

import (
	"ModEd/asset/controller"
	util "ModEd/asset/util"
	"fmt"
)

func AcceptanceTestHandler(facade *controller.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printAcceptanceTestOptions()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			ListAllTORs(facade)
			torid := util.GetUintInput("Enter TOR ID: ")
			PrintCategoriesByTOR(torid, facade)
			WaitForEnter()
		case "2":
			WaitForEnter()
		}
	}
}

func printAcceptanceTestOptions() {
	fmt.Println(":/Procurement/Main")
	fmt.Println()
	fmt.Println("--AcceptanceTest Functions--")
	fmt.Println("  1:\tAccept")
	fmt.Println("  2:\tReject")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
func PrintCategor7tiesByTOR(torID uint, facade *controller.ProcurementControllerFacade) error {
	details, err := facade.AcceptanceTest.GetQuotationDetailsByTOR(torID)
	if err != nil {
		return err
	}

	categorySet := make(map[uint]bool)
	for _, d := range details {
		categorySet[d.CategoryID] = true
	}

	var categoryIDs []uint
	for id := range categorySet {
		categoryIDs = append(categoryIDs, id)
	}

	condition := map[string]interface{}{"id": categoryIDs}
	categories, err := controller.NewCategoryController().List(condition)
	if err != nil {
		return err
	}

	for _, category := range categories {
		fmt.Println("Category Name:", category.CategoryName)
		if category.Description != nil {
			fmt.Println("Description:", *category.Description)
		} else {
			fmt.Println("Description: (none)")
		}
		fmt.Println("-----")
	}
	// fmt.Println(model.Category.GetID(id))
	return nil
}

func PrintCategoriesByTOR(torID uint, facade *controller.ProcurementControllerFacade) error {
	details, err := facade.AcceptanceTest.GetQuotationDetailsByTOR(torID)
	if err != nil {
		return err
	}

	categorySet := make(map[uint]bool)
	for _, d := range details {
		categorySet[d.CategoryID] = true
	}

	var categoryIDs []uint
	for id := range categorySet {
		categoryIDs = append(categoryIDs, id)
	}

	categories, err := facade.AcceptanceTest.GetCategoriesByIDs(categoryIDs)
	if err != nil {
		return err
	}

	for _, category := range categories {
		fmt.Println("Category Name:", category.CategoryName)
		if category.Description != nil {
			fmt.Println("Description:", *category.Description)
		} else {
			fmt.Println("Description: (none)")
		}
		fmt.Println("-----")
	}

	return nil
}
