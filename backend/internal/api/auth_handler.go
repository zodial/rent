package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/zodial/rent/backend/internal/auth"
	"github.com/zodial/rent/backend/internal/service"
)

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB, jwtManager *auth.JWTManager) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/admins", func(c *gin.Context) {
			var payload struct {
				Name     string `json:"name" binding:"required"`
				Email    string `json:"email" binding:"required"`
				Password string `json:"password" binding:"required"`
				Role     string `json:"role" binding:"required"`
			}
			if err := c.ShouldBindJSON(&payload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			admin, err := service.RegisterAdmin(c.Request.Context(), db, "default", payload.Name, payload.Email, payload.Password, payload.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, admin)
		})

		v1.POST("/auth/token", func(c *gin.Context) {
			var payload struct {
				Email    string `json:"email" binding:"required"`
				Password string `json:"password" binding:"required"`
			}
			if err := c.ShouldBindJSON(&payload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			admin, err := service.Authenticate(c.Request.Context(), db, "default", payload.Email, payload.Password)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
				return
			}
			// generate token
			tok, err := jwtManager.GenerateToken(admin.ID, admin.TenantID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": tok})
		})
	}
}
