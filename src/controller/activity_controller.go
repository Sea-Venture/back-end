package controller

import (

	"seaventures/src/models"
	"seaventures/src/service"

	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateActivity(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity created successfully"})
}

func GetAllActivities(c *gin.Context) {
	activities, err := service.GetAllActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func GetActivityByID(c *gin.Context) {
	id := c.Param("id")
	activity, err := service.GetActivityByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func UpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateActivity(id, &activity); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity updated successfully"})


}

func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteActivity(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}

func GetActivityDescriptionByActivityID(c *gin.Context) {
	id := c.Param("id")
	desc, err := service.GetActivityDescriptionByActivityID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": desc})
}




