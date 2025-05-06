package credential

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// DefaultProvider implements CredentialProvider with in-memory storage
type DefaultProvider struct {
	users     map[string]*user
	mu        sync.RWMutex
	expiryAge time.Duration
}

type user struct {
	Username     string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}

// NewDefaultProvider creates a new DefaultProvider with the given session expiry duration
func NewDefaultProvider(expiryAge time.Duration) *DefaultProvider {
	return &DefaultProvider{
		users:     make(map[string]*user),
		expiryAge: expiryAge,
	}
}

// Authenticate implements CredentialProvider
func (p *DefaultProvider) Authenticate(ctx context.Context, username, password string) (*UserContext, error) {
	p.mu.RLock()
	user, exists := p.users[username]
	p.mu.RUnlock()

	if !exists {
		return nil, ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate a random user ID
	userID := make([]byte, 16)
	if _, err := rand.Read(userID); err != nil {
		return nil, err
	}

	now := time.Now()
	return &UserContext{
		UserID:    base64.URLEncoding.EncodeToString(userID),
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: now,
		ExpiresAt: now.Add(p.expiryAge),
	}, nil
}

// CreateUser implements CredentialProvider
func (p *DefaultProvider) CreateUser(ctx context.Context, username, password string, role string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exists := p.users[username]; exists {
		return ErrUserExists
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.users[username] = &user{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Role:         role,
		CreatedAt:    time.Now(),
	}

	return nil
}

// DeleteUser implements CredentialProvider
func (p *DefaultProvider) DeleteUser(ctx context.Context, username string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exists := p.users[username]; !exists {
		return ErrUserNotFound
	}

	delete(p.users, username)
	return nil
}

// UpdatePassword implements CredentialProvider
func (p *DefaultProvider) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	user, exists := p.users[username]
	if !exists {
		return ErrUserNotFound
	}

	// Verify old password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return ErrInvalidCredentials
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	return nil
}
