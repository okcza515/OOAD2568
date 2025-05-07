# Core Handler

## The main context managing the menu and handlers:

- Fields:

  - title: The title shown above the menu.

  - menu: A map of user input keys to menu items.

- Key methods:

  - `NewHandlerContext()`: Initializes a new handler context.

  - `AddHandler(userInput, headerLabel, strategy)`: Adds a new menu item.

  - `HandleInput(userInput)`: Executes the strategy for the given input.

  - `ShowMenu()`: Displays the menu sorted by input keys.

  - `SetMenuTitle(title)`: Sets the menu title.

  - `AddBackHandler(strategy)`: Adds a predefined "back" handler.

## Example:

```go
package main

import (
	"fmt"
	"core/handler"
)

func helloWorld() error {
	fmt.Println("Hello, World!")
	return nil
}

type generalHandler struct {}

func (handler generalHandler) Execute() error {
	fmt.Println("Bye bye")
	return nil
}

func main() {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Main Menu")

	ctx.AddHandler("1", "Say Hello", handler.FuncStrategy{Action: helloWorld}})
	ctx.AddHandler("2", "Say Goodbye", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Goodbye!")
			return nil
		},
	})
    ctx.AddHandler("3", "Say Hello", generalHandler)
	ctx.AddBackHandler(handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Exiting...")
			return nil
		},
	})

	for {
		ctx.ShowMenu()
		var userInput string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&userInput)
		if err := ctx.HandleInput(userInput); err != nil {
			fmt.Println("Error:", err)
		}
		if userInput == "back" {
			break
		}
	}
}
```

## Example Cases

Here are three example cases from the usage in `main.go`:

---

### Case 1: Using `FuncStrategy` with a Named Function

```go
ctx.AddHandler("1", "Say Hello", handler.FuncStrategy{Action: helloWorld})
```

You can use any `func() error` function (like helloWorld) and wrap it with
`FuncStrategy{Action: func}` to convert it into a MenuStrategy that fits the HandlerContext system.

### Case 2: Using `FuncStrategy` with an Inline Anonymous Function

```go
ctx.AddHandler("2", "Say Goodbye", handler.FuncStrategy{
    Action: func() error {
        fmt.Println("Goodbye!")
        return nil
    },
})
```
You can define quick, one-off behaviors inline using anonymous func() error blocks
and wrap them directly with FuncStrategy. This saves you from having to create separate named functions

### Case 3: Using a Custom Struct That Implements the MenuStrategy Interface

```go
type generalHandler struct {}

func (handler generalHandler) Execute() error {
    fmt.Println("Bye bye")
    return nil
}

ctx.AddHandler("3", "Say Hello", generalHandler)
```
Instead of wrapping a func() error with FuncStrategy, you can create your own
struct type that directly implements the MenuStrategy interface by defining an Execute method.
This approach is useful for handling more complex behaviors, adding state, or encapsulating related logic inside a type.