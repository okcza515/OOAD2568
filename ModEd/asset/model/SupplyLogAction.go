package model

type SupplyLogAction int

const (
	ADD_NEW SupplyLogAction = iota
	UPDATEINFO
	RESTOCK
	DEPLETE
	DELETEFROMSYSTEM
)
