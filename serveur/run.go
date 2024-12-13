package engine

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Run(jeu *Engine) {
    fmt.Println("Running server on port 8080")
    
    http.HandleFunc("/", jeu.Handler)
    router := gin.Default()
    router.GET("/api", func(c *gin.Context) {
        // Add your API handling logic here
        c.JSON(http.StatusOK, gin.H{"message": "API endpoint"})
    })

    router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
    })
    router.Run("localhost:8080")
}
