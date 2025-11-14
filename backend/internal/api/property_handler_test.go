package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zodial/rent/backend/internal/api"
	"github.com/zodial/rent/backend/internal/db"
	"github.com/zodial/rent/backend/internal/models"
)

func TestCreatePropertyHandler(t *testing.T) {
	// Setup in-memory DB
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

	r := gin.New()
	api.RegisterPropertyRoutes(r, gdb)

	payload := map[string]string{"name": "API Property", "address": "1 API Rd"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/properties", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.Property
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	assert.Equal(t, "API Property", resp.Name)
}
