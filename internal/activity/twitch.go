package activity

import (
	"context"
	"fmt"
)

type GetTwitchUserInput struct {
	Username string
}

type GetTwitchUserOutput struct {
	BroadcasterID string
}

func (a *Activity) GetTwitchUser(ctx context.Context, input GetTwitchUserInput) (*GetTwitchUserOutput, error) {
	response, err := a.GetUsers(input.Username)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, fmt.Errorf("no users found for username: %s", input.Username)
	}
	id := response.Users[0].Id
	return &GetTwitchUserOutput{
		BroadcasterID: id,
	}, nil
}

type GetClipsInput struct {
	BroadcasterID string
}

type GetClipsOutput struct {
	ClipURLs []string
}

func (a *Activity) GetClipsFromUser(ctx context.Context, input GetClipsInput) ([]string, error) {
	clips, err := a.GetClips(input.BroadcasterID)
	if err != nil {
		return nil, err
	}

	var clipURLs []string
	for _, clip := range clips.Clips {
		clipURLs = append(clipURLs, clip.URL)
	}

	return clipURLs, nil
}
