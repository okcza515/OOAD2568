package menu

type ProcurementMenuEnum string

const (
	MENU_PROCUREMENT_MAIN   ProcurementMenuEnum = "procurement-main"
	MENU_INSTRUMENT_REQUEST ProcurementMenuEnum = "instrument-request"
	MENU_PROCUREMENT        ProcurementMenuEnum = "procurement"
	MENU_APPROVAL           ProcurementMenuEnum = "approval"
	MENU_QUOTATION          ProcurementMenuEnum = "quotation"
	MENU_ACCEPTANCE         ProcurementMenuEnum = "acceptance"
	MENU_TOR                ProcurementMenuEnum = "tor"
	MENU_ACCEPTEDINSTRUMENT  ProcurementMenuEnum = "accepted-instrument"
)
