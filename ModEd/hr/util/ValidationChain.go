package util

import (
	"flag"
	"fmt"
	"strings"
)

type ValidationChain struct {
	fs     *flag.FlagSet
	fields map[string]*FieldValidator
}

func NewValidationChain(fs *flag.FlagSet) *ValidationChain {
	return &ValidationChain{
		fs:     fs,
		fields: make(map[string]*FieldValidator),
	}
}

func (vc *ValidationChain) Field(flagName string) *FieldValidator {
	if fv, exists := vc.fields[flagName]; exists {
		return fv
	}

	fv := &FieldValidator{
		fs:       vc.fs,
		flagName: flagName,
		rules:    make([]func(value string) error, 0),
	}

	// Check if the flag exists in the FlagSet
	if vc.fs.Lookup(flagName) == nil {
		// Record error within the FieldValidator itself
		fv.addCreationError(fmt.Errorf("attempted to validate undefined flag '%s'", flagName))
	}

	vc.fields[flagName] = fv
	return fv
}

// Validate executes all validation rules defined for all fields.
// It collects and returns all validation errors found.
func (vc *ValidationChain) Validate() error {
	var allErrors []string

	// Ensure flags are parsed before validation
	if !vc.fs.Parsed() {
		return fmt.Errorf("flagset must be parsed before calling Validate()")
	}

	for _, fieldValidator := range vc.fields {
		// Skip validation if the flag itself doesn't exist (error recorded in Field())
		if fieldValidator.fs.Lookup(fieldValidator.flagName) == nil && len(fieldValidator.creationErrors) > 0 {
			// Add only the creation error about the missing flag
			for _, err := range fieldValidator.creationErrors {
				allErrors = append(allErrors, err.Error())
			}
			continue
		}

		fieldErrors := fieldValidator.validateField()
		for _, err := range fieldErrors {
			allErrors = append(allErrors, err.Error())
		}
	}

	if len(allErrors) > 0 {
		return fmt.Errorf("validation failed:\n - %s", strings.Join(allErrors, "\n - "))
	}

	return nil
}
