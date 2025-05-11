// Wrote by MEP-1001
package authentication

import (
	"context"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DBUser struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

type DBAuthProvider struct {
	db        *gorm.DB
	expiryAge time.Duration
}

func NewDBAuthProvider(db *gorm.DB, expiryAge time.Duration) *DBAuthProvider {
	db.AutoMigrate(&DBUser{})

	return &DBAuthProvider{
		db:        db,
		expiryAge: expiryAge,
	}
}

func (p *DBAuthProvider) Authenticate(ctx context.Context, username, password string) (*UserContext, error) {
	var user DBUser
	if err := p.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidAuthentications
	}

	now := time.Now()
	return &UserContext{
		UserID:    strconv.FormatUint(uint64(user.ID), 10),
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: now,
		ExpiresAt: now.Add(p.expiryAge),
	}, nil
}

func (p *DBAuthProvider) CreateUser(ctx context.Context, username, password string, role string) error {
	var count int64
	if err := p.db.Model(&DBUser{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := DBUser{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Role:         role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return p.db.Create(&user).Error
}

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

func (p *DBAuthProvider) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	var user DBUser
	if err := p.db.Where("username = ?", username).First(&user).Error; err != nil {
		return ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return ErrInvalidAuthentications
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	user.UpdatedAt = time.Now()
	return p.db.Save(&user).Error
}

func (p *DBAuthProvider) ListUsers(ctx context.Context) ([]UserContext, error) {
	var dbUsers []DBUser
	if err := p.db.Find(&dbUsers).Error; err != nil {
		return nil, err
	}

	userContexts := make([]UserContext, len(dbUsers))
	for i, user := range dbUsers {
		userContexts[i] = UserContext{
			UserID:    strconv.FormatUint(uint64(user.ID), 10),
			Username:  user.Username,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			ExpiresAt: user.CreatedAt.Add(p.expiryAge),
		}
	}

	return userContexts, nil
}

func (p *DBAuthProvider) GetCurrentUser(ctx context.Context) (*UserContext, error) {
	userCtx, ok := ctx.Value("user").(*UserContext)
	if !ok || userCtx == nil {
		return nil, ErrUnauthorized
	}
	return userCtx, nil
}

func (p *DBAuthProvider) UpdateUserRole(ctx context.Context, username, newRole string) error {
	var user DBUser
	if err := p.db.Where("username = ?", username).First(&user).Error; err != nil {
		return ErrUserNotFound
	}

	user.Role = newRole
	return p.db.Save(&user).Error
}
