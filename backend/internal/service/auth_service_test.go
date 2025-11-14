package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zodial/rent/backend/internal/db"
	"github.com/zodial/rent/backend/internal/models"
	"github.com/zodial/rent/backend/internal/service"
)

func TestRegisterAndAuthenticateAdmin(t *testing.T) {
	dbPath := os.Getenv("TEST_DATABASE_URL")
	if dbPath == "" {
		dbPath = ":memory:"
	}
	gdb, err := db.Open(dbPath)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	if err := models.Migrate(gdb); err != nil {
		t.Fatalf("migrate failed: %v", err)
	}

	ctx := context.Background()
	admin, err := service.RegisterAdmin(ctx, gdb, "default", "Alice", "alice@example.com", "s3cr3t", "SUPER_ADMIN")
	if err != nil {
		t.Fatalf("RegisterAdmin failed: %v", err)
	}
	assert.Equal(t, "Alice", admin.Name)

	auth, err := service.Authenticate(ctx, gdb, "default", "alice@example.com", "s3cr3t")
	if err != nil {
		t.Fatalf("Authenticate failed: %v", err)
	}
	assert.Equal(t, admin.ID, auth.ID)

	_, err = service.Authenticate(ctx, gdb, "default", "alice@example.com", "wrongpwd")
	assert.Equal(t, service.ErrInvalidCredentials, err)
}
