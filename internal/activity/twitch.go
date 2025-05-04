package activity

import "context"

type GetTwitchUserInput struct {
	Username string
}

type GetTwitchUserOutput struct {
	BroadcasterID string
}

func (a *Activity) GetTwitchUser(ctx context.Context, input GetTwitchUserInput) (*GetTwitchUserOutput, error) {
	broadcaster, err := a.GetUsers(input.Username)
	if err != nil {
		return nil, err
	}
	id := broadcaster.Users[0].BroadcasterID
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
