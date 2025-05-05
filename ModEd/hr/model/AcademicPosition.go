package model

import "fmt"

type AcademicPosition int

const (
	NONE AcademicPosition = iota
	ASSISTANT_PROF
	ASSOCIATE_PROF
	PROFESSOR
)

func ParseAcademicPosition(posStr string) (AcademicPosition, error) {
	switch posStr {
	case "assistant", "ASSISTANT_PROF":
		return ASSISTANT_PROF, nil
	case "associate", "ASSOCIATE_PROF":
		return ASSOCIATE_PROF, nil
	case "professor", "PROFESSOR":
		return PROFESSOR, nil
	case "none", "NONE":
		return NONE, nil
	default:
		return NONE, fmt.Errorf("invalid academic position: %s", posStr)
	}
}
