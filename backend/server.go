package backend

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StartGinServer starts the Gin HTTP server
func StartGinServer(port string) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Printf("Starting Gin Server on %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start Gin server: %v", err)
	}
}
