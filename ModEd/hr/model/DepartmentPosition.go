package model

type DepartmentPosition int

const (
	HEAD DepartmentPosition = iota
	DEPUTY
	SECRETARY
	NONE_POSITION
)
