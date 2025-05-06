package credential

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// DBUser represents the user model in the database
type DBUser struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

// DBAuthProvider implements authentication using a database
type DBAuthProvider struct {
	db        *gorm.DB
	expiryAge time.Duration
}

// NewDBAuthProvider creates a new database-backed authentication provider
func NewDBAuthProvider(db *gorm.DB, expiryAge time.Duration) *DBAuthProvider {
	// Auto-migrate the user table
	db.AutoMigrate(&DBUser{})

	return &DBAuthProvider{
		db:        db,
		expiryAge: expiryAge,
	}
}

// Authenticate implements database authentication
func (p *DBAuthProvider) Authenticate(ctx context.Context, username, password string) (*UserContext, error) {
	var user DBUser
	if err := p.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, ErrUserNotFound
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Create user context with session expiry
	now := time.Now()
	return &UserContext{
		UserID:    string(user.ID),
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: now,
		ExpiresAt: now.Add(p.expiryAge),
	}, nil
}

// CreateUser implements database user creation
func (p *DBAuthProvider) CreateUser(ctx context.Context, username, password string, role string) error {
	// Check if user exists
	var count int64
	if err := p.db.Model(&DBUser{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrUserExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user
	user := DBUser{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Role:         role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return p.db.Create(&user).Error
}

// DeleteUser implements database user deletion
func (p *DBAuthProvider) DeleteUser(ctx context.Context, username string) error {
	result := p.db.Where("username = ?", username).Delete(&DBUser{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}

// UpdatePassword implements database password update
func (p *DBAuthProvider) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	var user DBUser
	if err := p.db.Where("username = ?", username).First(&user).Error; err != nil {
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

	// Update password
	user.PasswordHash = string(hashedPassword)
	user.UpdatedAt = time.Now()
	return p.db.Save(&user).Error
}
