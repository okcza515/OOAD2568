package validation

import (
	"regexp"
	"strconv"
	"sync"
)

const (
	emailRegexString   = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	numericRegexString = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberRegexString  = "^[0-9]+$"
)

func lazyRegexCompile(str string) func() *regexp.Regexp {
	var regex *regexp.Regexp
	var once sync.Once
	return func() *regexp.Regexp {
		once.Do(func() {
			regex = regexp.MustCompile(str)
		})
		return regex
	}
}

type Validator struct {
}

var validator *Validator

func NewValidator() *Validator {
	if validator == nil {
		validator = &Validator{}
	}
	return validator
}

func (v *Validator) IsNumberValid(num string) bool {
	if !v.IsStringNotEmpty(num) {
		return false
	}
	numericRegex := lazyRegexCompile(numericRegexString)()
	if numericRegex == nil {
		return false
	}
	return numericRegex.MatchString(num)
}

func (v *Validator) ParseNumber(num string) (float64, bool) {
	if !v.IsStringNotEmpty(num) {
		return 0, false
	}
	numericRegex := lazyRegexCompile(numericRegexString)()
	if numericRegex == nil || !numericRegex.MatchString(num) {
		return 0, false
	}
	value, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, false
	}
	return value, true
}

func (v *Validator) IsUintValid(num string) bool {
	if !v.IsStringNotEmpty(num) {
		return false
	}
	numberRegex := lazyRegexCompile(numberRegexString)()
	if numberRegex == nil {
		return false
	}
	return numberRegex.MatchString(num)
}

func (v *Validator) ParseUint(num string) (uint, bool) {
	if !v.IsStringNotEmpty(num) {
		return 0, false
	}
	numberRegex := lazyRegexCompile(numberRegexString)()
	if numberRegex == nil || !numberRegex.MatchString(num) {
		return 0, false
	}
	value, err := strconv.ParseUint(num, 10, 64)
	if err != nil {
		return 0, false
	}
	return uint(value), true
}

func (v *Validator) IsStringNotEmpty(str string) bool {
	return str != ""
}

func (v *Validator) IsEmailValid(email string) bool {
	if !v.IsStringNotEmpty(email) {
		return false
	}

	emailRegex := lazyRegexCompile(emailRegexString)()
	if emailRegex == nil {
		return false
	}
	return emailRegex.MatchString(email)
}

func (v *Validator) IsPhoneNumberValid(phone string) bool {
	if !v.IsStringNotEmpty(phone) {
		return false
	}
	if len(phone) != 10 {
		return false
	}
	if phone[0] != '0' {
		return false
	}
	for _, char := range phone {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
