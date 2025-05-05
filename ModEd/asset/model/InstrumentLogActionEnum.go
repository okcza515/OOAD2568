package model

import "errors"

// MEP-1012 Asset

type InstrumentLogActionEnum string

var enumMap = map[string]InstrumentLogActionEnum{
	"AddNew":    INS_LOG_ADDNEW,
	"Update":    INS_LOG_UPDATE,
	"Borrow":    INS_LOG_BORROW,
	"Return":    INS_LOG_RETURN,
	"Move":      INS_LOG_MOVE,
	"Broken":    INS_LOG_BROKEN,
	"Repair":    INS_LOG_REPAIR,
	"Lost":      INS_LOG_LOST,
	"Found":     INS_LOG_FOUND,
	"Salvaging": INS_LOG_SALVAGING,
	"Salvage":   INS_LOG_SALVAGE,
	"Donate":    INS_LOG_DONATE,
}

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

func ToInstrumentActionEnum(str string) (InstrumentLogActionEnum, error) {
	action, ok := enumMap[str]
	if !ok {
		return "", errors.New("err: cannot map string into enum")
	}

	return action, nil
}
