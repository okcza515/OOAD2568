package handler

import (
	"fmt"
)

// MenuState interface for managing menu states
type MenuState interface {
	Enter() error
	Exit() error
	HandleInput(input string) (MenuState, error)
}

// MenuStateManager manages menu state transitions
type MenuStateManager struct {
	currentState MenuState
}

// NewMenuStateManager creates a new menu state manager
func NewMenuStateManager(initialState MenuState) *MenuStateManager {
	return &MenuStateManager{
		currentState: initialState,
	}
}

// Run executes the menu state manager loop
func (m *MenuStateManager) Run() error {
	if err := m.currentState.Enter(); err != nil {
		return err
	}

	for {
		var input string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&input)

		nextState, err := m.currentState.HandleInput(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if nextState == nil {
			return m.currentState.Exit()
		}

		if err := m.currentState.Exit(); err != nil {
			return err
		}

		m.currentState = nextState

		if err := m.currentState.Enter(); err != nil {
			return err
		}
	}
}

// Define FuncStrategy to encapsulate action functions
type FuncStrategy struct {
	Action func() error
}

// Define NewHandlerContext to create a new handler context
func NewHandlerContext() *HandlerContext {
	return &HandlerContext{
		handlers: make(map[string]FuncStrategy),
	}
}

// Define HandlerContext to manage menu states
type HandlerContext struct {
	menuTitle string
	handlers  map[string]FuncStrategy
}

func (ctx *HandlerContext) SetMenuTitle(title string) {
	ctx.menuTitle = title
}

func (ctx *HandlerContext) AddHandler(key, description string, strategy FuncStrategy) {
	ctx.handlers[key] = strategy
}

func (ctx *HandlerContext) AddBackHandler(strategy FuncStrategy) {
	ctx.AddHandler("back", "Back", strategy)
}

func (ctx *HandlerContext) ShowMenu() {
	fmt.Printf("\n%s Menu:\n", ctx.menuTitle)
	for key := range ctx.handlers {
		fmt.Printf("%s. %s\n", key, key)
	}
}

func (ctx *HandlerContext) HandleInput(input string) error {
	if strategy, exists := ctx.handlers[input]; exists {
		return strategy.Action()
	}
	fmt.Println("Invalid option")
	return nil
}
