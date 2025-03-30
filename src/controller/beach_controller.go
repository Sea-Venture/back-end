package controller

import (
	"net/http"
    "seaventures/src/models"
    "seaventures/src/service"
    "github.com/gin-gonic/gin"
)

func CreateBeach(c *gin.Context) {
	var beach models.Beach
	if err := c.ShouldBindJSON(&beach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateBeach(&beach); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Beach created successfully"})
}

func GetAllBeaches(c *gin.Context) {
	beaches, err := service.GetAllBeaches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, beaches)
}

func GetBeachByID(c *gin.Context) {
	id := c.Param("id")
	beach, err := service.GetBeachByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, beach)
}

func UpdateBeach(c *gin.Context) {
	id := c.Param("id")
	var beach models.Beach
	if err := c.ShouldBindJSON(&beach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateBeach(id, &beach); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Beach updated successfully"})
}

func DeleteBeach(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteBeach(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Beach deleted successfully"})
}

func GetBeachesByLocationID(c *gin.Context) {
	id := c.Param("id")
	beaches, err := service.GetBeachesByLocationID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, beaches)
}

func GetBeachDescriptionByBeachID(c *gin.Context) {
	id := c.Param("id")
	description, err := service.GetBeachDescriptionByBeachID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": description})
}

