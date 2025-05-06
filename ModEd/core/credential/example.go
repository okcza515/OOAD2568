package credential

import (
	"context"
	"fmt"
	"time"
)

// ExampleUsage demonstrates how to use the credential middleware
func ExampleUsage() {
	// Create a new credential provider with 24-hour session expiry
	provider := NewDefaultProvider(24 * time.Hour)

	// Create the middleware
	middleware := NewMiddleware(provider)

	// Create a new user
	ctx := context.Background()
	err := middleware.CreateUser(ctx, "admin", "securepassword", "admin")
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}

	// Authenticate the user
	userCtx, err := middleware.Authenticate(ctx, "admin", "securepassword")
	if err != nil {
		fmt.Printf("Authentication failed: %v\n", err)
		return
	}

	// Add the user context to the request context
	ctx = WithContext(ctx, userCtx)

	// Later, retrieve the user context
	if user, ok := FromContext(ctx); ok {
		fmt.Printf("Authenticated as: %s (Role: %s)\n", user.Username, user.Role)
	}

	// Update password
	err = middleware.UpdatePassword(ctx, "admin", "securepassword", "newpassword")
	if err != nil {
		fmt.Printf("Error updating password: %v\n", err)
		return
	}

	// Delete user
	err = middleware.DeleteUser(ctx, "admin")
	if err != nil {
		fmt.Printf("Error deleting user: %v\n", err)
		return
	}
}

// ExampleCustomProvider shows how to create a custom credential provider
type CustomProvider struct {
	// Add your custom fields here
}

func (p *CustomProvider) Authenticate(ctx context.Context, username, password string) (*UserContext, error) {
	// Implement your custom authentication logic
	return nil, nil
}

func (p *CustomProvider) CreateUser(ctx context.Context, username, password string, role string) error {
	// Implement your custom user creation logic
	return nil
}

func (p *CustomProvider) DeleteUser(ctx context.Context, username string) error {
	// Implement your custom user deletion logic
	return nil
}

func (p *CustomProvider) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	// Implement your custom password update logic
	return nil
}

// ExampleCustomProviderUsage shows how to use a custom provider
func ExampleCustomProviderUsage() {
	// Create your custom provider
	customProvider := &CustomProvider{}

	// Create middleware with your custom provider
	middleware := NewMiddleware(customProvider)

	// Use the middleware as before
	ctx := context.Background()
	userCtx, err := middleware.Authenticate(ctx, "user", "password")
	if err != nil {
		fmt.Printf("Authentication failed: %v\n", err)
		return
	}

	// Add user context to request context
	ctx = WithContext(ctx, userCtx)
}
