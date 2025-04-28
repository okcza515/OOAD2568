package util

import (
	"flag"
	"fmt"
)

// ValidateRequiredFlags builds a chain of required flag validators and validates them.
// flags is a slice of flag names that are required.
func ValidateRequiredFlags(fs *flag.FlagSet, flags []string) error {
	if len(flags) == 0 {
		return nil
	}
	// Build the chain from the slice.
	head := NewRequiredFlagValidator(fs, flags[0])
	current := head
	for _, flagName := range flags[1:] {
		validator := NewRequiredFlagValidator(fs, flagName)
		current.SetNext(validator)
		current = validator
	}
	return head.Validate()
}

// ValidationChain provides a fluent interface for building a validation chain.
type ValidationChain struct {
	fs            *flag.FlagSet
	head          ValidatorHandler
	tail          ValidatorHandler
	creationError error // Stores errors during validator creation (e.g., bad regex)
}

// NewValidationChain creates a new validation chain builder.
func NewValidationChain(fs *flag.FlagSet) *ValidationChain {
	return &ValidationChain{fs: fs}
}

// append adds a validator to the end of the chain.
func (vc *ValidationChain) append(validator ValidatorHandler) {
	if vc.head == nil {
		vc.head = validator
		vc.tail = validator
	} else {
		vc.tail.SetNext(validator)
		vc.tail = validator
	}
}

// Required adds a required flag validation rule.
func (vc *ValidationChain) Required(flagName string) *ValidationChain {
	if vc.creationError != nil {
		return vc // Don't add more if an error already occurred
	}
	validator := NewRequiredFlagValidator(vc.fs, flagName)
	vc.append(validator)
	return vc
}

// Length adds a flag length validation rule.
func (vc *ValidationChain) Length(flagName string, requiredLen int) *ValidationChain {
	if vc.creationError != nil {
		return vc
	}
	validator := NewLengthFlagValidator(vc.fs, flagName, requiredLen)
	vc.append(validator)
	return vc
}

// Regex adds a regex pattern validation rule.
func (vc *ValidationChain) Regex(flagName, pattern string) *ValidationChain {
	if vc.creationError != nil {
		return vc
	}
	validator, err := NewRegexFlagValidator(vc.fs, flagName, pattern)
	if err != nil {
		vc.creationError = fmt.Errorf("failed to create regex validator for '%s': %w", flagName, err)
		return vc
	}
	vc.append(validator)
	return vc
}

func (vc *ValidationChain) AllowedValues(flagName string, allowed []string) *ValidationChain {
	if vc.creationError != nil {
		return vc
	}
	if len(allowed) == 0 {
		// Optional: Could set creationError or just skip adding the validator
		// For now, let's skip adding if 'allowed' is empty
		return vc
	}
	validator := NewAllowedValuesFlagValidator(vc.fs, flagName, allowed)
	vc.append(validator)
	return vc
}

// Validate executes the entire validation chain.
// It first checks for any errors during validator creation.
func (vc *ValidationChain) Validate() error {
	if vc.creationError != nil {
		return vc.creationError // Return creation error first
	}
	if vc.head == nil {
		return nil // No validation rules added
	}
	return vc.head.Validate() // Execute the chain
}
