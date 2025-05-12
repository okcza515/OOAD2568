package menu

import (
	"fmt"

	"gorm.io/gorm"
)

type CommonCLIMenu struct {
	db   *gorm.DB
	menu *MenuHandler
	path string
}

func NewCommonCLIMenu(db *gorm.DB, path string) *CommonCLIMenu {
	menu := NewMenuHandler()
	cli := &CommonCLIMenu{
		db:   db,
		menu: menu,
		path: path,
	}

	menu.AppendItem("1", "Read file", &ReadFileHandler{path: path})
	menu.AppendItem("2", "Register", &RegisterHandler{db: db, path: path})
	menu.AppendItem("3", "Retrieve", &RetrieveHandler{db: db})
	menu.AppendItem("4", "Delete", &DeleteHandler{db: db})
	menu.AppendItem("5", "Clear Database", &ClearDBHandler{db: db})
	menu.AppendItem("6", "Test", &TestHandler{db: db})
	menu.AppendItem("exit", "Exit", &ExitHandler{})

	return cli
}

func (c *CommonCLIMenu) Run() {
	for {
		c.menu.DisplayMenu()
		choice := c.menu.GetMenuChoice()
		if choice == "exit" {
			fmt.Println("Goodbye!")
			return
		}
		c.menu.Execute(choice, nil)
	}
}
