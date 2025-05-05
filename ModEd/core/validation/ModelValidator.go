package validation

// Wrote by MEP-1010

import (
	"fmt"
	"reflect"
	"strings"
)

type ModelValidtorEnum string

const (
	VALIDATE_REQUIRED ModelValidtorEnum = "required"
	VALIDATE_EMAIL    ModelValidtorEnum = "email"
	VALIDATE_NUMERIC  ModelValidtorEnum = "numeric"
	VALIDATE_UINT     ModelValidtorEnum = "uint"
	VALIDATE_DATE     ModelValidtorEnum = "date"
	VALIDATE_PHONE    ModelValidtorEnum = "phone"
)

type validationFunc func(value string) bool

type validation struct {
	Name string
	Tags []validationFunc
}

type ModelValidator struct {
	validator   *Validator
	validateTag map[ModelValidtorEnum]validationFunc
}

var modelValidatorInstance *ModelValidator

func NewModelValidator() *ModelValidator {
	if modelValidatorInstance == nil {
		modelValidatorInstance = &ModelValidator{
			validator:   NewValidator(),
			validateTag: make(map[ModelValidtorEnum]validationFunc),
		}
		modelValidatorInstance.validateTag[VALIDATE_PHONE] = modelValidatorInstance.validator.IsPhoneNumberValid
		modelValidatorInstance.validateTag[VALIDATE_EMAIL] = modelValidatorInstance.validator.IsEmailValid
		modelValidatorInstance.validateTag[VALIDATE_REQUIRED] = modelValidatorInstance.validator.IsStringNotEmpty
		modelValidatorInstance.validateTag[VALIDATE_NUMERIC] = modelValidatorInstance.validator.IsNumberValid
		modelValidatorInstance.validateTag[VALIDATE_UINT] = modelValidatorInstance.validator.IsUintValid
		// modelValidatorInstance.validateTag[VALIDATE_DATE] = modelValidatorInstance.validator.IsDate
		// modelValidatorInstance.validateTag[VALIDATE_TIME] = modelValidatorInstance.validator.IsTime
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

	validations, err := getValidatorFields(st)
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
			return fmt.Errorf("field %s is not of a supported type", validation.Name)
		}

		for _, tag := range validation.Tags {
			if !tag(value) {
				return fmt.Errorf("validation failed for field %s with value %s in tags %s", validation.Name, value)
			}
		}
	}
	return nil
}

func getValidatorFields(st reflect.Type) ([]validation, error) {
	var validations []validation

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if tag, ok := field.Tag.Lookup("validation"); ok {
			// Split the tag value by comma to get individual validation tags
			tags := []validationFunc{}
			tagsValue := strings.Split(tag, ",")
			for _, tagValue := range tagsValue {
				tagValue = strings.TrimSpace(tagValue)

				if tagValue == "" {
					continue
				}
				// Get the validation function for the tag
				validationFunc := getValidationFunc(ModelValidtorEnum(tagValue))
				if validationFunc == nil {
					return nil, fmt.Errorf("validation function not found for tag: %s", tagValue)
				}
				// Append the validation function to the tags slice
				tags = append(tags, validationFunc)
			}

			validations = append(validations, validation{
				Name: field.Name,
				Tags: tags,
			})
		}
	}

	return validations, nil
}

func getValidationFunc(tag ModelValidtorEnum) validationFunc {
	return modelValidatorInstance.validateTag[tag]
}
