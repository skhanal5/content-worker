package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"fmt"
	"strings"

	"go.temporal.io/sdk/workflow"
)

type PublishClipsInput struct {
	Streamer string `json:"streamer,required"`
}

// Note: Temporal doesn't really recommend using child workflows for this use-case.
// I am using this workflow to orchestrate the major components of my application logic
// for now. Later, I will consolidate each child worklow into this single workflow.
func PublishClipsWorkflow(ctx workflow.Context, input PublishClipsInput) (error) {
	
	ctx = workflow.WithActivityOptions(ctx, ActivityOptions)
	logger := workflow.GetLogger(ctx)

	// inherit the same options in the child workflow
	childWorkflowOptions := workflow.ChildWorkflowOptions{}
  	ctx = workflow.WithChildOptions(ctx, childWorkflowOptions)

	childInput := RetrieveClipsInput{
		Streamer: input.Streamer,
	}
	var result RetrieveClipsOutput
	err := workflow.ExecuteChildWorkflow(ctx, RetrieveClipsWorkflow, childInput).Get(ctx, &result)
	if err != nil {
		return  err
	}
	
	var a activity.Activity
	for _, clip := range result.Clips {
		mp4Url, err := getMP4URL(clip.ThumbnailURL)
		if err != nil {
			logger.Info("Can't get mp4 url for clip", "clip_id", clip.ID, "error", err)
			continue
		}
		input := activity.DownloadClipInput{
			Streamer: input.Streamer,
			ClipID:  clip.ID,
			ClipURL: mp4Url,
		}
		output := &DownloadClipOutput{}
		err = workflow.ExecuteActivity(ctx, a.DownloadClip, input).Get(ctx, &output.OutputPath)
		if err != nil {
			return nil
		}
	}
	return nil
}


/*
	Given a thumbnail url, replace the preview portion
	of it with a ".mp4" extension
	
	Note, the thumbnail url can be in a few different formats:
		1. https://static-cdn.jtvnw.net/twitch-clips/OU9JTTVUC3JhpdNpQUuLZQ/AT-cm%7COU9JTTVUC3JhpdNpQUuLZQ-preview-480x272.jpg
		2. https://static-cdn.jtvnw.net/twitch-clips-thumbnails-prod/VictoriousArbitraryDonkeyDoubleRainbow-Fn2PCb9JMKEIrGza/cbfdf779-1f77-4f70-81fe-cdb75dc0f3b0/preview-480x272.jpg

	But the MP4 location looks like: https://clips-media-assets2.twitch.tv/-x-v1x9bhwciZ23fVnULCQ/AT-cm%7C-x-v1x9bhwciZ23fVnULCQ

	Ignoring the first format for now, since I don't have a way of getting the MP4 URL
*/
func getMP4URL(thumbnailURL string) (string, error) {
	if strings.Contains(thumbnailURL, "twitch-clips-thumbnails-prod") {
		return "", fmt.Errorf("thumbnail url is not in the expected format: %s", thumbnailURL)
	}
	index := strings.Index(thumbnailURL, "-preview")
	rawURL := thumbnailURL[:index]
    rawURL = strings.Replace(rawURL, "static-cdn.jtvnw.net/twitch-clips", "clips-media-assets2.twitch.tv", 1)
    rawURL += ".mp4"
	return rawURL, nil
}