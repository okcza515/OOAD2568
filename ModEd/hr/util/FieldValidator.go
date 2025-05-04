package util

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

type FieldValidator struct {
	fs             *flag.FlagSet
	flagName       string
	rules          []func(value string) error // Slice of validation functions for this field
	creationErrors []error                    // Errors during rule setup (e.g., bad regex)
}

// addRule adds a validation function to the field's rule list.
func (fv *FieldValidator) addRule(rule func(value string) error) {
	fv.rules = append(fv.rules, rule)
}

// addCreationError records an error encountered during rule definition.
func (fv *FieldValidator) addCreationError(err error) {
	fv.creationErrors = append(fv.creationErrors, err)
}

// Required marks the field as required.
func (fv *FieldValidator) Required() *FieldValidator {
	fv.addRule(func(value string) error {
		f := fv.fs.Lookup(fv.flagName)
		if f == nil || f.Value.String() == "" {
			if defVal, ok := f.Value.(flag.Getter); ok && defVal.String() == "" {
				return fmt.Errorf("required flag '%s' is missing or empty", fv.flagName)
			}
		}
		return nil
	})
	return fv
}

// Length adds a validation rule for exact string length.
func (fv *FieldValidator) Length(requiredLen int) *FieldValidator {
	fv.addRule(func(value string) error {
		// Only apply length check if value is not empty (let Required handle empty)
		if value != "" && len(value) != requiredLen {
			return fmt.Errorf("flag '%s' must be %d characters long, but got %d", fv.flagName, requiredLen, len(value))
		}
		return nil
	})
	return fv
}

// Regex adds a validation rule using a regular expression.
func (fv *FieldValidator) Regex(pattern string) *FieldValidator {
	re, err := regexp.Compile(pattern)
	if err != nil {
		fv.addCreationError(fmt.Errorf("invalid regex pattern for '%s': %w", fv.flagName, err))
		fv.addRule(func(value string) error { return nil })
		return fv
	}
	fv.addRule(func(value string) error {
		// Only apply regex check if value is not empty
		if value != "" && !re.MatchString(value) {
			return fmt.Errorf("flag '%s' value '%s' does not match pattern '%s'", fv.flagName, value, pattern)
		}
		return nil
	})
	return fv
}

// AllowedValues adds a validation rule ensuring the value is one of the allowed strings (case-insensitive).
func (fv *FieldValidator) AllowedValues(allowed []string) *FieldValidator {
	if len(allowed) == 0 {
		fv.addCreationError(fmt.Errorf("AllowedValues for '%s' called with empty list", fv.flagName))
		fv.addRule(func(value string) error { return nil }) // Add dummy rule
		return fv
	}

	allowedMap := make(map[string]struct{}, len(allowed))
	for _, val := range allowed {
		allowedMap[strings.ToLower(val)] = struct{}{}
	}

	fv.addRule(func(value string) error {
		if value == "" { // Skip check if empty (let Required handle)
			return nil
		}
		if _, ok := allowedMap[strings.ToLower(value)]; !ok {
			return fmt.Errorf("flag '%s' has invalid value '%s'. Allowed values are: %v", fv.flagName, value, allowed)
		}
		return nil
	})
	return fv
}

func (fv *FieldValidator) IsStudentID() *FieldValidator {
	// Student ID is 11 digits long
	return fv.Length(11).Regex(`^[0-9]{11}$`)
}

func (fv *FieldValidator) IsInstructorID() *FieldValidator {
	return fv.Required()
}

func (fv *FieldValidator) IsDate() *FieldValidator {
	// dd-mm-yyyy format
	return fv.Regex(`^\d{2}-\d{2}-\d{4}$`)
}

func (fv *FieldValidator) IsEmail() *FieldValidator {
	// Simple email regex
	return fv.Regex(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
}

func (fv *FieldValidator) IsDigit() *FieldValidator {
	return fv.Regex(`^\d+$`)
}

// validateField applies all registered rules for this specific field.
func (fv *FieldValidator) validateField() []error {
	var errors []error
	errors = append(errors, fv.creationErrors...)

	if len(fv.creationErrors) > 0 {
		return errors
	}

	f := fv.fs.Lookup(fv.flagName)
	if f == nil {
		errors = append(errors, fmt.Errorf("validation rule defined for non-existent flag '%s'", fv.flagName))
		return errors
	}
	value := f.Value.String()

	isEmpty := value == ""
	if defVal, ok := f.Value.(flag.Getter); ok {
		isEmpty = isEmpty && defVal.String() == ""
	}

	for i, rule := range fv.rules {
		isLikelyRequiredRule := i == 0

		if !isEmpty || isLikelyRequiredRule {
			if err := rule(value); err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}
