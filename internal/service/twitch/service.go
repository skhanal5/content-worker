package twitch

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type TwitchManager interface {
	GetClipInformation(clipSlug string) (*ClipMetadataResponse, error)
	GetUserClips(user string, limit int, filter string) (*UserClipsResponse, error)
}

const (
	gqlUrl = "https://gql.twitch.tv"
)

type TwitchService struct {
	gqlClient *resty.Client
}

func NewTwitchService(gqlClientId string) *TwitchService {
	return &TwitchService{
		gqlClient: resty.New().
			SetTimeout(10*time.Second).
			SetBaseURL(gqlUrl).
			SetDebug(true).
			SetHeader("Client-ID", gqlClientId),
	}
}

func (t *TwitchService) GetClipInformation(clipSlug string) (*ClipMetadataResponse, error) {
	queryBody := []map[string]interface{}{
		{
			"operationName": "VideoAccessToken_Clip",
			"variables": map[string]interface{}{
				"slug": clipSlug,
			},
			"extensions": map[string]interface{}{
				"persistedQuery": map[string]interface{}{
					"version":    1,
					"sha256Hash": "36b89d2507fce29e5ca551df756d27c1cfe079e2609642b4390aa4c35796eb11",
				},
			},
		},
	}

	var gqlResponses []ClipMetadataResponse // was type array
	_, err := t.gqlClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(queryBody).
		SetResult(&gqlResponses).
		Post("/gql")

	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	return &gqlResponses[0], nil

	// clip := gqlResponses[0].Data.Clip
	// qualities := clip.VideoQualities
	// bestQualityURL := qualities[0].SourceURL

	// u, err := url.Parse(bestQualityURL)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to parse sourceURL: %v", err)
	// }

	// q := u.Query()
	// q.Set("sig", clip.PlaybackAccessToken.Signature)
	// q.Set("token", clip.PlaybackAccessToken.Value)
	// u.RawQuery = q.Encode()

	// return u.String(), nil
}

func (t *TwitchService) GetUserClips(streamer string, limit int, filter string) (*UserClipsResponse, error) {
	queryBody := []map[string]interface{}{
		{
			"operationName": "ClipsCards__User",
			"variables": map[string]interface{}{
				"login": streamer,
				"limit": limit,
				"criteria": map[string]interface{}{
					"filter": filter,
				},
			},
			"extensions": map[string]interface{}{
				"persistedQuery": map[string]interface{}{
					"version":    1,
					"sha256Hash": "4eb8f85fc41a36c481d809e8e99b2a32127fdb7647c336d27743ec4a88c4ea44",
				},
			},
		},
	}

	var response []UserClipsResponse
	_, err := t.gqlClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(queryBody).
		SetResult(&response).
		Post("/gql")

	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	if len(response[0].Errors) > 0 {
		return nil, fmt.Errorf("error from Twitch API: %s", response[0].Errors[0].Message)
	}

	return &response[0], nil
}
