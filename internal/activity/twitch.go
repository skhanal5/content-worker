package activity

type GetTwitchUserInput struct {
	Username string 
}

type GetTwitchUserOutput struct {
	BroadcasterID string
}


func (a *Activity) GetTwitchUser(input GetTwitchUserInput) (*GetTwitchUserOutput, error) {
	broadcaster, err := a.twitchService.GetUsers(input.Username)
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


func (a *Activity) GetClipsFromUser(input GetClipsInput) ([]string, error) {
	clips, err := a.twitchService.GetClips(input.BroadcasterID)
	if err != nil {
		return nil, err
	}

	var clipURLs []string
	for _, clip := range clips.Clips {
		clipURLs = append(clipURLs, clip.URL)
	}

	return clipURLs, nil
}