package credential

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AuthMenuState struct {
	middleware *Middleware
	ctx        context.Context
}

func (a *AuthMenuState) Render() {
	fmt.Println("\n=== Authentication Menu ===")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Change Password")
	fmt.Println("4. Delete Account")
	fmt.Println("5. Exit")
	fmt.Print("Select an option: ")
}

func (a *AuthMenuState) HandleUserInput(input string) error {
	switch input {
	case "1":
		return a.handleLogin()
	case "2":
		return a.handleRegister()
	case "3":
		return a.handleChangePassword()
	case "4":
		return a.handleDeleteAccount()
	case "5":
		return fmt.Errorf("exit")
	default:
		return fmt.Errorf("invalid option")
	}
}

func (a *AuthMenuState) handleLogin() error {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	userCtx, err := a.middleware.Authenticate(a.ctx, username, password)
	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}

	a.ctx = WithContext(context.Background(), userCtx)
	fmt.Printf("Login successful! Welcome %s (Role: %s)\n", userCtx.Username, userCtx.Role)

	return fmt.Errorf("login_success")
}

func (a *AuthMenuState) handleRegister() error {
	var username, password, role string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	fmt.Print("Role (user/admin): ")
	fmt.Scanln(&role)

	err := a.middleware.CreateUser(a.ctx, username, password, role)
	if err != nil {
		return fmt.Errorf("registration failed: %v", err)
	}

	fmt.Println("Registration successful!")
	return nil
}

func (a *AuthMenuState) handleChangePassword() error {
	var username, oldPass, newPass string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Current Password: ")
	fmt.Scanln(&oldPass)
	fmt.Print("New Password: ")
	fmt.Scanln(&newPass)

	err := a.middleware.UpdatePassword(a.ctx, username, oldPass, newPass)
	if err != nil {
		return fmt.Errorf("password change failed: %v", err)
	}

	fmt.Println("Password changed successfully!")
	return nil
}

func (a *AuthMenuState) handleDeleteAccount() error {
	var username string
	fmt.Print("Username to delete: ")
	fmt.Scanln(&username)

	err := a.middleware.DeleteUser(a.ctx, username)
	if err != nil {
		return fmt.Errorf("account deletion failed: %v", err)
	}

	fmt.Println("Account deleted successfully!")
	return nil
}

func (a *AuthMenuState) GetContext() context.Context {
	return a.ctx
}

func NewAuthMenuState(db *gorm.DB) *AuthMenuState {
	provider := NewDBAuthProvider(db, 24*time.Hour)
	middleware := NewMiddleware(provider)
	return &AuthMenuState{
		middleware: middleware,
		ctx:        context.Background(),
	}
}
