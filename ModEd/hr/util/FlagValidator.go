package util

import (
	"flag"
	"fmt"
	"regexp"
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

type LengthFlagValidator struct {
	BaseValidator
	fs          *flag.FlagSet
	flagName    string
	requiredLen int
}

type RegexFlagValidator struct {
	BaseValidator
	fs       *flag.FlagSet
	flagName string
	pattern  *regexp.Regexp
}

// NewRequiredFlagValidator creates a new validator for a required flag.
func NewRequiredFlagValidator(fs *flag.FlagSet, flagName string) *RequiredFlagValidator {
	return &RequiredFlagValidator{
		fs:       fs,
		flagName: flagName,
	}
}

// NewLengthFlagValidator creates a new validator for checking the length of a flag.
func NewLengthFlagValidator(fs *flag.FlagSet, flagName string, requiredLen int) *LengthFlagValidator {
	return &LengthFlagValidator{
		fs:          fs,
		flagName:    flagName,
		requiredLen: requiredLen,
	}
}

func NewRegexFlagValidator(fs *flag.FlagSet, flagName, pattern string) (*RegexFlagValidator, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}
	return &RegexFlagValidator{
		fs:       fs,
		flagName: flagName,
		pattern:  re,
	}, nil
}

func (v *RequiredFlagValidator) Validate() error {
	f := v.fs.Lookup(v.flagName)
	if f == nil || f.Value.String() == "" {
		return fmt.Errorf("required flag '%s' is missing", v.flagName)
	}
	// Continue to next validator in the chain.
	return v.BaseValidator.Validate()
}

// Validate returns an error if the flag's value doesn’t have the required length.
func (v *LengthFlagValidator) Validate() error {
	f := v.fs.Lookup(v.flagName)
	if f == nil {
		return fmt.Errorf("flag '%s' not found", v.flagName)
	}

	value := f.Value.String()
	if len(value) != v.requiredLen {
		return fmt.Errorf("flag '%s' must be %d characters long", v.flagName, v.requiredLen)
	}

	return v.BaseValidator.Validate()
}

func (v *RegexFlagValidator) Validate() error {
	f := v.fs.Lookup(v.flagName)
	if f == nil {
		return fmt.Errorf("flag '%s' not found", v.flagName)
	}
	value := f.Value.String()
	if !v.pattern.MatchString(value) {
		return fmt.Errorf("flag '%s' does not match the required pattern", v.flagName)
	}
	return v.BaseValidator.Validate()
}
