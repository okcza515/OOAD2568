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
		printInstrumentList(facade)

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

func printInstrumentList(facade *controller.ProcurementControllerFacade) {
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
	for _, inst := range instruments {
		fmt.Println(inst)
	}
}

func HandleInstrumentDetails(facade *controller.ProcurementControllerFacade) {
	util.ClearScreen()
	id := util.GetUintInput("Enter Instrument ID: ")
	instrument, err := facade.Instrument.RetrieveByID(id)
	if err != nil {
		fmt.Println("Failed to retrieve instrument:", err)
	} else {
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
	util.ClearScreen()
	fmt.Println("Create Instruments from Accepted Request")
	acceptanceID := util.GetUintInput("Enter Acceptance Approval ID: ")

	err := facade.Instrument.CreateInstrumentsFromAcceptance(acceptanceID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Instruments created successfully for Acceptance ID %d\n", acceptanceID)
	}
	util.PressEnterToContinue()
}
