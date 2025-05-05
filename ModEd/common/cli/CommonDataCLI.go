package main

import (
	"flag"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MenuItemHandler interface {
	ExecuteItem(parameters []string)
}

type MenuHandler struct {
	defaultHandler MenuItemHandler
	backHandler    MenuItemHandler
	itemHandlerMap map[string]MenuItemHandler
	itemLabelMap   map[string]string
	items          []string
}

func NewMenuHandler() *MenuHandler {
	return &MenuHandler{
		itemHandlerMap: make(map[string]MenuItemHandler),
		itemLabelMap:   make(map[string]string),
		items:          []string{},
	}
}

func (handler *MenuHandler) AppendItem(key string, label string, itemHandler MenuItemHandler) {
	handler.itemHandlerMap[key] = itemHandler
	handler.itemLabelMap[key] = label
	handler.items = append(handler.items, key)
}

func (handler *MenuHandler) SetBackHandler(itemHandler MenuItemHandler) {
	handler.backHandler = itemHandler
}

func (handler *MenuHandler) SetDefaultHandler(itemHandler MenuItemHandler) {
	handler.defaultHandler = itemHandler
}

func (handler *MenuHandler) Execute(selectedMenu string, parameters []string) {
	if selectedMenu == "back" && handler.backHandler != nil {
		handler.backHandler.ExecuteItem(parameters)
		return
	}

	if itemHandler, exists := handler.itemHandlerMap[selectedMenu]; exists {
		itemHandler.ExecuteItem(parameters)
		return
	}

	if handler.defaultHandler != nil {
		handler.defaultHandler.ExecuteItem(parameters)
	} else {
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}

func (handler *MenuHandler) DisplayMenu() {
	fmt.Println("\nSelect an option:")
	for i, key := range handler.items {
		fmt.Printf("%d. %s\n", i, handler.itemLabelMap[key])
	}
	fmt.Print("choice: ")
}

func (handler *MenuHandler) GetMenuChoice() string {
	var choiceIndex int
	fmt.Scan(&choiceIndex)

	if choiceIndex >= 0 && choiceIndex < len(handler.items) {
		return handler.items[choiceIndex]
	}

	return ""
}

func main() {
	var (
		database string
		path     string
	)
	flag.StringVar(&database, "database", "data/ModEd.bin", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "", "Path to CSV or JSON.")
	flag.Parse()

	args := flag.Args()
	fmt.Printf("args: %v\n", args)

	db := ConnectDB()

	menu := NewMenuHandler()

	menu.AppendItem("readfile", "Read file", &ReadFileHandler{path: path})
	menu.AppendItem("register", "Register", &RegisterHandler{db: db, path: path})
	menu.AppendItem("retrieve", "Retrieve", &RetrieveHandler{db: db})
	menu.AppendItem("delete", "Delete", &DeleteHandler{db: db})
	menu.AppendItem("cleardb", "Clear DB", &ClearDBHandler{db: db})
	menu.AppendItem("exit", "Exit", &ExitHandler{})
	menu.AppendItem("test", "Test", &TestHandler{db: db})

	menu.SetDefaultHandler(&DefaultHandler{})

	for {
		menu.DisplayMenu()
		choice := menu.GetMenuChoice()
		menu.Execute(choice, args)
	}
}

func ConnectDB() *gorm.DB {
	connector, err := gorm.Open(sqlite.Open("data/ModEd.bin"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return connector
}

func confirmAction(prompt string) bool {
	var response string
	fmt.Print(prompt)
	fmt.Scan(&response)
	return response == "y" || response == "Y"
}
