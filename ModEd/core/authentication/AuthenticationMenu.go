package authentication

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type AuthMenuItemHandler interface {
	ExecuteItem() error
}

type AuthMenuHandler struct {
	itemHandlerMap map[string]AuthMenuItemHandler
	itemLabelMap   map[string]string
	items          []string
}

func NewAuthMenuHandler() *AuthMenuHandler {
	return &AuthMenuHandler{
		itemHandlerMap: make(map[string]AuthMenuItemHandler),
		itemLabelMap:   make(map[string]string),
		items:          []string{},
	}
}

func (handler *AuthMenuHandler) AppendItem(key string, label string, itemHandler AuthMenuItemHandler) {
	handler.itemHandlerMap[key] = itemHandler
	handler.itemLabelMap[key] = label
	handler.items = append(handler.items, key)
}

func (handler *AuthMenuHandler) Execute(selectedMenu string) error {
	// Convert numeric input to menu key
	if index, err := strconv.Atoi(selectedMenu); err == nil {
		if index > 0 && index <= len(handler.items) {
			selectedMenu = handler.items[index-1]
		}
	}

	if itemHandler, exists := handler.itemHandlerMap[selectedMenu]; exists {
		return itemHandler.ExecuteItem()
	}
	return fmt.Errorf("invalid option")
}

func (handler *AuthMenuHandler) DisplayMenu() {
	fmt.Println("\n=== Authentication Menu ===")
	for i, key := range handler.items {
		fmt.Printf("%d. %s\n", i+1, handler.itemLabelMap[key])
	}
	fmt.Print("Select an option: ")
}

func (handler *AuthMenuHandler) GetMenuChoice() string {
	var choiceIndex int
	fmt.Scan(&choiceIndex)

	if choiceIndex > 0 && choiceIndex <= len(handler.items) {
		return handler.items[choiceIndex-1]
	}

	return ""
}

type AuthMenuState struct {
	middleware *Middleware
	ctx        context.Context
	handler    *AuthMenuHandler
}

type LoginHandler struct {
	state *AuthMenuState
}

func (h *LoginHandler) ExecuteItem() error {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	userCtx, err := h.state.middleware.Authenticate(h.state.ctx, username, password)
	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}

	h.state.ctx = WithContext(context.Background(), userCtx)
	fmt.Printf("Login successful! Welcome %s (Role: %s)\n", userCtx.Username, userCtx.Role)

	return fmt.Errorf("login_success")
}

type RegisterHandler struct {
	state *AuthMenuState
}

func (h *RegisterHandler) ExecuteItem() error {
	var username, password, role string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	fmt.Print("Role (user/admin): ")
	fmt.Scanln(&role)

	err := h.state.middleware.CreateUser(h.state.ctx, username, password, role)
	if err != nil {
		return fmt.Errorf("registration failed: %v", err)
	}

	fmt.Println("Registration successful!")
	return nil
}

type ChangePasswordHandler struct {
	state *AuthMenuState
}

func (h *ChangePasswordHandler) ExecuteItem() error {
	var username, oldPass, newPass string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Current Password: ")
	fmt.Scanln(&oldPass)
	fmt.Print("New Password: ")
	fmt.Scanln(&newPass)

	err := h.state.middleware.UpdatePassword(h.state.ctx, username, oldPass, newPass)
	if err != nil {
		return fmt.Errorf("password change failed: %v", err)
	}

	fmt.Println("Password changed successfully!")
	return nil
}

type DeleteAccountHandler struct {
	state *AuthMenuState
}

func (h *DeleteAccountHandler) ExecuteItem() error {
	var username string
	fmt.Print("Username to delete: ")
	fmt.Scanln(&username)

	err := h.state.middleware.DeleteUser(h.state.ctx, username)
	if err != nil {
		return fmt.Errorf("account deletion failed: %v", err)
	}

	fmt.Println("Account deleted successfully!")
	return nil
}

type ExitHandler struct{}

func (h *ExitHandler) ExecuteItem() error {
	return fmt.Errorf("exit")
}

func (a *AuthMenuState) Render() {
	a.handler.DisplayMenu()
}

func (a *AuthMenuState) HandleUserInput(input string) error {
	return a.handler.Execute(input)
}

func (a *AuthMenuState) GetContext() context.Context {
	return a.ctx
}

func NewAuthMenuState(db *gorm.DB) *AuthMenuState {
	provider := NewDBAuthProvider(db, 24*time.Hour)
	middleware := NewMiddleware(provider)
	state := &AuthMenuState{
		middleware: middleware,
		ctx:        context.Background(),
		handler:    NewAuthMenuHandler(),
	}

	// Initialize menu items
	state.handler.AppendItem("login", "Login", &LoginHandler{state: state})
	state.handler.AppendItem("register", "Register", &RegisterHandler{state: state})
	state.handler.AppendItem("changepass", "Change Password", &ChangePasswordHandler{state: state})
	state.handler.AppendItem("delete", "Delete Account", &DeleteAccountHandler{state: state})
	state.handler.AppendItem("exit", "Exit", &ExitHandler{})

	return state
}
