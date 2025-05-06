package credential

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ExampleApplication shows how to use the authentication system in a real application
func ExampleApplication(db *gorm.DB) {
	// Create a database-backed authentication provider with 24-hour session expiry
	authProvider := NewDBAuthProvider(db, 24*time.Hour)

	// Create the middleware
	middleware := NewMiddleware(authProvider)

	// Example: User registration
	ctx := context.Background()
	err := middleware.CreateUser(ctx, "john.doe", "securepass123", "user")
	if err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
		return
	}

	// Example: User login
	userCtx, err := middleware.Authenticate(ctx, "john.doe", "securepass123")
	if err != nil {
		fmt.Printf("Authentication failed: %v\n", err)
		return
	}

	// Add user context to request context
	ctx = WithContext(ctx, userCtx)

	// Example: Access user information in a handler
	if user, ok := FromContext(ctx); ok {
		fmt.Printf("User logged in: %s (Role: %s)\n", user.Username, user.Role)
	}

	// Example: Password update
	err = middleware.UpdatePassword(ctx, "john.doe", "securepass123", "newpass456")
	if err != nil {
		fmt.Printf("Failed to update password: %v\n", err)
		return
	}

	// Example: User deletion
	err = middleware.DeleteUser(ctx, "john.doe")
	if err != nil {
		fmt.Printf("Failed to delete user: %v\n", err)
		return
	}
}

// ExampleMiddleware shows how to create a custom authentication provider
func ExampleMiddleware() {
	// Create a custom authentication provider
	customProvider := NewCustomAuthProvider(24 * time.Hour)

	// Create middleware with custom provider
	middleware := NewMiddleware(customProvider)

	// Use the middleware
	ctx := context.Background()
	userCtx, err := middleware.Authenticate(ctx, "user", "pass")
	if err != nil {
		fmt.Printf("Authentication failed: %v\n", err)
		return
	}

	// Add user context to request context
	ctx = WithContext(ctx, userCtx)

	// Access user information
	if user, ok := FromContext(ctx); ok {
		fmt.Printf("User authenticated: %s\n", user.Username)
	}
}

// ExampleProtectedHandler shows how to protect routes/handlers
func ExampleProtectedHandler(ctx context.Context) error {
	// Check if user is authenticated
	user, ok := FromContext(ctx)
	if !ok {
		return fmt.Errorf("unauthorized: no user context")
	}

	// Check if session has expired
	if time.Now().After(user.ExpiresAt) {
		return fmt.Errorf("unauthorized: session expired")
	}

	// Check user role
	if user.Role != "admin" {
		return fmt.Errorf("unauthorized: insufficient permissions")
	}

	// Proceed with protected operation
	fmt.Printf("Protected operation executed by %s\n", user.Username)
	return nil
}
