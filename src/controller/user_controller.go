package controller

import (
	"net/http"
	"os"
	"seaventures/src/helpers"
	"seaventures/src/models"
	"seaventures/src/service"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := helpers.GenerateJWT(user.ID, user.Role, secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "token": token, "user": user})
}

func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.Login(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := helpers.GenerateJWT(user.ID, user.Role, secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user": user})
}

func AddProfilePic(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var profilePicData struct {
		ProfilePic string `json:"profile_pic"`
	}
	if err := c.ShouldBindJSON(&profilePicData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.GetUserById(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	err = service.UpdateProfilePic(user, profilePicData.ProfilePic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update profile picture"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile picture updated successfully"})
}

func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You have accessed a protected endpoint"})
}
