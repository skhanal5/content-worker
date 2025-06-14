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

type GetClipSlugsInput struct {
	BroadcasterID string
}

type GetClipSlugsOutput struct {
	ClipIds []string 
}

func (a *Activity) GetClipSlugs(ctx context.Context, input GetClipSlugsInput) (*GetClipSlugsOutput, error) {
	clips, err := a.GetClips(input.BroadcasterID)
	if err != nil {
		return nil, err
	}
	if clips == nil {	
		return nil, fmt.Errorf("no clips found for broadcaster id: %s", input.BroadcasterID)
	}
	
	clipsOutput := &GetClipSlugsOutput{
		ClipIds: []string{},
	}
	for _, clip := range clips.Clips {
		if clip.Duration < 15 || clip.ID == ""{
			continue
		}
		clipsOutput.ClipIds = append(clipsOutput.ClipIds, clip.ID)
	}

	return clipsOutput, nil
}

type GetDownloadLinksInput struct {
	ClipIds []string
}

type GetDownloadLinksOutput struct {
	DownloadLinks []string 
}

func (a *Activity) GetDownloadLinks(ctx context.Context,  input GetDownloadLinksInput) (*GetDownloadLinksOutput, error) {
	output := &GetDownloadLinksOutput{
		DownloadLinks: []string{},
	}
	for _, clipId := range input.ClipIds {
		res, err := a.GetDownloadLink(clipId)
		if err != nil {
			return nil, err
		}
		output.DownloadLinks = append(output.DownloadLinks, res)
	}
	return output, nil
}