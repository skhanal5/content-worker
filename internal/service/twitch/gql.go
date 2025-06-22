package twitch

import (
	"fmt"
	"net/url"
)

func (t *TwitchService) GetDownloadLink(clipSlug string) (string, error) {
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

	var gqlResponses []GraphQLResponse
	_, err :=t.gqlClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(queryBody).
		SetResult(&gqlResponses).
		Post("/gql")

    if err != nil {
        return "", fmt.Errorf("request failed: %v", err)
    }

	clip := gqlResponses[0].Data.Clip
	qualities := clip.VideoQualities
	bestQualityURL := qualities[0].SourceURL
	
	u, err := url.Parse(bestQualityURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse sourceURL: %v", err)
	}

	q := u.Query()
	q.Set("sig", clip.PlaybackAccessToken.Signature)
	q.Set("token", clip.PlaybackAccessToken.Value)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

