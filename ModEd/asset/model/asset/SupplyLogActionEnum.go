package asset

type SupplyLogAction int

const (
	SUP_ADDNEW SupplyLogAction = iota
	SUP_UPDATEINFO
	SUP_RESTOCK
	SUP_DEPLETE
	SUP_DELETEFROMSYSTEM
)

var SupplyLogActionLabel = map[SupplyLogAction]string{
	SUP_ADDNEW:           "AddNew",
	SUP_UPDATEINFO:       "UpdateInfo",
	SUP_RESTOCK:          "Restock",
	SUP_DEPLETE:          "Deplete",
	SUP_DELETEFROMSYSTEM: "DeleteFromSystem",
}

func (status SupplyLogAction) String() string {
	return SupplyLogActionLabel[status]
}
