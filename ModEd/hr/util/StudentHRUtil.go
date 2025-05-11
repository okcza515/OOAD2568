package util

import (
	"fmt"

	commonModel "ModEd/common/model"
)

func StatusToString(status commonModel.StudentStatus) string {
	switch status {
	case commonModel.ACTIVE:
		return "ACTIVE"
	case commonModel.GRADUATED:
		return "GRADUATED"
	case commonModel.DROP:
		return "DROP"
	default:
		return "ACTIVE"
	}
}

func StatusFromString(status string) (commonModel.StudentStatus, error) {
	switch status {
	case "ACTIVE", "Active":
		return commonModel.ACTIVE, nil
	case "GRADUATED", "Graduated":
		return commonModel.GRADUATED, nil
	case "DROP", "Drop":
		return commonModel.DROP, nil
	default:
		return commonModel.ACTIVE, fmt.Errorf("invalid status: %s (must be ACTIVE, GRADUATED, or DROP)", status)
	}
}

func ProgramTypeToString(program commonModel.ProgramType) string {
	switch program {
	case commonModel.REGULAR:
		return "Regular"
	case commonModel.INTERNATIONAL:
		return "International"
	default:
		return "Unknown"
	}
}

func ProgramTypeFromString(program string) (commonModel.ProgramType, error) {
	switch program {
	case "Regular":
		return commonModel.REGULAR, nil
	case "International":
		return commonModel.INTERNATIONAL, nil
	default:
		return commonModel.REGULAR, fmt.Errorf("invalid program type: %s (must be Regular or International)", program)
	}
}

func IfNotEmpty(newValue, fallback string) string {
	if newValue != "" {
		return newValue
	}
	return fallback
}

func IfNotZero(newValue, fallback int) int {
	if newValue != 0 {
		return newValue
	}
	return fallback
}
