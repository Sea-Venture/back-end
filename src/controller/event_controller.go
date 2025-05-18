package controller

import (
	"net/http"
	"seaventures/src/models"
	"seaventures/src/service"

	"github.com/gin-gonic/gin"
)


func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully"})
}

func GetEvents(c *gin.Context) {
	events, err := service.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

func GetEventByID(c *gin.Context) {
	id := c.Param("id")
	event, err := service.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateEvent(id, &event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func GetEventByActivityID(c *gin.Context) {
	activityID := c.Param("id")
	events, err := service.GetEventByActivityID(activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

func GetEventByLocationID(c *gin.Context) {
	locationID := c.Param("id")
	events, err := service.GetEventByLocationID(locationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

func GetEventByLocationIDAndActivityID(c *gin.Context) {
	locationID := c.Param("location_id")
	activityID := c.Param("activity_id")
	events, err := service.GetEventByLocationIDAndActivityID(locationID, activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}




