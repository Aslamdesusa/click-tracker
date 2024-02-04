// http/controller.go
package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type ClickResponse struct {
	Message string `json:"message"`
	Data    struct {
		UserAgent string          `json:"user_agent"`
		GeoInfo   json.RawMessage `json:"geo_info"`
	} `json:"data"`
}

func TrackClick(c *gin.Context) {
	// Extract User-Agent from the request headers
	userAgent := c.GetHeader("User-Agent")

	// Get the user's IP address from the request
	ip := c.ClientIP()

	// Get geo information based on the user's IP
	geoInfo, err := getGeoInfo(ip)
	if err != nil {
		fmt.Println("Error getting geo info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Construct the JSON response
	response := ClickResponse{
		Message: "Click tracked successfully",
		Data: struct {
			UserAgent string          `json:"user_agent"`
			GeoInfo   json.RawMessage `json:"geo_info"`
		}{
			UserAgent: userAgent,
			GeoInfo:   geoInfo,
		},
	}

	logClick(response)

	c.JSON(http.StatusOK, response)
}

func getGeoInfo(ip string) (json.RawMessage, error) {
	// Use the IPinfo API to get geo information based on IP
	apiKey := os.Getenv("IP_INFO_API_KEY")
	url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, apiKey)

	resp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(resp.Body()), nil
}

func logClick(response ClickResponse) {
	// Implement your logging logic here (e.g., store in a database)
	fmt.Printf("UserAgent: %s, GeoInfo: %s\n", response.Data.UserAgent, response.Data.GeoInfo)
}
