package util

import "flag"

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
