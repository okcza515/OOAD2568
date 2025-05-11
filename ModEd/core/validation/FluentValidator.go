package validation

// Wrote by MEP-1010, MEP-1004, MEP-1002

import (
	"fmt"
	"strings"
)

type InputGetterFunc func(prompt string) string

type FieldConfig struct {
	Name   string
	Prompt string
}

type FieldValidator struct {
	chain         *ValidationChain
	fieldName     string
	prompt        string
	value         string
	rules         []func(string) string
	isRequired    bool
	inputObtained bool
}

type ValidationChain struct {
	validator   *Validator
	inputGetter InputGetterFunc
}

func NewValidationChain(inputGetter InputGetterFunc) *ValidationChain {
	return &ValidationChain{
		validator:   NewValidator(),
		inputGetter: inputGetter,
	}
}

func (vc *ValidationChain) Field(config FieldConfig) *FieldValidator {
	return &FieldValidator{
		chain:      vc,
		fieldName:  config.Name,
		prompt:     config.Prompt,
		rules:      make([]func(string) string, 0),
		isRequired: false,
	}
}

func (fv *FieldValidator) obtainInput() {
	if !fv.inputObtained {
		fv.value = fv.chain.inputGetter(fv.prompt)
		fv.inputObtained = true
	}
}

func (fv *FieldValidator) Required() *FieldValidator {
	fv.isRequired = true
	fv.rules = append([]func(string) string{
		func(val string) string {
			if val == "" {
				return fmt.Sprintf("field '%s' is required", fv.fieldName)
			}
			return ""
		},
	}, fv.rules...)
	return fv
}

func (fv *FieldValidator) IsStudentCode() *FieldValidator {
	fv.rules = append(fv.rules, func(val string) string {
		if !fv.chain.validator.IsStudentID(val) {
			return fmt.Sprintf("field '%s' ('%s') must be a valid student ID (11 digits)", fv.fieldName, val)
		}
		return ""
	})
	return fv
}

func (fv *FieldValidator) IsEmail() *FieldValidator {
	fv.rules = append(fv.rules, func(val string) string {
		if !fv.chain.validator.IsEmailValid(val) {
			return fmt.Sprintf("field '%s' ('%s') must be a valid email", fv.fieldName, val)
		}
		return ""
	})
	return fv
}

func (fv *FieldValidator) IsDateTime() *FieldValidator {
	fv.rules = append(fv.rules, func(val string) string {
		if !fv.chain.validator.IsDateTimeValid(val) {
			return fmt.Sprintf("field '%s' ('%s') must be a valid date/time (YYYY-MM-DD HH:MM:SS)", fv.fieldName, val)
		}
		return ""
	})
	return fv
}

func (fv *FieldValidator) IsDate() *FieldValidator {
	fv.rules = append(fv.rules, func(val string) string {
		if !fv.chain.validator.IsDateValid(val) {
			return fmt.Sprintf("field '%s' ('%s') must be a valid date (YYYY-MM-DD)", fv.fieldName, val)
		}
		return ""
	})
	return fv
}

func (fv *FieldValidator) Length(length int) *FieldValidator {
	fv.rules = append(fv.rules, func(val string) string {
		if len(val) != length {
			return fmt.Sprintf("field '%s' ('%s') must have a length of %d", fv.fieldName, val, length)
		}
		return ""
	})
	return fv
}

func (fv *FieldValidator) IsPhoneNumber() *FieldValidator {
	fv.rules = append(fv.rules, func(val string) string {
		if !fv.chain.validator.IsPhoneNumberValid(val) {
			return fmt.Sprintf("field '%s' ('%s') must be a valid phone number (10 digits, starts with 0)", fv.fieldName, val)
		}
		return ""
	})
	return fv
}

func (fv *FieldValidator) IsAllDigits() *FieldValidator {
	fv.rules = append(fv.rules, func(val string) string {
		if !fv.chain.validator.IsUintValid(val) {
			return fmt.Sprintf("field '%s' ('%s') must consist of digits only", fv.fieldName, val)
		}
		return ""
	})
	return fv
}

func (fv *FieldValidator) AllowedValues(allowed []string) *FieldValidator {
	fv.rules = append(fv.rules, func(val string) string {
		if !fv.chain.validator.IsValueAllowed(val, allowed) {
			return fmt.Sprintf("field '%s' ('%s') must be one of: %s", fv.fieldName, val, strings.Join(allowed, ", "))
		}
		return ""
	})
	return fv
}

func (fv *FieldValidator) GetInput() string {
	for {
		fv.obtainInput()

		// If the field is not required and the input is empty, it's considered valid without running further rules.
		if !fv.isRequired && fv.value == "" {
			return fv.value
		}

		var allErrors []string

		for _, rule := range fv.rules {
			errMsg := rule(fv.value)
			if errMsg != "" {
				allErrors = append(allErrors, errMsg)
			}
		}

		if len(allErrors) == 0 {
			return fv.value
		}

		fmt.Printf("Input errors for '%s':\n", fv.fieldName)
		for _, e := range allErrors {
			fmt.Printf("- %s\n", e)
		}
		fmt.Printf("Please try again.\n")
		fv.inputObtained = false
	}
}

func (fv *FieldValidator) GetParsedNumber() float64 {
	for {
		val := fv.GetInput()
		if parsed, ok := fv.chain.validator.ParseNumber(val); ok {
			return parsed
		}
		fmt.Printf("Input '%s' is not a valid number. Please try again.\n", val)
		fv.inputObtained = false
	}
}

func (fv *FieldValidator) GetParsedUint() uint {
	for {
		val := fv.GetInput()
		if parsed, ok := fv.chain.validator.ParseUint(val); ok {
			return parsed
		}
		fmt.Printf("Input '%s' is not a valid unsigned integer. Please try again.\n", val)
		fv.inputObtained = false
	}
}
