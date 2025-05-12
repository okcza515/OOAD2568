// MEP-1014
package helper

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

func HandleInstrumentOption(facade *controller.ProcurementControllerFacade) {
	for {
		util.ClearScreen()
		PrintInstrumentList(facade)

		fmt.Println(":/Instrument Menu")
		fmt.Println(" back:\tBack to previous menu")
		fmt.Println()

		cmd := util.GetCommandInput()

		if cmd == "Back" || cmd == "back" {
			return
		} else {
			fmt.Println("Invalid command!")
			util.PressEnterToContinue()
		}
	}
}

func PrintInstrumentList(facade *controller.ProcurementControllerFacade) {
	instruments, err := facade.Instrument.ListAllInstruments()
	if err != nil {
		fmt.Println("Failed to fetch instruments:", err)
		return
	}

	if len(instruments) == 0 {
		fmt.Println("No instruments found.")
		return
	}

	fmt.Println("Available Instruments:")
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("%-5s | %-25s | %-15s | %-10s | %-15s\n", "ID", "Label", "Code", "Status", "Location")
	fmt.Println("---------------------------------------------------------------------------------------------")
	for _, inst := range instruments {
		fmt.Printf("%-5d | %-25s | %-15s | %-10s | %-15s\n",
			inst.ID,
			inst.InstrumentLabel,
			inst.InstrumentCode,
			inst.InstrumentStatus,
			inst.Location,
		)
	}
	fmt.Println("---------------------------------------------------------------------------------------------")
}

func HandleInstrumentDetails(facade *controller.ProcurementControllerFacade) {
	id := util.GetUintInput("Enter Instrument ID: ")

	if _, err := facade.Instrument.ListAllInstruments(); err != nil {
		fmt.Println("Instrument Controller is not initialized or database connection failed:", err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("Attempting to fetch instrument details...")
	instrument, err := facade.Instrument.RetrieveByID(id)

	if err != nil {
		fmt.Println("Failed to retrieve instrument:", err)
		util.PressEnterToContinue()
		return
	}

	if instrument.ID == 0 {
		fmt.Println("Instrument not found.")
		util.PressEnterToContinue()
		return
	}

	util.ClearScreen()
	fmt.Println("Instrument Details:")
	fmt.Printf("  ID: %d\n", instrument.ID)
	fmt.Printf("  Label: %s\n", instrument.InstrumentLabel)
	fmt.Printf("  Code: %s\n", instrument.InstrumentCode)
	fmt.Printf("  Status: %s\n", instrument.InstrumentStatus)
	fmt.Printf("  Room ID: %s\n", instrument.RoomID)
	fmt.Printf("  Location: %s\n", instrument.Location)
	fmt.Printf("  Category ID: %d\n", instrument.CategoryID)
	fmt.Printf("  Cost: %.2f\n", instrument.Cost)
	fmt.Printf("  Budget Year: %d\n", instrument.BudgetYear)

	if instrument.InstrumentBrand != nil {
		fmt.Printf("  Brand: %s\n", *instrument.InstrumentBrand)
	}
	if instrument.InstrumentModel != nil {
		fmt.Printf("  Model: %s\n", *instrument.InstrumentModel)
	}

	util.PressEnterToContinue()
}

func HandleImportInstrument(facade *controller.ProcurementControllerFacade) {
	filename := util.GetStringInput("Enter path to the CSV file (data/instruments.csv): ")

	err := ImportInstrumentsFromCSV(facade.GetDB(), filename)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Import successful.")
	}
	util.PressEnterToContinue()
}

func ImportInstrumentsFromCSV(db *gorm.DB, filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %v", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("failed to read CSV record: %v", err)
		}

		inst := model.Instrument{
			InstrumentLabel:    record[1],
			InstrumentCode:     record[2],
			Description:        &record[3],
			InstrumentStatus:   model.InstrumentStatusEnum(record[4]),
			RoomID:             record[5],
			Location:           record[6],
			CategoryID:         parseUint(record[7]),
			Cost:               parseFloat(record[8]),
			InstrumentSerialID: &record[9],
			BudgetYear:         parseInt(record[10]),
			BudgetSource:       &record[11],
			InstrumentBrand:    &record[12],
			InstrumentModel:    &record[13],
		}

		err = db.Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "instrument_code"}},
				DoUpdates: clause.AssignmentColumns([]string{"instrument_label", "description", "instrument_status", "room_id", "location", "category_id", "cost", "instrument_serial_id", "budget_year", "budget_source", "instrument_brand", "instrument_model"}),
			},
		).Create(&inst).Error

		if err != nil {
			return fmt.Errorf("failed to insert or update instrument with code %s: %v", inst.InstrumentCode, err)
		}
	}

	fmt.Println("Instruments imported successfully.")
	return nil
}

func parseUint(value string) uint {
	v, _ := strconv.ParseUint(value, 10, 64)
	return uint(v)
}

func parseFloat(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

func parseInt(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

func HandleCreateInstrumentFromAcceptance(facade *controller.ProcurementControllerFacade) {
	fmt.Println("List of Approved Accepted Requests:")

	PrintAcceptanceList(facade)

	acceptanceID := util.GetUintInput("Enter Acceptance Approval ID: ")

	util.ClearScreen()

	err := facade.Instrument.CreateInstrumentsFromAcceptance(acceptanceID)
	if err != nil {
		fmt.Println("Error creating instruments:", err)
		util.PressEnterToContinue()
		return
	}

	err = facade.Acceptance.UpdateStatusToImported(acceptanceID)
	if err != nil {
		fmt.Println("Error updating status to 'Imported':", err)
	} else {
		fmt.Printf("Instruments created and status updated to 'Imported' for Acceptance ID %d\n", acceptanceID)
	}

	util.PressEnterToContinue()
}

func PrintAcceptanceList(facade *controller.ProcurementControllerFacade) {
	acceptedRequests, err := facade.Acceptance.ListAllApprovals()
	if err != nil {
		fmt.Println("Failed to fetch acceptance approvals:", err)
		return
	}

	var approvedRequests []model.AcceptanceApproval
	for _, req := range acceptedRequests {
		if req.Status == model.AcceptanceStatusApproved {
			approvedRequests = append(approvedRequests, req)
		}
	}

	if len(approvedRequests) == 0 {
		fmt.Println("No approved acceptance requests found.")
		return
	}

	fmt.Println("---------------------------------------------------------------------------------------------")
	for _, a := range approvedRequests {
		approverID := "waiting"
		if a.ApproverID != nil && *a.ApproverID != 0 {
			approverID = fmt.Sprintf("%d", *a.ApproverID)
		}
		approvalTime := "N/A"
		if a.ApprovalTime != nil {
			approvalTime = a.ApprovalTime.Format("2006-01-02 15:04:05")
		}
		fmt.Printf("  ApprovalID: %d | ProcurementID: %d | Status: %s | Approver ID: %s | Approval Time: %s\n",
			a.AcceptanceApprovalID,
			a.ProcurementID,
			a.Status,
			approverID,
			approvalTime,
		)
	}
	fmt.Println("---------------------------------------------------------------------------------------------")
}
