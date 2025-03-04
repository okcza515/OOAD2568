package model

type StudentStatus int

const (
	ACTIVE StudentStatus = iota
	GRADUATED
	DROP
)
