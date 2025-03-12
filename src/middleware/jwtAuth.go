package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// AuthMiddleware validates the OAuth2 token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			c.Abort()
			return
		}

		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: tokenString})
		token, err := tokenSource.Token()
		if err != nil || !token.Valid() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
