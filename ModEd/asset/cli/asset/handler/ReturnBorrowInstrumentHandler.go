package handler

// Wrote by MEP-1012

import (
	"ModEd/asset/controller"
)

type ReturnHandlerStrategy struct {
	controller *controller.BorrowInstrumentController
}

func NewDeleteHandlerStrategy(controller *controller.BorrowInstrumentController) *ReturnHandlerStrategy {
	return &ReturnHandlerStrategy{controller: controller}
}

//func (handler ReturnHandlerStrategy) Execute() error {
//
//	records, err := handler.controller.List(nil)
//	if err != nil {
//		return err
//	}
//
//	fmt.Printf("Total %v record(s)\n", len(records))
//	fmt.Println()
//
//	for _, record := range records {
//		fmt.Println(record)
//	}
//
//	fmt.Print("Enter ID to return : ")
//	var inputBuffer string
//	fmt.Scanln(&inputBuffer)
//
//	idNum, err := strconv.Atoi(inputBuffer)
//	if err != nil {
//		fmt.Println("Invalid ID format. Please enter a valid number.")
//		return nil
//	}
//	err = handler.controller.DeleteByID(uint(idNum))
//	if err != nil {
//		fmt.Println("Error deleting record:", err)
//		return nil
//	}
//	fmt.Println("Record successfully deleted!")
//	return nil
//}
