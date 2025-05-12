package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"fmt"
	"strings"
)

type ControllerExportFunc func(filePath string) error

type ExportHandler struct {
	controllerFunc ControllerExportFunc
}

func NewExportHandlerStrategy(controllerFunc ControllerExportFunc) *ExportHandler {
	return &ExportHandler{controllerFunc: controllerFunc}
}

func (handler ExportHandler) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	filePath := validator.Field(validation.FieldConfig{
		Name:   "path",
		Prompt: "Enter the full file path for export (e.g., export_data.csv or path/to/data.json): ",
	}).Required().GetInput()

	if idx := strings.LastIndex(filePath, "."); idx == -1 || idx == len(filePath)-1 {
		fmt.Println("Warning: File path does not appear to have a valid extension (e.g., .csv, .json).")
	}

	if err := handler.controllerFunc(filePath); err != nil {
		return fmt.Errorf("failed to export file: %v", err)
	}

	fileExt := ""
	if dotIndex := strings.LastIndex(filePath, "."); dotIndex != -1 && dotIndex < len(filePath)-1 {
		fileExt = strings.ToUpper(filePath[dotIndex+1:])
	}

	if fileExt != "" {
		fmt.Printf("Export to %s successful: %s\n", fileExt, filePath)
	} else {
		fmt.Printf("Export successful: %s\n", filePath)
	}
	return nil
}
