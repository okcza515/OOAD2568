package main

import (
	"ModEd/project/utils"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openDatabase(database string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}

func main() {
	builder := utils.NewMenuBuilder(&utils.MenuItem{
		Title: "Main",
		Children: []*utils.MenuItem{
			{
				Title: "SubMenu",
				Children: []*utils.MenuItem{
					{
						Title: "SayHi",
						Action: func(io *utils.MenuIO) {
							io.Println("Hi")
						},
					},
				},
			},
		},
	}, nil, nil)

	builder.AddMenuPath([]string{"File", "New"}, func(io *utils.MenuIO) {
		text, err := io.ReadInput()
		if err != nil {
			return
		}

		fmt.Println(text)
	})

	builder.AddMenuPath([]string{"File", "Open"}, func(io *utils.MenuIO) {
		io.Println("File opened")
	})

	builder.AddMenuPath([]string{"Edit", "Undo"}, func(io *utils.MenuIO) {
		io.Println("Undo action")
	})

	builder.Show()
}
