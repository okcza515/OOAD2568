// Wrote by MEP-1001
package authentication

import (
	"context"
	"fmt"
)

type LoginHandler struct {
	State *AuthMenuState
}

func (h *LoginHandler) Execute() error {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	userCtx, err := h.State.middleware.Authenticate(h.State.ctx, username, password)
	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}

	h.State.ctx = WithContext(context.Background(), userCtx)
	fmt.Printf("Login successful! Welcome %s (Role: %s)\n", userCtx.Username, userCtx.Role)
	return fmt.Errorf("login_success")
}

type RegisterHandler struct {
	State *AuthMenuState
}

func (h *RegisterHandler) Execute() error {
	var username, password, role string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	fmt.Print("Role (user/admin): ")
	fmt.Scanln(&role)

	err := h.State.middleware.CreateUser(h.State.ctx, username, password, role)
	if err != nil {
		return fmt.Errorf("registration failed: %v", err)
	}

	fmt.Println("Registration successful!")
	return nil
}

type ChangePasswordHandler struct {
	State *AuthMenuState
}

func (h *ChangePasswordHandler) Execute() error {
	var username, oldPass, newPass string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Current Password: ")
	fmt.Scanln(&oldPass)
	fmt.Print("New Password: ")
	fmt.Scanln(&newPass)

	err := h.State.middleware.UpdatePassword(h.State.ctx, username, oldPass, newPass)
	if err != nil {
		return fmt.Errorf("password change failed: %v", err)
	}

	fmt.Println("Password changed successfully!")
	return nil
}

type DeleteAccountHandler struct {
	State *AuthMenuState
}

func (h *DeleteAccountHandler) Execute() error {
	var username string
	fmt.Print("Username to delete: ")
	fmt.Scanln(&username)

	err := h.State.middleware.DeleteUser(h.State.ctx, username)
	if err != nil {
		return fmt.Errorf("account deletion failed: %v", err)
	}

	fmt.Println("Account deleted successfully!")
	return nil
}

type ExitHandler struct{}

func (h *ExitHandler) Execute() error {
	return fmt.Errorf("exit")
}
