package twitch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	oauthBaseURL = "https://id.twitch.tv/oauth2/token?client_id=%s&grant_type=client_credentials"
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

func (t *TwitchService) GetClips(broadcasterId string) (*UsersResponse, error) {
	resp, err := t.client.R().
		SetQueryParam("first", "1").
		SetResult(&UsersResponse{}).
		Get("/clips?broadcaster_id=" + broadcasterId)
	if err != nil {
		return &UsersResponse{}, err
	}
	return resp.Result().(*UsersResponse), nil
}

func RefreshAuthMiddleware(c *resty.Client, res *resty.Response) error {
	if res.StatusCode() == http.StatusUnauthorized {
		authURL := fmt.Sprintf(oauthBaseURL, c.Header.Get("Client-ID"))
		resp, err := c.R().Post(authURL)
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

func NewTwitchClient(baseURL string, clientId string) *TwitchService {
	return &TwitchService{
		client: resty.New().
			SetTimeout(10*time.Second).
			SetBaseURL(baseURL).
			SetHeader("Client-ID", clientId).
			OnAfterResponse(RefreshAuthMiddleware),
	}
}
