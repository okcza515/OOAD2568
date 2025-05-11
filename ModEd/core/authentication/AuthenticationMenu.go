// Wrote by MEP-1001
package authentication

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AuthMenuState struct {
	middleware *Middleware
	ctx        context.Context
	handlers   map[string]func() error
}

func NewAuthMenuState(db *gorm.DB) *AuthMenuState {
	provider := NewDBAuthProvider(db, 24*time.Hour)
	middleware := NewMiddleware(provider)
	state := &AuthMenuState{
		middleware: middleware,
		ctx:        context.Background(),
		handlers:   make(map[string]func() error),
	}

	// Initialize menu items
	state.handlers["1"] = func() error {
		var username, password string
		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		fmt.Scanln(&password)

		userCtx, err := state.middleware.Authenticate(state.ctx, username, password)
		if err != nil {
			return fmt.Errorf("login failed: %v", err)
		}

		state.ctx = WithContext(context.Background(), userCtx)
		fmt.Printf("Login successful! Welcome %s (Role: %s)\n", userCtx.Username, userCtx.Role)
		return fmt.Errorf("login_success")
	}

	state.handlers["2"] = func() error {
		var username, password, role string
		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		fmt.Scanln(&password)
		fmt.Print("Role (user/admin): ")
		fmt.Scanln(&role)

		err := state.middleware.CreateUser(state.ctx, username, password, role)
		if err != nil {
			return fmt.Errorf("registration failed: %v", err)
		}

		fmt.Println("Registration successful!")
		return nil
	}

	state.handlers["3"] = func() error {
		var username, oldPass, newPass string
		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Current Password: ")
		fmt.Scanln(&oldPass)
		fmt.Print("New Password: ")
		fmt.Scanln(&newPass)

		err := state.middleware.UpdatePassword(state.ctx, username, oldPass, newPass)
		if err != nil {
			return fmt.Errorf("password change failed: %v", err)
		}

		fmt.Println("Password changed successfully!")
		return nil
	}

	state.handlers["4"] = func() error {
		var username string
		fmt.Print("Username to delete: ")
		fmt.Scanln(&username)

		err := state.middleware.DeleteUser(state.ctx, username)
		if err != nil {
			return fmt.Errorf("account deletion failed: %v", err)
		}

		fmt.Println("Account deleted successfully!")
		return nil
	}

	state.handlers["exit"] = func() error {
		return fmt.Errorf("exit")
	}

	return state
}

func (a *AuthMenuState) Render() {
	fmt.Println("\n=== Authentication Menu ===")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Change Password")
	fmt.Println("4. Delete Account")
	fmt.Println("exit. Exit")
	fmt.Print("Select an option: ")
}

func (a *AuthMenuState) HandleUserInput(input string) error {
	if handler, exists := a.handlers[input]; exists {
		return handler()
	}
	return fmt.Errorf("invalid option")
}

func (a *AuthMenuState) GetContext() context.Context {
	return a.ctx
}
