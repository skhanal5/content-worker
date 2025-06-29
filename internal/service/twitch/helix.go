package twitch

import (
	"fmt"
	"time"
)

const (
    users = "/users"
    clips = "/clips"
)


func (t *TwitchService) GetUsers(username string) (*UsersResponse, error) {
	resp, err := t.helixClient.R().
		SetResult(&UsersResponse{}).
        SetQueryParam("login", username).
		Get(users)
	if err != nil {
		return &UsersResponse{}, fmt.Errorf("failed to get users: %s", err)
	}
	return resp.Result().(*UsersResponse), nil
}


func (t *TwitchService) GetClips(broadcasterId string, daysBack int) (*ClipsResponse, error) {
    weekAgo := time.Now().AddDate(0,0, -1 * daysBack)
    weekAgoParam := weekAgo.Format(time.RFC3339)
    resp, err := t.helixClient.R().
        SetQueryParams(map[string]string{
            "broadcaster_id":  broadcasterId,
            "started_at": weekAgoParam,
        }).
		SetResult(&ClipsResponse{}).
		Get(clips)

	if err != nil {
		return &ClipsResponse{}, fmt.Errorf("failed to get clips: %s", err)
	}
	return resp.Result().(*ClipsResponse), nil
}

