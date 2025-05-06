package validation

// Wrote by MEP-1010

import (
	"fmt"
	"reflect"
	"strings"
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

type validation struct {
	Name string
	Tags []validationFunc
}

type ModelValidator struct {
	validator   *Validator
	validateTag map[ModelValidatorEnum]validationFunc
}

var modelValidatorInstance *ModelValidator

func NewModelValidator() *ModelValidator {
	if modelValidatorInstance == nil {
		modelValidatorInstance = &ModelValidator{
			validator:   NewValidator(),
			validateTag: make(map[ModelValidatorEnum]validationFunc),
		}
		modelValidatorInstance.validateTag[GORM_VALIDATE_NOT_NULL] = modelValidatorInstance.validator.IsStringNotEmpty

		modelValidatorInstance.validateTag[VALIDATE_STUDENT_ID] = modelValidatorInstance.validator.IsStudentID
		modelValidatorInstance.validateTag[VALIDATE_PHONE] = modelValidatorInstance.validator.IsPhoneNumberValid
		modelValidatorInstance.validateTag[VALIDATE_EMAIL] = modelValidatorInstance.validator.IsEmailValid
		modelValidatorInstance.validateTag[VALIDATE_DATETIME] = modelValidatorInstance.validator.IsDateTimeValid
	}
	return modelValidatorInstance

}

func (v ModelValidator) ModelValidate(model interface{}) error {
	st := reflect.TypeOf(model)
	val := reflect.ValueOf(model)

	if st.Kind() == reflect.Ptr {
		st = st.Elem()
		val = val.Elem()
	}

	if st.Kind() != reflect.Struct {
		return fmt.Errorf("input is not a struct")
	}

	validations, err := getValidator(st)
	if err != nil {
		return err
	}

	for _, validation := range validations {
		fieldValue := val.FieldByName(validation.Name)
		if !fieldValue.IsValid() {
			return fmt.Errorf("field %s does not exist in the model", validation.Name)
		}

		var value string
		switch fieldValue.Kind() {
		case reflect.String:
			value = fieldValue.Interface().(string)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = fmt.Sprintf("%d", fieldValue.Interface())
		case reflect.Float32, reflect.Float64:
			value = fmt.Sprintf("%f", fieldValue.Interface())
		default:
			value = fieldValue.Interface().(string)
		}

		for _, tag := range validation.Tags {
			if !tag(value) {
				return fmt.Errorf("validation failed for field %s with value %s in tags", validation.Name, value)
			}
		}
	}

	return nil
}

func getValidator(st reflect.Type) ([]validation, error) {
	var validations []validation

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		var tags []validationFunc

		tags = append(tags, getStructFieldTags(field, CONST_GORM_STRUCT_FIELD_TAGS)...)
		tags = append(tags, getStructFieldTags(field, CONST_VALIDATION_STRUCT_FIELD_TAGS)...)

		if len(tags) != 0 {
			validations = append(validations, validation{
				Name: field.Name,
				Tags: tags,
			})
		}
	}
	return validations, nil
}

func getStructFieldTags(field reflect.StructField, fieldName string) []validationFunc {
	var tags []validationFunc
	if tag, ok := field.Tag.Lookup(fieldName); ok {
		tagsValue := strings.Split(tag, ",")
		for _, tagValue := range tagsValue {
			tagValue = strings.TrimSpace(tagValue)

			if tagValue == "" {
				continue
			}

			validationFunc := getValidationFunc(ModelValidatorEnum(tagValue))
			if validationFunc == nil {
				continue
			}
			tags = append(tags, validationFunc)
		}
	}

	return tags
}

func getValidationFunc(tag ModelValidatorEnum) validationFunc {
	return modelValidatorInstance.validateTag[tag]
}
