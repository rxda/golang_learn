package main

import (
    dataapi "ginwithmodule/pkg/myapi/data"
    person "ginwithmodule/pkg/myapi/person"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/health", GetHealthHandler)
    router.GET("/health-dataapi", dataapi.GetDataAPIHealthHandler)
    router.GET("/person",person.GetPersonHandler)
    s := &http.Server{
        Addr:    ":8000",
        Handler: router,
    }
    s.ListenAndServe()
}

// GetHealthHandler - GET /health to expose service health
func GetHealthHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "code":    0,
        "message": "Service is alive!",
    })
}