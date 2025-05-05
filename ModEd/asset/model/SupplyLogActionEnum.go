package model

import "errors"

// MEP-1012 Asset

type SupplyLogActionEnum int

var enumSupplyMap = map[string]SupplyLogActionEnum{
	"AddNew":           SUP_ADDNEW,
	"UpdateInfo":       SUP_UPDATEINFO,
	"Restock":          SUP_RESTOCK,
	"Deplete":          SUP_DEPLETE,
	"DeleteFromSystem": SUP_DELETEFROMSYSTEM,
}

const (
	SUP_ADDNEW SupplyLogActionEnum = iota
	SUP_UPDATEINFO
	SUP_RESTOCK
	SUP_DEPLETE
	SUP_DELETEFROMSYSTEM
)

var SupplyLogActionLabel = map[SupplyLogActionEnum]string{
	SUP_ADDNEW:           "AddNew",
	SUP_UPDATEINFO:       "UpdateInfo",
	SUP_RESTOCK:          "Restock",
	SUP_DEPLETE:          "Deplete",
	SUP_DELETEFROMSYSTEM: "DeleteFromSystem",
}

func (status SupplyLogActionEnum) String() string {
	return SupplyLogActionLabel[status]
}

func ToSupplyActionEnum(str string) (SupplyLogActionEnum, error) {
	action, ok := enumSupplyMap[str]
	if !ok {
		return -1, errors.New("err: cannot map string into enum")
	}

	return action, nil
}
