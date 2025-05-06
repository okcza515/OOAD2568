package cli

import (
	"context"
	"fmt"

	"ModEd/core/credential"

	"gorm.io/gorm"
)

type CredentialCLI struct {
	db *gorm.DB
}

func NewCredentialCLI() *CredentialCLI {
	return &CredentialCLI{}
}

func (c *CredentialCLI) SetDB(db *gorm.DB) {
	c.db = db
}

func (c *CredentialCLI) ExecuteItem(parameters []string) {
	ctx := context.Background()
	if err := credential.RequireAdmin(ctx); err == nil {
		return
	}

	authMenu := credential.NewAuthMenuState(c.db)

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
