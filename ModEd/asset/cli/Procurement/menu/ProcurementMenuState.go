// File: ProcurementMenuState.go
package menu

import (
	"ModEd/asset/cli/Procurement/helper"
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
	"time"
)

type ProcurementMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewProcurementMenuState(manager *cli.CLIMenuStateManager) *ProcurementMenuState {
	facade, err := controller.CreateProcurementControllerFacade()
	if err != nil {
		fmt.Println("Failed to create ProcurementControllerFacade:", err)
		return nil
	}

	handlerContext := handler.NewHandlerContext()
	menu := &ProcurementMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	// 1. Create TOR and Procurement
	handlerContext.AddHandler("1", "Create TOR and Procurement", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Create TOR and Procurement")
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if err != nil {
				fmt.Println("Failed to list Instrument Requests:", err)
				util.PressEnterToContinue()
				return err
			}

			approvedRequests := []model.InstrumentRequest{}
			for _, req := range *requests {
				if req.Status == model.InstrumentRequestStatusApproved {
					approvedRequests = append(approvedRequests, req)
				}
			}

			if len(approvedRequests) == 0 {
				fmt.Println("No approved Instrument Requests available. Please complete Budget Approvals first.")
				util.PressEnterToContinue()
				return nil
			}

			fmt.Println("\n--- Approved Instrument Requests ---")
			for _, req := range approvedRequests {
				fmt.Printf("  ID: %d | Department ID: %d | Status: %s\n",
					req.InstrumentRequestID, req.DepartmentID, req.Status)
			}
			// Get Input
			requestID := util.GetUintInput("Enter Instrument Request ID: ")
			scope := util.GetStringInput("Enter TOR Scope: ")
			deliverables := util.GetStringInput("Enter TOR Deliverables: ")
			timeline := util.GetStringInput("Enter TOR Timeline: ")
			committee := util.GetStringInput("Enter TOR Committee: ")

			tor := controller.NewTORBuilder().
				WithInstrumentRequestID(requestID).
				WithScope(scope).
				WithDeliverables(deliverables).
				WithTimeline(timeline).
				WithCommittee(committee).
				WithCreatedAt(time.Now()).
				Build()

			err = facade.TOR.CreateTOR(tor)
			if err != nil {
				fmt.Println("Failed to create TOR:", err)
				util.PressEnterToContinue()
				return err
			}
			fmt.Println("TOR created successfully with ID:", tor.TORID)

			procurement := controller.NewProcurementBuilder().
				WithTOR(tor).
				WithStatus(model.ProcurementStatusPending).
				WithCreatedAt(time.Now()).
				Build()

			err = facade.Procurement.CreateProcurement(procurement)
			if err != nil {
				fmt.Println("Failed to create Procurement:", err)
			} else {
				fmt.Println("Procurement created successfully!")
			}
			util.PressEnterToContinue()
			return nil
		},
	})

	// List All Procurements
	handlerContext.AddHandler("2", "List All Procurements", handler.FuncStrategy{
		Action: func() error {
			procurements, err := facade.Procurement.ListAllProcurement()
			if err != nil {
				fmt.Println("Failed to list procurements:", err)
				util.PressEnterToContinue()
				return err
			}

			// Now the check is inside the helper
			if len(*procurements) == 0 {
				fmt.Println("No procurements available.")
				util.PressEnterToContinue()
				return nil
			}

			helper.DisplayProcurementList(*procurements)
			util.PressEnterToContinue()
			return nil
		},
	})

	// Listing All TORs
	handlerContext.AddHandler("3", "List All TORs", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("List All TORs")

			// Correct method call
			tors, err := facade.TOR.GetAllTORs()
			if err != nil {
				fmt.Println("Failed to list TORs:", err)
				util.PressEnterToContinue()
				return err
			}

			// Display the list — no need for "*"
			helper.DisplayTORList(tors)
			util.PressEnterToContinue()
			return nil
		},
	})

	// Viewing Procurement Details
	handlerContext.AddHandler("4", "View Procurement Detail by ID", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("View Procurement Detail by ID")
			procurement, err := helper.SelectProcurement(facade)
			if err != nil {
				fmt.Println("Failed to retrieve procurement details:", err)
				return err
			}

			fmt.Println("\n--- Procurement Details ---")
			fmt.Printf("ID: %d\n", procurement.ProcurementID)
			fmt.Printf("Status: %s\n", procurement.Status)
			fmt.Printf("Created At: %s\n", procurement.CreatedAt.Format("2006-01-02 15:04:05"))

			util.PressEnterToContinue()
			return nil
		},
	})
	// Viewing TOR Details
	handlerContext.AddHandler("5", "View TOR Detail by ID", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("View TOR Detail by ID")
			tor, err := helper.SelectTOR(facade)
			if err != nil {
				fmt.Println("Failed to retrieve TOR details:", err)
				return err
			}

			fmt.Println("\n--- TOR Details ---")
			fmt.Printf("ID: %d\n", tor.TORID)
			fmt.Printf("Scope: %s\n", tor.Scope)
			fmt.Printf("Deliverables: %s\n", tor.Deliverables)
			fmt.Printf("Timeline: %s\n", tor.Timeline)
			fmt.Printf("Committee: %s\n", tor.Committee)

			util.PressEnterToContinue()
			return nil
		},
	})

	// Deleting Procurement
	handlerContext.AddHandler("6", "Delete Procurement", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Delete Procurement")
			procurement, err := helper.SelectProcurement(facade)
			if err != nil {
				return err
			}
			return helper.DeleteEntity(facade.Procurement.DeleteProcurement, "Procurement", procurement.ProcurementID)
		},
	})

	// Deleting TOR
	handlerContext.AddHandler("7", "Delete TOR", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Delete TOR")
			tor, err := helper.SelectTOR(facade)
			if err != nil {
				return err
			}
			return helper.DeleteEntity(facade.TOR.DeleteTOR, "TOR", tor.TORID)
		},
	})

	manager.AddMenu(string(MENU_QUOTATION), NewQuotationMenuState(manager))
	handlerContext.AddHandler("8", "Quotation Management", handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_QUOTATION))))

	// Back Handler
	handlerContext.AddBackHandler(handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_PROCUREMENT_MAIN))))

	return menu
}

func (menu *ProcurementMenuState) Render() {
	fmt.Println()
	fmt.Println(":/procurement")
	fmt.Println()
	fmt.Println("Procurement Management Menu:")
	menu.handlerContext.ShowMenu()
	fmt.Println()
}

func (menu *ProcurementMenuState) HandleUserInput(input string) error {
	return menu.handlerContext.HandleInput(input)
}
