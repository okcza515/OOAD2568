package validation

// Wrote by MEP-1010, MEP-1004, MEP-1002

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type ModelValidatorEnum string

const (
	VALIDATE_EMAIL         ModelValidatorEnum = "email"
	VALIDATE_DATETIME      ModelValidatorEnum = "datetime"
	VALIDATE_PHONE         ModelValidatorEnum = "phone"
	VALIDATE_STUDENT_ID    ModelValidatorEnum = "studentId"
	GORM_VALIDATE_NOT_NULL ModelValidatorEnum = "not null"

	CONST_GORM_STRUCT_FIELD_TAGS       = "gorm"
	CONST_VALIDATION_STRUCT_FIELD_TAGS = "validation"
)

type validationFunc func(value string) bool

// fieldValidationRule stores the rules for a single field.
type fieldValidationRule struct {
	FieldName      string
	ValidationTags []struct {
		Tag  ModelValidatorEnum
		Func validationFunc
	}
}

type ModelValidator struct {
	validator   *Validator
	validateTag map[ModelValidatorEnum]validationFunc
}

var (
	modelValidatorInstance *ModelValidator
	modelValidatorOnce     sync.Once
)

func NewModelValidator() *ModelValidator {
	modelValidatorOnce.Do(func() {
		validator := NewValidator()
		validateTagMap := make(map[ModelValidatorEnum]validationFunc)

		validateTagMap[GORM_VALIDATE_NOT_NULL] = validator.IsStringNotEmpty
		validateTagMap[VALIDATE_STUDENT_ID] = validator.IsStudentID
		validateTagMap[VALIDATE_PHONE] = validator.IsPhoneNumberValid
		validateTagMap[VALIDATE_EMAIL] = validator.IsEmailValid
		validateTagMap[VALIDATE_DATETIME] = validator.IsDateTimeValid

		modelValidatorInstance = &ModelValidator{
			validator:   validator,
			validateTag: validateTagMap,
		}
	})
	return modelValidatorInstance
}

func (mv *ModelValidator) ModelValidate(model interface{}) error {
	st := reflect.TypeOf(model)
	val := reflect.ValueOf(model)

	if st.Kind() == reflect.Ptr {
		st = st.Elem()
		val = val.Elem()
	}

	if st.Kind() != reflect.Struct {
		return fmt.Errorf("input is not a struct, got %s", st.Kind())
	}

	fieldRules, err := mv.getFieldValidationRules(st)
	if err != nil {
		return fmt.Errorf("error getting validation rules: %w", err)
	}

	for _, rule := range fieldRules {
		fieldValue := val.FieldByName(rule.FieldName)
		if !fieldValue.IsValid() {
			return fmt.Errorf("field '%s' configured for validation does not exist in the model", rule.FieldName)
		}

		// Convert field value to string for validation
		// %v is a general verb that works for many types.
		// If specific string formatting is needed for certain types,
		// the switch statement can be reintroduced or refined.
		valueStr := fmt.Sprintf("%v", fieldValue.Interface())

		for _, tagValidation := range rule.ValidationTags {
			if !tagValidation.Func(valueStr) {
				return fmt.Errorf("validation failed for field '%s' with value '%s' for rule '%s'", rule.FieldName, valueStr, tagValidation.Tag)
			}
		}
	}

	return nil
}

// getFieldValidationRules extracts validation rules for each field of a struct type.
func (mv *ModelValidator) getFieldValidationRules(st reflect.Type) ([]fieldValidationRule, error) {
	var rules []fieldValidationRule

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		var currentFieldTags []struct {
			Tag  ModelValidatorEnum
			Func validationFunc
		}

		// Parse gorm tags
		gormTags := mv.parseStructFieldTags(field, CONST_GORM_STRUCT_FIELD_TAGS)
		currentFieldTags = append(currentFieldTags, gormTags...)

		// Parse validation tags
		validationTags := mv.parseStructFieldTags(field, CONST_VALIDATION_STRUCT_FIELD_TAGS)
		currentFieldTags = append(currentFieldTags, validationTags...)

		if len(currentFieldTags) > 0 {
			rules = append(rules, fieldValidationRule{
				FieldName:      field.Name,
				ValidationTags: currentFieldTags,
			})
		}
	}
	return rules, nil
}

// parseStructFieldTags parses a specific tag (like "gorm" or "validation") from a struct field.
func (mv *ModelValidator) parseStructFieldTags(field reflect.StructField, tagName string) []struct {
	Tag  ModelValidatorEnum
	Func validationFunc
} {
	var parsedTags []struct {
		Tag  ModelValidatorEnum
		Func validationFunc
	}
	tagValueStr, ok := field.Tag.Lookup(tagName)
	if !ok {
		return parsedTags // No tag found
	}

	tagValues := strings.Split(tagValueStr, ",")
	for _, tv := range tagValues {
		trimmedTag := ModelValidatorEnum(strings.TrimSpace(tv))
		if trimmedTag == "" {
			continue
		}

		fn := mv.getValidationFuncByTag(trimmedTag)
		if fn != nil {
			parsedTags = append(parsedTags, struct {
				Tag  ModelValidatorEnum
				Func validationFunc
			}{Tag: trimmedTag, Func: fn})
		}
	}
	return parsedTags
}

// getValidationFuncByTag retrieves the validation function associated with a ModelValidatorEnum tag.
func (mv *ModelValidator) getValidationFuncByTag(tag ModelValidatorEnum) validationFunc {
	// mv.validateTag is initialized in NewModelValidator
	return mv.validateTag[tag]
}
