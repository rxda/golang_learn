package person

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name string
	Age int
	high float32
}


// GetDataAPIHealthHandler GET /health-dataapi to expose heathy check result of data API
func GetPersonHandler(c *gin.Context) {
	// do something to check heathy of data API

	person:=Person{Name:"HanRuida",Age:18,high:180}

	c.JSON(http.StatusOK, gin.H{
		"persons": person,
	})
}
