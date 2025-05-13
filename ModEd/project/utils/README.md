# CLI Menu Builder (Tree-Based)

A simple Go package to build and display interactive CLI menus using a tree structure. Each menu item can contain submenus or trigger an action with full access to input/output streams.

---

## ✨ Features

- Nested tree-style menus
- Interactive CLI navigation
- Dynamic menu building with `AddMenuPath`
- Actions can receive user input/output through `MenuIO`
- Supports `Back` and `Exit` navigation automatically

---

## 📦 Installation

1. Clone or copy this package into your project (e.g., `utils/menu.go`)
2. Import it:

```go
import "yourmodule/utils" // Update to match your module name
```

## 🚀 Usage

```go
package main

import (
	"fmt"
	"yourmodule/utils" // Replace with your module name
)

func main() {
	builder := utils.NewMenuBuilder(&utils.MenuItem{
		Title: "Main",
		Children: []*utils.MenuItem{
			{
				Title: "SubMenu",
				Children: []*utils.MenuItem{
					{
						Title: "SayHi",
						Action: func(io *core.MenuIO) {
							io.Println("Hi")
						},
					},
				},
			},
		},
	}, nil, nil)

	builder.AddMenuPath([]string{"File", "New"}, func(io *core.MenuIO) {
		text, err := io.ReadInput()
		if err != nil {
			return
		}

		io.Println(text)
	})

	builder.AddMenuPath([]string{"File", "Open"}, func(io *core.MenuIO) {
		io.Println("File opened")
	})

	builder.AddMenuPath([]string{"Edit", "Undo"}, func(io *core.MenuIO) {
		io.Println("Undo action")
	})

	builder.Show()
}
```

## Example CLI Output
```
> Main
[1] SubMenu
[2] File
[3] Edit
[0] Exit
Select an option: 1

/Main/SubMenu> SubMenu
[1] SayHi
[0] Back
Select an option: 1
Hi

/Main/SubMenu> SubMenu
[0] Back
Select an option: 0

> Main
Select an option: 0

```