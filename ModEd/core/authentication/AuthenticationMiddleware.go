package authentication

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidAuthentications = errors.New("invalid authentications")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserExists             = errors.New("user already exists")
	ErrUnauthorized           = errors.New("unauthorized: requires admin role")
)

type UserContext struct {
	UserID    string
	Username  string
	Role      string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type AuthenticationProvider interface {
	Authenticate(ctx context.Context, username, password string) (*UserContext, error)

	CreateUser(ctx context.Context, username, password string, role string) error

	DeleteUser(ctx context.Context, username string) error

	UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error

	ListUsers(ctx context.Context) ([]UserContext, error)

	GetCurrentUser(ctx context.Context) (*UserContext, error)

	UpdateUserRole(ctx context.Context, username, role string) error
}

type Middleware struct {
	provider AuthenticationProvider
}

func NewMiddleware(provider AuthenticationProvider) *Middleware {
	return &Middleware{
		provider: provider,
	}
}

func (m *Middleware) Authenticate(ctx context.Context, username, password string) (*UserContext, error) {
	return m.provider.Authenticate(ctx, username, password)
}

func (m *Middleware) CreateUser(ctx context.Context, username, password string, role string) error {
	return m.provider.CreateUser(ctx, username, password, role)
}

func (m *Middleware) DeleteUser(ctx context.Context, username string) error {
	return m.provider.DeleteUser(ctx, username)
}

func (m *Middleware) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	return m.provider.UpdatePassword(ctx, username, oldPassword, newPassword)
}

func (m *Middleware) ListUsers(ctx context.Context) ([]UserContext, error) {
	return m.provider.ListUsers(ctx)
}
func (m *Middleware) GetCurrentUser(ctx context.Context) (*UserContext, error) {
	return m.provider.GetCurrentUser(ctx)
}

func (m *Middleware) UpdateUserRole(ctx context.Context, username, role string) error {
	return m.provider.UpdateUserRole(ctx, username, role)
}

func WithContext(ctx context.Context, userCtx *UserContext) context.Context {
	return context.WithValue(ctx, "user", userCtx)
}

func FromContext(ctx context.Context) (*UserContext, bool) {
	user, ok := ctx.Value("user").(*UserContext)
	return user, ok
}

func RequireAdmin(ctx context.Context) error {
	user, ok := FromContext(ctx)
	if !ok {
		return ErrUnauthorized
	}
	if user.Role != "admin" {
		return ErrUnauthorized
	}
	return nil
}
