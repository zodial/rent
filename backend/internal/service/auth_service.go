package service

import (
	"context"
	"errors"

	"github.com/zodial/rent/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// RegisterAdmin creates an administrator with a bcrypt-hashed password.
func RegisterAdmin(ctx context.Context, db *gorm.DB, tenantID, name, email, password, role string) (*models.Administrator, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin := &models.Administrator{
		TenantID:     tenantID,
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
		Role:         role,
	}
	if err := db.WithContext(ctx).Create(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

// Authenticate verifies email+password and returns the admin if valid.
func Authenticate(ctx context.Context, db *gorm.DB, tenantID, email, password string) (*models.Administrator, error) {
	var admin models.Administrator
	if err := db.WithContext(ctx).Where("tenant_id = ? AND email = ?", tenantID, email).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}
	return &admin, nil
}
