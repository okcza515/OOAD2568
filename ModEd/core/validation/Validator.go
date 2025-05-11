package validation

// Wrote by MEP-1010, MEP-1004, MEP-1002

import (
	"regexp"
	"strconv"
	"sync"
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

// Pre-initialize regexes to avoid repeated compilation
var (
	getEmailRegex       = lazyRegexCompile("^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
	getNumericRegex     = lazyRegexCompile("^[-+]?[0-9]+(?:\\.[0-9]+)?$")
	getNumberRegex      = lazyRegexCompile("^[0-9]+$")
	getDateTimeRegex    = lazyRegexCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)
	getDateRegex        = lazyRegexCompile(`^\d{4}-\d{2}-\d{2}$`)
	getStudentCodeRegex = lazyRegexCompile(`^\d{11}$`)
)

type Validator struct {
}

var (
	validatorInstance *Validator
	validatorOnce     sync.Once
)

func NewValidator() *Validator {
	validatorOnce.Do(func() {
		validatorInstance = &Validator{}
	})
	return validatorInstance
}

func (v *Validator) IsNumberValid(num string) bool {
	if !v.IsStringNotEmpty(num) {
		return false
	}
	return getNumericRegex().MatchString(num)
}

func (v *Validator) ParseNumber(num string) (float64, bool) {
	if !v.IsStringNotEmpty(num) {
		return 0, false
	}
	if !getNumericRegex().MatchString(num) {
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
	return getNumberRegex().MatchString(num)
}

func (v *Validator) ParseUint(num string) (uint, bool) {
	if !v.IsStringNotEmpty(num) {
		return 0, false
	}
	if !getNumberRegex().MatchString(num) {
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
	return getEmailRegex().MatchString(email)
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

func (v *Validator) IsDateTimeValid(dateTime string) bool {
	if !v.IsStringNotEmpty(dateTime) {
		return false
	}
	return getDateTimeRegex().MatchString(dateTime)
}

func (v *Validator) IsDateValid(date string) bool {
	if !v.IsStringNotEmpty(date) {
		return false
	}
	return getDateRegex().MatchString(date)
}

func (v *Validator) IsStudentID(studentId string) bool {
	if !v.IsStringNotEmpty(studentId) {
		return false
	}
	return getStudentCodeRegex().MatchString(studentId)
}

func (v *Validator) IsValueAllowed(value string, allowedValues []string) bool {
	for _, allowed := range allowedValues {
		if value == allowed {
			return true
		}
	}
	return false
}
