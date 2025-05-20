package middleware

import (
	"context"
	"net/http"
	"strings"

	"seaventures/src/config"

	"github.com/gin-gonic/gin"
)

func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}
		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		client, err := config.FirebaseApp.Auth(context.Background())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Firebase Auth client"})
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			return
		}

		c.Set("firebase_uid", token.UID)
		c.Set("email", token.Claims["email"])
		c.Next()
	}
}
