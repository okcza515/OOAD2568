package model

import (
	"fmt"
	"strings"
)

type DepartmentPosition int

const (
	NONE_POSITION DepartmentPosition = iota
	HEAD
	DEPUTY
	SECRETARY
)

var departmentPositionMap = map[string]DepartmentPosition{
	"head":      HEAD,
	"deputy":    DEPUTY,
	"secretary": SECRETARY,
	"none":      NONE_POSITION,
}

func ParseDepartmentPosition(posStr string) (DepartmentPosition, error) {
	position, ok := departmentPositionMap[strings.ToLower(posStr)]
	if !ok {
		return NONE_POSITION, fmt.Errorf("invalid department position: %s", posStr)
	}
	return position, nil
}
