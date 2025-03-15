package model

type InstrumentLogActionEnum int

const (
	INS_LOG_ADDNEW InstrumentLogActionEnum = iota
	INS_LOG_UPDATE
	INS_LOG_BORROW
	INS_LOG_RETURN
	INS_LOG_MOVE
	INS_LOG_BROKEN
	INS_LOG_REPAIR
	INS_LOG_LOST
	INS_LOG_FOUND
	INS_LOG_SALVAGING
	INS_LOG_SALVAGE
	INS_LOG_DONATE
)

var InstrumentLogActionLabel = map[InstrumentLogActionEnum]string{
	INS_LOG_ADDNEW:    "AddNew",
	INS_LOG_UPDATE:    "Update",
	INS_LOG_BORROW:    "Borrow",
	INS_LOG_RETURN:    "Return",
	INS_LOG_MOVE:      "Move",
	INS_LOG_BROKEN:    "Broken",
	INS_LOG_REPAIR:    "Repair",
	INS_LOG_LOST:      "Lost",
	INS_LOG_FOUND:     "Found",
	INS_LOG_SALVAGING: "Salvaging",
	INS_LOG_SALVAGE:   "Salvage",
	INS_LOG_DONATE:    "Donate",
}

func (program InstrumentLogActionEnum) ToString() string {
	return InstrumentLogActionLabel[program]
}
