package activity

import (
	"context"
	"fmt"
	"net/url"
)

type GetClipSlugsInput struct {
	Broadcaster string
	Limit       int
	Filter      string
}

type GetClipSlugsOutput struct {
	ClipIds []string
}

func (a *Activity) GetClipSlugs(ctx context.Context, input GetClipSlugsInput) (*GetClipSlugsOutput, error) {

	clips, err := a.GetUserClips(input.Broadcaster, input.Limit, input.Filter)
	if err != nil {
		return nil, err
	}
	if clips == nil {
		return nil, fmt.Errorf("no clips found for broadcaster id: %s", input.Broadcaster)
	}

	clipEdges := clips.Data.User.Clips.Edges
	slugs := []string{}

	for _, clipEdge := range clipEdges {
		clipNode := clipEdge.Node
		slugs = append(slugs, clipNode.Slug)
	}

	return &GetClipSlugsOutput{
		ClipIds: slugs,
	}, nil

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
	// TODO: Consider if this should be multiple activity calls instead of 1 call
	for _, clipId := range input.ClipIds {
		res, err := a.GetClipInformation(clipId)
		if err != nil {
			return nil, err
		}
		if res == nil {
			return nil, fmt.Errorf("no clip found for id: %s", clipId)
		}

		clip := res.Data.Clip
		qualities := clip.VideoQualities
		bestQualityURL := qualities[0].SourceURL

		u, err := url.Parse(bestQualityURL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse sourceURL: %v", err)
		}

		q := u.Query()
		q.Set("sig", clip.PlaybackAccessToken.Signature)
		q.Set("token", clip.PlaybackAccessToken.Value)
		u.RawQuery = q.Encode()

		output.DownloadLinks = append(output.DownloadLinks, u.String())
	}
	return output, nil
}
