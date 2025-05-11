// Wrote by MEP-1001
package authentication

import (
	"context"
	"fmt"
	"os"
	"time"

	"ModEd/core/handler"

	"gorm.io/gorm"
)

type AuthMenuState struct {
	middleware     *Middleware
	ctx            context.Context
	handlerContext *handler.HandlerContext
}

type LoginHandler struct {
	state *AuthMenuState
}

func (h LoginHandler) Execute() error {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	userCtx, err := h.state.middleware.Authenticate(h.state.ctx, username, password)
	if err != nil {
		if err == ErrRoleNotAllowed {
			fmt.Println("Access denied: Your role is not allowed to use this program")
			return nil
		}
		return fmt.Errorf("login failed: %v", err)
	}

	h.state.ctx = WithContext(context.Background(), userCtx)
	fmt.Printf("Login successful! Welcome %s (Role: %s)\n", userCtx.Username, userCtx.Role)

	return fmt.Errorf("login_success")
}

type RegisterHandler struct {
	state *AuthMenuState
}

func (h RegisterHandler) Execute() error {
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

func (h ChangePasswordHandler) Execute() error {
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

func (h DeleteAccountHandler) Execute() error {
	var username, password string
	fmt.Print("Username to delete: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	// Verify credentials directly with provider to bypass role check
	provider := NewDBAuthProvider(h.state.middleware.provider.(*DBAuthProvider).db, 24*time.Hour)
	_, err := provider.Authenticate(h.state.ctx, username, password)
	if err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	// Proceed with deletion
	err = h.state.middleware.DeleteUser(h.state.ctx, username)
	if err != nil {
		return fmt.Errorf("account deletion failed: %v", err)
	}

	fmt.Println("Account deleted successfully!")
	return nil
}

type ExitHandler struct{}

func (h ExitHandler) Execute() error {
	os.Exit(0)
	return nil
}

func NewAuthMenuState(db *gorm.DB) *AuthMenuState {
	provider := NewDBAuthProvider(db, 24*time.Hour)
	middleware := NewMiddleware(provider)

	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Authentication Menu")

	state := &AuthMenuState{
		middleware:     middleware,
		ctx:            context.Background(),
		handlerContext: ctx,
	}

	ctx.AddHandler("1", "Login", LoginHandler{state})
	ctx.AddHandler("2", "Register", RegisterHandler{state})
	ctx.AddHandler("3", "Change Password", ChangePasswordHandler{state})
	ctx.AddHandler("4", "Delete Account", DeleteAccountHandler{state})
	ctx.AddHandler("5", "exit", ExitHandler{})

	return state
}

func (a *AuthMenuState) Render() {
	a.handlerContext.ShowMenu()
}

func (a *AuthMenuState) HandleUserInput(input string) error {
	return a.handlerContext.HandleInput(input)
}

func (a *AuthMenuState) GetContext() context.Context {
	return a.ctx
}
