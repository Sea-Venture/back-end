package controller

import (
	"net/http"
	"seaventures/src/service"

	"github.com/gin-gonic/gin"
)

// GET /api/user/forecast/advanced?beach=Galle
func GetAdvancedForecastHandler(c *gin.Context) {
    beach := c.Query("beach")
    if beach == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Beach parameter is required"})
        return
    }

    result, err := service.GetAdvancedForecast(beach)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)
}