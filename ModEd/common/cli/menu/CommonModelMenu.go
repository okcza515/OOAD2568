package menu

import (
	"gorm.io/gorm"
)

type CommonModelMenu struct {
	db   *gorm.DB
	menu *MenuHandler
}

func NewCommonModelMenu(db *gorm.DB) *CommonModelMenu {
	menu := NewMenuHandler()
	modelMenu := &CommonModelMenu{
		db:   db,
		menu: menu,
	}

	menu.AppendItem("1", "Student", &RegisterModelHandler{db: db, modelType: 1})
	menu.AppendItem("2", "Instructor", &RegisterModelHandler{db: db, modelType: 2})
	menu.AppendItem("3", "Department", &RegisterModelHandler{db: db, modelType: 3})
	menu.AppendItem("4", "Faculty", &RegisterModelHandler{db: db, modelType: 4})
	menu.AppendItem("back", "Back", &BackHandler{})

	return modelMenu
}

func (m *CommonModelMenu) Run() {
	for {
		m.menu.DisplayMenu()
		choice := m.menu.GetMenuChoice()
		if choice == "back" {
			return
		}
		m.menu.Execute(choice, nil)
	}
}
