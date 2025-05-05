package handler

import (
	procurement "ModEd/asset/controller"
	// model "ModEd/asset/model"
	// util "ModEd/asset/util"
	// "encoding/json"
	// "fmt"
	// "time"
)

func QuotationSupplierHandler(facade *procurement.ProcurementControllerFacade) {}

// 	inputBuffer := ""

// 	for inputBuffer != "back" {
// 		util.ClearScreen()
// 		printQuotationSupplierOptions()
// 		inputBuffer = util.GetCommandInput()

// 		switch inputBuffer {
// 		case "1":
// 			fmt.Println("Import Quotations")
// 			WaitForEnter()
// 		case "2":
// 			fmt.Println("List by TOR ID")
// 			WaitForEnter()
// 		case "3":
// 			fmt.Println("Quotation Selection")
// 			WaitForEnter()
// 		case "4":
// 			fmt.Println("Quotaion ")
// 			WaitForEnter()
// 		case "5":
// 			fmt.Println("Delete Procurement")
// 			ListAllProcurements(facade)
// 			id := util.GetUintInput("Enter procurement ID to delete: ")
// 			err := facade.Procurement.Delete(id)
// 			if err != nil {
// 				fmt.Printf("Failed to delete procurement with ID %d: %v\n", id, err)
// 				return
// 			}
// 			fmt.Printf("Procurement with ID %d deleted successfully.\n", id)
// 			WaitForEnter()
// 		case "6":
// 		case "7":
// 		case "8":
// 		case "9":
// 		case "":
// 		}

// 		util.ClearScreen()
// 	}

// 	util.ClearScreen()
// }

// func printQuotationSupplierOptions() {
// 	fmt.Println(":/Procurement/Main")
// 	fmt.Println()
// 	fmt.Println("--View Quotation Functions--")
// 	fmt.Println("  1:\tImport Quotations")
// 	fmt.Println("  2:\tList by TOR ID")
// 	fmt.Println("--Quotation Selection Functions--")
// 	fmt.Println("  3:\tView Quotation by ID")
// 	fmt.Println("  4:\tQuotation by ")
// 	fmt.Println("  5:\tDelete Quotation")

// 	fmt.Println("--Supplier Functions--")
// 	fmt.Println("  6:\tList All Supplier")
// 	fmt.Println("  7:\tView Supplier by ID")
// 	fmt.Println("  8:\tUpdate Supplier Status")
// 	fmt.Println("  9:\tDelete Supplier")
// 	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit")
// 	fmt.Println()
// }

// func ListAllQuotations() {
// 	filePath := filepath.Join("data", "quotations.json")

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		fmt.Println("Failed to open quotations file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	var quotations []quotation.Quotation
// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&quotations); err != nil {
// 		fmt.Println("Failed to decode quotations JSON:", err)
// 		return
// 	}

// 	fmt.Println("Quotation List:")
// 	for _, q := range quotations {
// 		submitTime := "-"
// 		if q.SubmittedTime != nil {
// 			submitTime = q.SubmittedTime.Format("2006-01-02 15:04:05")
// 		}
// 		fmt.Printf("ID: %d, Supplier: %s, Amount: %.2f, Status: %s, SubmittedTime: %s\n",
// 			q.QuotationID, q.SupplierName, q.Amount, q.Status, submitTime)
// 	}
// }
