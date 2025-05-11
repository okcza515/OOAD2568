package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"fmt"
)

type ControllerImportFunc func(filePath string) error

type ImportHandler struct{ controllerFunc ControllerImportFunc }

func NewImportHandlerStrategy(controllerFunc ControllerImportFunc) *ImportHandler {
	return &ImportHandler{controllerFunc: controllerFunc}
}

func (handler ImportHandler) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	path := validator.Field(validation.FieldConfig{Name: "path", Prompt: "Enter the path to the CSV file: "}).Required().GetInput()

	if err := handler.controllerFunc(path); err != nil {
		return fmt.Errorf("failed to import file: %v", err)
	}

	fmt.Println("Import successful")
	return nil
}
