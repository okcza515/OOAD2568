package model

type InstrumentStatusEnum int

const (
	STATUS_AVAILABLE InstrumentStatusEnum = iota
	STATUS_BORROWED
	STATUS_BROKEN
	STATUS_LOST
	STATUS_SALVAGING
	STATUS_SALVAGED
	STATUS_DONATED
)

var InstrumentStatusEnumLabel = map[InstrumentStatusEnum]string{
	STATUS_AVAILABLE: "Available",
	STATUS_BORROWED:  "Borrowed",
	STATUS_BROKEN:    "Broken",
	STATUS_LOST:      "Lost",
	STATUS_SALVAGING: "Salvaging",
	STATUS_SALVAGED:  "Salvaged",
	STATUS_DONATED:   "Donated",
}

func (status InstrumentStatusEnum) String() string {
	return InstrumentStatusEnumLabel[status]
}
