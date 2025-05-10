package cli

import (
	"context"
	"fmt"

	"ModEd/core/authentication"

	"gorm.io/gorm"
)

type AuthenticationCLI struct {
	db *gorm.DB
}

func NewAuthenticationCLI() *AuthenticationCLI {
	return &AuthenticationCLI{}
}

func (c *AuthenticationCLI) SetDB(db *gorm.DB) {
	c.db = db
}

func (c *AuthenticationCLI) ExecuteItem(parameters []string) {
	ctx := context.Background()
	if err := authentication.RequireAdmin(ctx); err == nil {
		return
	}

	authMenu := authentication.NewAuthMenuState(c.db)

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
