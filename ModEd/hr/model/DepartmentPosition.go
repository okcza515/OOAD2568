package model

import "fmt"

type DepartmentPosition int

const (
	HEAD DepartmentPosition = iota
	DEPUTY
	SECRETARY
	NONE_POSITION
)

func ParseDepartmentPosition(posStr string) (DepartmentPosition, error) {
	switch posStr {
	case "head", "HEAD":
		return HEAD, nil
	case "deputy", "DEPUTY":
		return DEPUTY, nil
	case "secret":
		return SECRETARY, nil
	case "none", "NONE_POSITION":
		return NONE_POSITION, nil
	default:
		return NONE_POSITION, fmt.Errorf("invalid department position: %s", posStr)
	}
}
