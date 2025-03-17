package model

type InternStatus int

const (
	NOT_STARTED InternStatus = iota
	ACTIVE
	COMPLETED
)