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
	case "ACTIVE":
		return commonModel.ACTIVE, nil
	case "GRADUATED":
		return commonModel.GRADUATED, nil
	case "DROP":
		return commonModel.DROP, nil
	default:
		return commonModel.ACTIVE, fmt.Errorf("invalid status: %s (must be ACTIVE, GRADUATED, or DROP)", status)
	}
}
