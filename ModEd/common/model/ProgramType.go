package model

type ProgramType int

const (
	REGULAR ProgramType = iota
	INTERNATIONAL
)

var ProgramTypeLabel = map[ProgramType]string{
	REGULAR:       "Regular",
	INTERNATIONAL: "International",
}

func (program ProgramType) String() string {
	return ProgramTypeLabel[program]
}
