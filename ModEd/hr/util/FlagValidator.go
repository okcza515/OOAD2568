package util

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
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

type AllowedValuesFlagValidator struct {
	BaseValidator
	fs           *flag.FlagSet
	flagName     string
	allowed      map[string]struct{}
	allowedSlice []string
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

func NewAllowedValuesFlagValidator(fs *flag.FlagSet, flagName string, allowed []string) *AllowedValuesFlagValidator {
	allowedMap := make(map[string]struct{}, len(allowed))
	for _, val := range allowed {
		// Store lowercase versions for case-insensitive comparison
		allowedMap[strings.ToLower(val)] = struct{}{}
	}
	return &AllowedValuesFlagValidator{
		fs:           fs,
		flagName:     flagName,
		allowed:      allowedMap,
		allowedSlice: allowed, // Store original for error message
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

func (v *AllowedValuesFlagValidator) Validate() error {
	f := v.fs.Lookup(v.flagName)
	if f == nil {
		// This might indicate a programming error (checking a non-existent flag)
		// Or it could be handled by a RequiredFlagValidator earlier in the chain.
		// Depending on desired behavior, you might return an error or let the chain continue.
		// Let's assume RequiredFlagValidator handles missing flags.
		return v.BaseValidator.Validate() // Proceed if flag doesn't exist (maybe it wasn't required)
	}

	value := f.Value.String()
	if value == "" {
		// If the value is empty, let RequiredFlagValidator handle it if needed.
		return v.BaseValidator.Validate()
	}

	// Check against the allowed values (case-insensitive)
	if _, ok := v.allowed[strings.ToLower(value)]; !ok {
		return fmt.Errorf("flag '%s' has invalid value '%s'. Allowed values are: %v", v.flagName, value, v.allowedSlice)
	}

	// Value is valid, continue chain
	return v.BaseValidator.Validate()
}
