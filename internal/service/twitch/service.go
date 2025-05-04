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
		SetQueryParam("first", "1").
		SetResult(&UsersResponse{}).
		Get("/users?login=" + username)
	if err != nil {
		return &UsersResponse{}, err
	}
	return resp.Result().(*UsersResponse), nil
}

func (t *TwitchService) GetClips(broadcasterId string) (*ClipsResponse, error) {
	resp, err := t.client.R().
		SetQueryParam("first", "1").
		SetResult(&UsersResponse{}).
		Get("/clips?broadcaster_id=" + broadcasterId)
	if err != nil {
		return &ClipsResponse{}, err
	}
	return resp.Result().(*ClipsResponse), nil
}

func refreshAuthMiddleware(clientId string, clientSecret string) func (c *resty.Client, res *resty.Response) error {
	return func(c *resty.Client, res *resty.Response) error {
		if res.StatusCode() == http.StatusUnauthorized {
			authURL := fmt.Sprintf(oauthBaseURL)
			resp, err := c.R().
				SetQueryParam("client_id", clientId).
				SetQueryParam("client_secret", clientSecret).
				SetQueryParam("grant_type", "client_credentials").
				Post(authURL)
			if err != nil {
				return err
			}
			if resp.StatusCode() != http.StatusOK {
				return fmt.Errorf("failed to refresh auth token: %s", resp.String())
			}
			var authResponse AuthResponse
			if err := json.Unmarshal(resp.Body(), &authResponse); err != nil {
				return err
			}
			c.SetHeader("Authorization", "Bearer "+authResponse.AccessToken)
		}
		return nil	
	}
}

func NewTwitchService(baseURL string, clientId string, clientSecret string) *TwitchService {
	return &TwitchService{
		client: resty.New().
			SetTimeout(10*time.Second).
			SetBaseURL(baseURL).
			SetHeader("Client-ID", clientId).
			OnAfterResponse(refreshAuthMiddleware(clientId, clientSecret)),
	}
}
