package menu

type AssetMenuEnum string

const (
	MENU_ASSET          AssetMenuEnum = "asset"
	MENU_CATEGORY       AssetMenuEnum = "category"
	MENU_INSTRUMENT     AssetMenuEnum = "instrument"
	MENU_INSTRUMENT_LOG AssetMenuEnum = "instrument-log"
	MENU_SUPPLY         AssetMenuEnum = "supply"
	MENU_SUPPLY_LOG     AssetMenuEnum = "supply-log"
	MENU_BORROW         AssetMenuEnum = "borrow"
	MENU_REPORT         AssetMenuEnum = "report"
)
