package data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDataAPIHealthHandler GET /health-dataapi to expose heathy check result of data API
func GetDataAPIHealthHandler(c *gin.Context) {
	// do something to check heathy of data API
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Data API is alive",
	})
}
