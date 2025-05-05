package model

import "fmt"

type DepartmentPosition int

const (
	NONE_POSITION DepartmentPosition = iota
	HEAD
	DEPUTY
	SECRETARY
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
