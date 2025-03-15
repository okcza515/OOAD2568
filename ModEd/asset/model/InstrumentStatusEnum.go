package model

type InstrumentStatusEnum int

const (
	INS_AVAILABLE InstrumentStatusEnum = iota
	INS_BORROWED
	INS_BROKEN
	INS_LOST
	INS_SALVAGING
	INS_SALVAGED
	INS_DONATED
)

var InstrumentStatusEnumLabel = map[InstrumentStatusEnum]string{
	INS_AVAILABLE: "Available",
	INS_BORROWED:  "Borrowed",
	INS_BROKEN:    "Broken",
	INS_LOST:      "Lost",
	INS_SALVAGING: "Salvaging",
	INS_SALVAGED:  "Salvaged",
	INS_DONATED:   "Donated",
}

func (status InstrumentStatusEnum) String() string {
	return InstrumentStatusEnumLabel[status]
}
