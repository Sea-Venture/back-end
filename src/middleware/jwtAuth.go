package middleware

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	publicKeys   map[string]*rsa.PublicKey
	lastKeyFetch time.Time
	mu           sync.Mutex
)

// fetchFirebasePublicKeys downloads and parses Google's public keys
func fetchFirebasePublicKeys() (map[string]*rsa.PublicKey, error) {
	resp, err := http.Get("https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var keyMap map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&keyMap); err != nil {
		return nil, err
	}

	keys := make(map[string]*rsa.PublicKey)
	for kid, certPEM := range keyMap {
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(certPEM))
		if err != nil {
			return nil, fmt.Errorf("error parsing public key: %v", err)
		}
		keys[kid] = key
	}

	return keys, nil
}

// getCachedPublicKey fetches/caches Firebase public keys
func getCachedPublicKey(kid string) (*rsa.PublicKey, error) {
	mu.Lock()
	defer mu.Unlock()

	if publicKeys == nil || time.Since(lastKeyFetch) > time.Hour {
		keys, err := fetchFirebasePublicKeys()
		if err != nil {
			return nil, err
		}
		publicKeys = keys
		lastKeyFetch = time.Now()
	}

	key, exists := publicKeys[kid]
	if !exists {
		return nil, fmt.Errorf("public key not found for kid: %s", kid)
	}
	return key, nil
}

// AuthMiddleware verifies Firebase ID token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			c.Abort()
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse token without verifying to extract "kid"
		parser := &jwt.Parser{}
		parsedToken, _, err := parser.ParseUnverified(idToken, jwt.MapClaims{})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		kid, ok := parsedToken.Header["kid"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing kid in token header"})
			c.Abort()
			return
		}

		publicKey, err := getCachedPublicKey(kid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get public key"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Validate issuer and audience
		if claims["iss"] != "https://securetoken.google.com/seaventure-e4ddc" ||
			claims["aud"] != "seaventure-e4ddc" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token issuer or audience mismatch"})
			c.Abort()
			return
		}

		// Add user info to context
		c.Set("userID", claims["user_id"])
		c.Set("email", claims["email"])

		c.Next()
	}
}