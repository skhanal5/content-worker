package twitch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

func addAuthHeaderMiddleware(clientId string, clientSecret string) func(c *resty.Client, req *resty.Request) error {
	return func(c *resty.Client, req *resty.Request) error {
		token, err := tokenProvider(clientId, clientSecret)()
		if err != nil {
			return fmt.Errorf("failed to fetch token: %w", err)
		}

		req.SetHeader("Authorization", "Bearer "+token)
		req.SetHeader("Client-ID", clientId)

		return nil
	}
}

func tokenProvider(clientId string, clientSecret string) func() (string, error) {
	var token string
	var tokenExpiry time.Time

	return func() (string, error) {
		if token != "" && time.Now().Before(tokenExpiry) {
			return token, nil
		}

		resp, err := resty.New().R().
			SetQueryParam("client_id", clientId).
			SetQueryParam("client_secret", clientSecret).
			SetQueryParam("grant_type", "client_credentials").
			Post(oauthBaseURL)
		if err != nil {
			return "", fmt.Errorf("failed to refresh token: %w", err)
		}

		if resp.StatusCode() != http.StatusOK {
			return "", fmt.Errorf("failed to refresh token: %s", resp.String())
		}

		var authResponse AuthResponse
		if err := json.Unmarshal(resp.Body(), &authResponse); err != nil {
			return "", fmt.Errorf("failed to parse token response: %w", err)
		}

		token = authResponse.AccessToken
		tokenExpiry = time.Now().Add(time.Duration(authResponse.ExpiresIn) * time.Second)

		return token, nil
	}
}
