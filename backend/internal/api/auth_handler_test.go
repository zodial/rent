package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zodial/rent/backend/internal/api"
	"github.com/zodial/rent/backend/internal/auth"
	"github.com/zodial/rent/backend/internal/db"
	"github.com/zodial/rent/backend/internal/models"
)

func TestAuthRegisterAndToken(t *testing.T) {
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

	// setup server
	r := gin.New()
	jm := auth.NewJWTManager("testsecret", time.Hour)
	api.RegisterAuthRoutes(r, gdb, jm)

	// register admin
	reg := map[string]string{"name": "Bob", "email": "bob@example.com", "password": "pw", "role": "SUPER_ADMIN"}
	b, _ := json.Marshal(reg)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/admins", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// login
	login := map[string]string{"email": "bob@example.com", "password": "pw"}
	b2, _ := json.Marshal(login)
	req2 := httptest.NewRequest(http.MethodPost, "/api/v1/auth/token", bytes.NewReader(b2))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
	var resp map[string]string
	if err := json.Unmarshal(w2.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal token response: %v", err)
	}
	if _, ok := resp["token"]; !ok {
		t.Fatalf("token not found in response")
	}
}
