package controller

import (
    "net/http"

    "seaventures/src/models"
    "seaventures/src/service"

    "github.com/gin-gonic/gin"
)

func CreateGuide(c *gin.Context) {
	var guide models.Guide

	if err := c.ShouldBindJSON(&guide); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateGuide(&guide); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guide created successfully"})

}

func GetAllGuides(c *gin.Context) {
	guides, err := service.GetAllGuides()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guides)
}

func GetGuideByID(c *gin.Context) {
	id := c.Param("id")
	guide, err := service.GetGuideByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guide)
}

func UpdateGuide(c *gin.Context) {
	id := c.Param("id")
	var guide models.Guide

	if err := c.ShouldBindJSON(&guide); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateGuide(id, &guide); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guide updated successfully"})
}

func DeleteGuide(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteGuide(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guide deleted successfully"})
}

func GetGuideByBeachID(c *gin.Context) {
	id := c.Param("id")
	guides, err := service.GetGuideByBeachID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guides)
}

func GetGuideByActivityID(c *gin.Context) {
	acid := c.Param("acid")
	guides, err := service.GetGuideByActivityID(acid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guides)
}

func GetGuideActivitiesByBeachIDAndActivityID(c *gin.Context) {
	beachID := c.Param("beach_id")
	activityID := c.Param("activity_id")
	guides, err := service.GetGuideActivitiesByBeachIDAndActivityID(beachID, activityID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guides)
}
