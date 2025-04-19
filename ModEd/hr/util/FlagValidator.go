package util

import (
	"flag"
	"fmt"
)

// ValidatorHandler defines the interface for flag validation handlers.
type ValidatorHandler interface {
	SetNext(handler ValidatorHandler)
	Validate() error
}

// BaseValidator provides default behavior to pass to the next handler.
type BaseValidator struct {
	next ValidatorHandler
}

func (b *BaseValidator) SetNext(handler ValidatorHandler) {
	b.next = handler
}

func (b *BaseValidator) Validate() error {
	if b.next != nil {
		return b.next.Validate()
	}
	return nil
}

// RequiredFlagValidator checks that a specific flag is provided.
type RequiredFlagValidator struct {
	BaseValidator
	fs       *flag.FlagSet
	flagName string
}

// NewRequiredFlagValidator creates a new validator for a required flag.
func NewRequiredFlagValidator(fs *flag.FlagSet, flagName string) *RequiredFlagValidator {
	return &RequiredFlagValidator{
		fs:       fs,
		flagName: flagName,
	}
}

func (v *RequiredFlagValidator) Validate() error {
	f := v.fs.Lookup(v.flagName)
	if f == nil || f.Value.String() == "" {
		return fmt.Errorf("required flag '%s' is missing", v.flagName)
	}
	// Continue to next validator in the chain.
	return v.BaseValidator.Validate()
}
