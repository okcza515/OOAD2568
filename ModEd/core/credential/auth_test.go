package credential

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB creates a test database connection
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	return db
}

// TestAuthenticationFlow tests the complete authentication flow
func TestAuthenticationFlow(t *testing.T) {
	// Setup
	db := setupTestDB(t)
	provider := NewDBAuthProvider(db, 24*time.Hour)
	middleware := NewMiddleware(provider)
	ctx := context.Background()

	// Test user creation
	t.Run("Create User", func(t *testing.T) {
		err := middleware.CreateUser(ctx, "testuser", "testpass", "user")
		if err != nil {
			t.Errorf("Failed to create user: %v", err)
		}
	})

	// Test duplicate user creation
	t.Run("Create Duplicate User", func(t *testing.T) {
		err := middleware.CreateUser(ctx, "testuser", "testpass", "user")
		if err != ErrUserExists {
			t.Errorf("Expected ErrUserExists, got: %v", err)
		}
	})

	// Test authentication
	t.Run("Authenticate User", func(t *testing.T) {
		userCtx, err := middleware.Authenticate(ctx, "testuser", "testpass")
		if err != nil {
			t.Errorf("Authentication failed: %v", err)
		}
		if userCtx.Username != "testuser" {
			t.Errorf("Expected username 'testuser', got: %s", userCtx.Username)
		}
		if userCtx.Role != "user" {
			t.Errorf("Expected role 'user', got: %s", userCtx.Role)
		}
	})

	// Test wrong password
	t.Run("Authenticate Wrong Password", func(t *testing.T) {
		_, err := middleware.Authenticate(ctx, "testuser", "wrongpass")
		if err != ErrInvalidCredentials {
			t.Errorf("Expected ErrInvalidCredentials, got: %v", err)
		}
	})

	// Test password update
	t.Run("Update Password", func(t *testing.T) {
		err := middleware.UpdatePassword(ctx, "testuser", "testpass", "newpass")
		if err != nil {
			t.Errorf("Failed to update password: %v", err)
		}

		// Verify new password works
		_, err = middleware.Authenticate(ctx, "testuser", "newpass")
		if err != nil {
			t.Errorf("Authentication with new password failed: %v", err)
		}
	})

	// Test user deletion
	t.Run("Delete User", func(t *testing.T) {
		err := middleware.DeleteUser(ctx, "testuser")
		if err != nil {
			t.Errorf("Failed to delete user: %v", err)
		}

		// Verify user is deleted
		_, err = middleware.Authenticate(ctx, "testuser", "newpass")
		if err != ErrUserNotFound {
			t.Errorf("Expected ErrUserNotFound, got: %v", err)
		}
	})
}

// TestContextManagement tests the context management functions
func TestContextManagement(t *testing.T) {
	ctx := context.Background()
	userCtx := &UserContext{
		UserID:    "123",
		Username:  "testuser",
		Role:      "user",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	// Test adding context
	ctx = WithContext(ctx, userCtx)

	// Test retrieving context
	retrievedUser, ok := FromContext(ctx)
	if !ok {
		t.Error("Failed to retrieve user context")
	}
	if retrievedUser.Username != userCtx.Username {
		t.Errorf("Expected username %s, got %s", userCtx.Username, retrievedUser.Username)
	}
}

// TestSessionExpiry tests session expiry functionality
func TestSessionExpiry(t *testing.T) {
	db := setupTestDB(t)
	provider := NewDBAuthProvider(db, 1*time.Hour) // 1 hour expiry
	middleware := NewMiddleware(provider)
	ctx := context.Background()

	// Create and authenticate user
	err := middleware.CreateUser(ctx, "expireuser", "testpass", "user")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	userCtx, err := middleware.Authenticate(ctx, "expireuser", "testpass")
	if err != nil {
		t.Fatalf("Authentication failed: %v", err)
	}

	// Verify expiry time
	if time.Until(userCtx.ExpiresAt) > 1*time.Hour {
		t.Error("Session expiry time is too long")
	}
	if time.Until(userCtx.ExpiresAt) < 0 {
		t.Error("Session has already expired")
	}
}
