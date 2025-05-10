package model

import (
	"fmt"
	"strings"
)

type AcademicPosition int

const (
	NONE AcademicPosition = iota
	ASSISTANT_PROF
	ASSOCIATE_PROF
	PROFESSOR
)

var academicPositionMap = map[string]AcademicPosition{
	"assistant": ASSISTANT_PROF,
	"associate": ASSOCIATE_PROF,
	"professor": PROFESSOR,
	"none":      NONE,
}

func ParseAcademicPosition(posStr string) (AcademicPosition, error) {
	position, ok := academicPositionMap[strings.ToLower(posStr)]
	if !ok {
		return NONE, fmt.Errorf("invalid academic position: %s", posStr)
	}
	return position, nil
}
