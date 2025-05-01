package handler

// import (
// 	procurement "ModEd/asset/controller"
// 	"ModEd/asset/util"
// 	"fmt"
// )

// func BudgetAllocationHandler(facade *procurement.ProcurementControllerFacade) {
// 	inputBuffer := ""

// 	for inputBuffer != "back" {
// 		util.ClearScreen()
// 		printBudgetAllocationOption()
// 		inputBuffer = util.GetCommandInput()

// 		switch inputBuffer {
// 		case "1":
// 			// fmt.Println("Create New Budget Allocation")
// 			// AllocationID := GetUintInput("Enter Allocation ID: ")
// 			fmt.Println("Allocate Budget successfully")
// 			WaitForEnter()
// 		case "2":
// 			id := GetUintInput("Enter BudgetID: ")
// 			NewAmount := GetFloatInput("Enter new amount: ")
// 			facade.BudgetAllocation.UpdateBudget(id, NewAmount)
// 			fmt.Println("Budget Updated")
// 			WaitForEnter()
// 		case "3":
// 			fmt.Println("Get Budget Allocation by ID")
// 			id := GetUintInput("Enter BudgetID: ")
// 			facade.BudgetAllocation.GetByID(id)
// 			WaitForEnter()
// 		}

// 		util.ClearScreen()
// 	}

// 	util.ClearScreen()
// }

// func printBudgetAllocationOption() {
// 	fmt.Println(":/Procurement/BudgetAllocation")
// 	fmt.Println()
// 	fmt.Println("--BudgetAllocation Function--")
// 	fmt.Println("  1:\tAllocate Budget")
// 	fmt.Println("  2:\tUpdate Budget")
// 	fmt.Println("  3:\tGet Budget Allocation by ID")
// 	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit)")
// 	fmt.Println()
// }

// func GetUintInput(prompt string) uint {
// 	var input uint
// 	fmt.Print(prompt)
// 	_, err := fmt.Scanln(&input)
// 	if err != nil {
// 		fmt.Println("Invalid input. Please enter a positive number.")
// 		return GetUintInput(prompt)
// 	}
// 	return input
// }

// func GetFloatInput(prompt string) float64 {
// 	var input float64
// 	fmt.Print(prompt)
// 	_, err := fmt.Scanln(&input)
// 	if err != nil {
// 		fmt.Println("Invalid input. Please enter a number.")
// 		return GetFloatInput(prompt)
// 	}
// 	return input
// }
