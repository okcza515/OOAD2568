// Wrote by MEP-1001
package authentication

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AuthenticationCLI struct {
	db         *gorm.DB
	middleware *Middleware
}

func NewAuthenticationCLI() *AuthenticationCLI {
	return &AuthenticationCLI{}
}

func (c *AuthenticationCLI) SetDB(db *gorm.DB) {
	c.db = db
	provider := NewDBAuthProvider(db, 24*time.Hour)
	c.middleware = NewMiddleware(provider)
}

func (c *AuthenticationCLI) SetAllowedRoles(roles []string) {
	if c.middleware != nil {
		c.middleware.SetAllowedRoles(roles)
	}
}

func (c *AuthenticationCLI) ExecuteItem(parameters []string) {
	ctx := context.Background()
	if err := RequireAdmin(ctx); err == nil {
		return
	}

	authMenu := NewAuthMenuState(c.db)
	if c.middleware != nil {
		authMenu.middleware = c.middleware
	}

	for {
		authMenu.Render()

		var input string
		fmt.Scanln(&input)

		err := authMenu.HandleUserInput(input)
		if err != nil {
			if err.Error() == "login_success" {
				ctx = authMenu.GetContext()
				return
			}
			if err.Error() == "exit" {
				return
			}
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}
}
