package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/zodial/rent/backend/internal/service"
)

// RegisterPropertyRoutes mounts property-related routes under /api/v1
func RegisterPropertyRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/properties", func(c *gin.Context) {
			var payload struct {
				Name    string `json:"name" binding:"required"`
				Address string `json:"address"`
			}
			if err := c.ShouldBindJSON(&payload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			p, err := service.CreateProperty(c.Request.Context(), db, "default", payload.Name, payload.Address)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, p)
		})
	}
}
