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

type ClipOutput struct {	
	ID string
	Title string
	URL   string
	ViewCount int
	Duration float32
	CreatedAt string
	ThumbnailURL string
}

type GetClipsOutput struct {
	Clips []ClipOutput 
}

func (a *Activity) GetClipsFromUser(ctx context.Context, input GetClipsInput) (*GetClipsOutput, error) {
	clips, err := a.GetClips(input.BroadcasterID)
	if err != nil {
		return nil, err
	}
	if clips == nil {	
		return nil, fmt.Errorf("no clips found for broadcaster id: %s", input.BroadcasterID)
	}
	
	clipsOutput := &GetClipsOutput{
		Clips: make([]ClipOutput, len(clips.Clips)),
	}
	for _, clip := range clips.Clips {
		clipsOutput.Clips = append(clipsOutput.Clips, ClipOutput{
			ID: clip.ID,
			Title: clip.Title,
			URL: clip.URL,
			ViewCount: clip.ViewCount,
			Duration: clip.Duration,
			CreatedAt: clip.CreatedAt,
			ThumbnailURL: clip.ThumbnailURL,
		})
	}

	return clipsOutput, nil
}
