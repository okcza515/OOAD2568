package menu

type ProcurementMenuEnum string

const (
	MENU_PROCUREMENT_MAIN   ProcurementMenuEnum = "procurement-main"
	MENU_INSTRUMENT_REQUEST ProcurementMenuEnum = "instrument-request"
	MENU_APPROVAL           ProcurementMenuEnum = "approval"
	MENU_QUOTATION          ProcurementMenuEnum = "quotation"
	MENU_TOR                ProcurementMenuEnum = "tor"
)
