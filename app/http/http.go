package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init(env string) {
	r := gin.Default()

	// Endpoint to track clicks
	r.GET("/track-click", TrackClick)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
