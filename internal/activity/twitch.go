package activity

import (
	"clip-farmer-workflow/internal/service/twitch"
	"context"
	"fmt"
	"sort"
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
	DaysAgo       int
	TopN int
}

type GetClipSlugsOutput struct {
	ClipIds []string
}

func (a *Activity) GetClipSlugs(ctx context.Context, input GetClipSlugsInput) (*GetClipSlugsOutput, error) {
	topN := input.TopN
	
	clips, err := a.GetClips(input.BroadcasterID, input.DaysAgo)
	if err != nil {
		return nil, err
	}
	if clips == nil {
		return nil, fmt.Errorf("no clips found for broadcaster id: %s", input.BroadcasterID)
	}

	filteredClips := []twitch.Clip{} 
	for _, clip := range clips.Clips {
		if clip.Duration >= 15 && clip.ID != "" {
			filteredClips = append(filteredClips, clip)
		}
	}

	sort.Slice(filteredClips, func(i, j int) bool {
		return filteredClips[i].ViewCount > filteredClips[j].ViewCount
	})

	if input.TopN > len(filteredClips) {
		topN = len(filteredClips)
	}
	topClips := filteredClips[:topN]

	clipsOutput := &GetClipSlugsOutput{
		ClipIds: make([]string, 0, topN),
	}
	for _, clip := range topClips {
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

func (a *Activity) GetDownloadLinks(ctx context.Context, input GetDownloadLinksInput) (*GetDownloadLinksOutput, error) {
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
