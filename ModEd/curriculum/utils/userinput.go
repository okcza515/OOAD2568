package utils

import "fmt"

type UserInputStep interface {
	Prompt() string
	FieldName() string
	Validate(input string) (any, bool)
}

type UintInputStep struct {
	PromptText    string
	FieldNameText string
}

func (s UintInputStep) Prompt() string    { return s.PromptText }
func (s UintInputStep) FieldName() string { return s.FieldNameText }
func (s UintInputStep) Validate(input string) (any, bool) {
	if !NewValidator().IsStringNotEmpty(input) {
		return nil, false
	}
	if val, ok := NewValidator().ParseUint(input); ok {
		return val, true
	}
	return nil, false
}

type StringInputStep struct {
	PromptText    string
	FieldNameText string
}

func (s StringInputStep) Prompt() string    { return s.PromptText }
func (s StringInputStep) FieldName() string { return s.FieldNameText }
func (s StringInputStep) Validate(input string) (any, bool) {
	if NewValidator().IsStringNotEmpty(input) {
		return input, true
	}
	return nil, false
}

type EmailInputStep struct {
	PromptText    string
	FieldNameText string
}

func (s EmailInputStep) Prompt() string    { return s.PromptText }
func (s EmailInputStep) FieldName() string { return s.FieldNameText }
func (s EmailInputStep) Validate(input string) (any, bool) {
	if NewValidator().IsEmailValid(input) {
		return input, true
	}
	return nil, false
}

func ExecuteUserInputStep(step UserInputStep) any {
	for {
		temp := GetUserInput(step.Prompt())
		if value, ok := step.Validate(temp); ok {
			return value
		}
		fmt.Printf("Invalid or empty %s\n", step.FieldName())
	}
}
