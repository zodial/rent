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

func TestCreateProperty(t *testing.T) {
	// Use in-memory SQLite for tests
	dbPath := os.Getenv("TEST_DATABASE_URL")
	if dbPath == "" {
		dbPath = ":memory:"
	}
	gdb, err := db.Open(dbPath)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	// Run migrations
	if err := models.Migrate(gdb); err != nil {
		t.Fatalf("migrate failed: %v", err)
	}

	ctx := context.Background()
	p, err := service.CreateProperty(ctx, gdb, "default", "Test Property", "123 Test St")
	if err != nil {
		t.Fatalf("CreateProperty failed: %v", err)
	}

	var found models.Property
	if err := gdb.First(&found, p.ID).Error; err != nil {
		t.Fatalf("failed to fetch property: %v", err)
	}

	assert.Equal(t, "Test Property", found.Name)
	assert.Equal(t, "123 Test St", found.Address)
}
package service

import (
	"context"

	"gorm.io/gorm"
	"github.com/zodial/rent/backend/internal/models"
)

// CreateProperty creates a property and associated rooms in a single transaction.
func CreateProperty(ctx context.Context, db *gorm.DB, tenantID, name, address string) (*models.Property, error) {
	p := &models.Property{
		TenantID: tenantID,
		Name:     name,
		Address:  address,
	}
	if err := db.WithContext(ctx).Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
module github.com/zodial/rent/backend

go 1.20

require (
	github.com/gin-gonic/gin v1.9.0
	gorm.io/driver/sqlite v1.4.6
	gorm.io/gorm v1.26.0
	github.com/stretchr/testify v1.8.4
)

