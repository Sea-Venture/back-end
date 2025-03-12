package config

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var OAuthConfig = &clientcredentials.Config{
	ClientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
	ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
	TokenURL:     os.Getenv("KEYCLOAK_TOKEN_URL"),
}


func GetToken() (*oauth2.Token, error) {
	ctx := context.Background()
	token, err := OAuthConfig.Token(ctx)
	if err != nil {
		log.Printf("Unable to retrieve token: %v", err)
		return nil, err
	}
	return token, nil
}
