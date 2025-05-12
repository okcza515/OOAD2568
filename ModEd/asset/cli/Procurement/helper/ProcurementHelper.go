// MEP-1014
package helper

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
)

func DisplayProcurementList(procurements []model.Procurement) {
	if len(procurements) == 0 {
		fmt.Println("No procurements available.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\n--- Available Procurements ---")
	for _, procurement := range procurements {
		fmt.Printf("  ID: %d | Status: %s | Created At: %s\n",
			procurement.ProcurementID, procurement.Status, procurement.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func DisplayTORList(tors []model.TOR) {
	if len(tors) == 0 {
		fmt.Println("No TORs available.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\n--- Available TORs ---")
	for _, tor := range tors {
		fmt.Printf("  ID: %d | Scope: %s | Committee: %s\n",
			tor.TORID, tor.Scope, tor.Committee)
	}
}

func SelectProcurement(facade *controller.ProcurementControllerFacade) (*model.Procurement, error) {
	procurements, err := facade.Procurement.ListAllProcurement()
	if err != nil || len(*procurements) == 0 {
		fmt.Println("No procurements available.")
		util.PressEnterToContinue()
		return nil, fmt.Errorf("no procurements found")
	}

	if len(*procurements) == 0 {
		fmt.Println("No procurements available.")
		util.PressEnterToContinue()
		return nil, fmt.Errorf("no procurements found")
	}

	DisplayProcurementList(*procurements)
	procurementID := util.GetUintInput("\nEnter Procurement ID: ")

	return facade.Procurement.GetProcurementByID(procurementID)
}

func SelectTOR(facade *controller.ProcurementControllerFacade) (*model.TOR, error) {
	tors, err := facade.TOR.GetAllTORs()
	if err != nil || len(tors) == 0 {
		fmt.Println("No TORs available.")
		util.PressEnterToContinue()
		return nil, fmt.Errorf("no TORs found")
	}

	DisplayTORList(tors)
	torID := util.GetUintInput("\nEnter TOR ID: ")

	return facade.TOR.GetTORByID(torID)
}

func DeleteEntity(deleteFunc func(uint) error, entityName string, id uint) error {
	fmt.Printf("Delete %s\n", entityName)

	err := deleteFunc(id)
	if err != nil {
		fmt.Printf("Failed to delete %s: %v\n", entityName, err)
		util.PressEnterToContinue()
		return err
	}

	fmt.Printf("%s deleted successfully.\n", entityName)
	util.PressEnterToContinue()
	return nil
}
