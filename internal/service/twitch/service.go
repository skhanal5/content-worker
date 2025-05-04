package twitch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	oauthBaseURL = "https://id.twitch.tv/oauth2/token"
)

type TwitchManager interface {
	GetUsers(broadcasterId string) (*UsersResponse, error)
	GetClips(username string) (*ClipsResponse, error)
}

type TwitchService struct {
	client  *resty.Client
	authURL string
}

func (t *TwitchService) GetUsers(username string) (*UsersResponse, error) {
	resp, err := t.client.R().
		SetResult(&UsersResponse{}).
		Get("/users?login=" + username)
	if err != nil {
		return &UsersResponse{}, err
	}
	return resp.Result().(*UsersResponse), nil
}

func (t *TwitchService) GetClips(broadcasterId string) (*ClipsResponse, error) {
	resp, err := t.client.R().
		SetResult(&ClipsResponse{}).
		Get("/clips?broadcaster_id=" + broadcasterId)
	if err != nil {
		return &ClipsResponse{}, err
	}
	return resp.Result().(*ClipsResponse), nil
}

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

func NewTwitchService(baseURL string, clientId string, clientSecret string) *TwitchService {
	return &TwitchService{
		client: resty.New().
			SetTimeout(10*time.Second).
			SetBaseURL(baseURL).
			OnBeforeRequest(addAuthHeaderMiddleware(clientId, clientSecret)),
	}
}
