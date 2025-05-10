package authentication

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory database: %v", err)
	}
	db.AutoMigrate(&DBUser{})
	return db
}

func TestUserRegistrationAndLogin(t *testing.T) {
	db := setupTestDB(t)
	provider := NewDBAuthProvider(db, 1*time.Hour)
	middleware := NewMiddleware(provider)

	ctx := context.Background()
	username := "testuser"
	password := "secure123"
	role := "admin"

	if err := middleware.CreateUser(ctx, username, password, role); err != nil {
		t.Fatalf("failed to register user: %v", err)
	}

	if err := middleware.CreateUser(ctx, username, password, role); err != ErrUserExists {
		t.Errorf("expected ErrUserExists, got %v", err)
	}

	userCtx, err := middleware.Authenticate(ctx, username, password)
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}
	if userCtx.Username != username || userCtx.Role != role {
		t.Errorf("unexpected user context: %+v", userCtx)
	}

	ctxWithUser := WithContext(context.Background(), userCtx)
	if err := RequireAdmin(ctxWithUser); err != nil {
		t.Errorf("expected admin access, got error: %v", err)
	}
}

func TestUnauthorizedAccess(t *testing.T) {
	ctx := context.Background()

	if err := RequireAdmin(ctx); err != ErrUnauthorized {
		t.Errorf("expected ErrUnauthorized for missing context, got: %v", err)
	}

	ctx = WithContext(ctx, &UserContext{
		Username: "user1",
		Role:     "user",
	})

	if err := RequireAdmin(ctx); err != ErrUnauthorized {
		t.Errorf("expected ErrUnauthorized for non-admin, got: %v", err)
	}
}
