package asset

type InstrumentLogActionEnum string

const (
	INS_LOG_ADDNEW    InstrumentLogActionEnum = "AddNew"
	INS_LOG_UPDATE    InstrumentLogActionEnum = "Update"
	INS_LOG_BORROW    InstrumentLogActionEnum = "Borrow"
	INS_LOG_RETURN    InstrumentLogActionEnum = "Return"
	INS_LOG_MOVE      InstrumentLogActionEnum = "Move"
	INS_LOG_BROKEN    InstrumentLogActionEnum = "Broken"
	INS_LOG_REPAIR    InstrumentLogActionEnum = "Repair"
	INS_LOG_LOST      InstrumentLogActionEnum = "Lost"
	INS_LOG_FOUND     InstrumentLogActionEnum = "Found"
	INS_LOG_SALVAGING InstrumentLogActionEnum = "Salvaging"
	INS_LOG_SALVAGE   InstrumentLogActionEnum = "Salvage"
	INS_LOG_DONATE    InstrumentLogActionEnum = "Donate"
)
