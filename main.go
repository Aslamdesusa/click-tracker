// app.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourapp/http"
)

func main() {
	r := gin.Default()

	// Endpoint to track clicks
	r.GET("/track-click", http.TrackClick)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
