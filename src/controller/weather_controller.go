package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

func GetWeatherById(c *gin.Context) {
	id := c.Param("id")
	weather := map[string]interface{}{
		"id":       id,
		"temperature": "30°C",
		"humidity":   "60%",
		"condition":  "Sunny",
	}

	c.JSON(http.StatusOK, gin.H{"weather": weather})
}

func GetWeatherByCoordinates(c *gin.Context) {
	latitude := c.Param("lat")
	longitude := c.Param("lng")
	weather := map[string]interface{}{
		"latitude":  latitude,
		"longitude": longitude,
		"temperature": "30°C",
		"humidity":   "60%",
		"condition":  "Sunny",
	}

	c.JSON(http.StatusOK, gin.H{"weather": weather})
}