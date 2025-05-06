package credential

import (
	"context"
	"errors"
	"time"
)

// Common errors
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
)

// UserContext represents the authenticated user's context
type UserContext struct {
	UserID    string
	Username  string
	Role      string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// CredentialProvider defines the interface for credential management
type CredentialProvider interface {
	// Authenticate validates user credentials and returns a UserContext if successful
	Authenticate(ctx context.Context, username, password string) (*UserContext, error)

	// CreateUser creates a new user with the given credentials
	CreateUser(ctx context.Context, username, password string, role string) error

	// DeleteUser removes a user from the system
	DeleteUser(ctx context.Context, username string) error

	// UpdatePassword changes a user's password
	UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error
}

// Middleware provides a base implementation of the CredentialProvider interface
type Middleware struct {
	provider CredentialProvider
}

// NewMiddleware creates a new credential middleware with the given provider
func NewMiddleware(provider CredentialProvider) *Middleware {
	return &Middleware{
		provider: provider,
	}
}

// Authenticate wraps the provider's Authenticate method
func (m *Middleware) Authenticate(ctx context.Context, username, password string) (*UserContext, error) {
	return m.provider.Authenticate(ctx, username, password)
}

// CreateUser wraps the provider's CreateUser method
func (m *Middleware) CreateUser(ctx context.Context, username, password string, role string) error {
	return m.provider.CreateUser(ctx, username, password, role)
}

// DeleteUser wraps the provider's DeleteUser method
func (m *Middleware) DeleteUser(ctx context.Context, username string) error {
	return m.provider.DeleteUser(ctx, username)
}

// UpdatePassword wraps the provider's UpdatePassword method
func (m *Middleware) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	return m.provider.UpdatePassword(ctx, username, oldPassword, newPassword)
}

// WithContext adds the UserContext to the given context
func WithContext(ctx context.Context, userCtx *UserContext) context.Context {
	return context.WithValue(ctx, "user", userCtx)
}

// FromContext retrieves the UserContext from the given context
func FromContext(ctx context.Context) (*UserContext, bool) {
	user, ok := ctx.Value("user").(*UserContext)
	return user, ok
}
