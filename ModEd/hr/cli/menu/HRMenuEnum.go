package menu

type HRMenuEnum string

const (
	MENU_HR      HRMenuEnum = "hr"
	MENU_ADD     HRMenuEnum = "add"
	MENU_LIST    HRMenuEnum = "list"
	MENU_UPDATE  HRMenuEnum = "update"
	MENU_DELETE  HRMenuEnum = "delete"
	MENU_IMPORT  HRMenuEnum = "import"
	MENU_EXPORT  HRMenuEnum = "export"
	MENU_PULL    HRMenuEnum = "pull"
	MENU_REQUEST HRMenuEnum = "request"
	MENU_REVIEW  HRMenuEnum = "review"
	MENU_MIGRATE HRMenuEnum = "migrate"
)
