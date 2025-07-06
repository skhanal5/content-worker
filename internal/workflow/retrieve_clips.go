package workflow

import (
	"clip-farmer-workflow/internal/activity"

	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

type RetrieveClipsWorkflowInput struct {
	Streamer string `json:"streamer"`
	Limit    int    `json:"limit"`
	Filter   string `json:"filter"`
}

func RetrieveClipsWorkflow(ctx workflow.Context, input RetrieveClipsWorkflowInput) error {
	ctx = workflow.WithActivityOptions(ctx, ActivityOptions)

	var a activity.Activity

	var getClipSlugsOutput activity.GetClipSlugsOutput
	getClipSlugsInput := activity.GetClipSlugsInput{
		Broadcaster: input.Streamer,
		Limit:       input.Limit,
		Filter:      input.Filter,
	}

	err := workflow.ExecuteActivity(ctx, a.GetClipSlugs, getClipSlugsInput).Get(ctx, &getClipSlugsOutput)
	if err != nil {
		return err
	}

	var getDownloadLinks activity.GetDownloadLinksOutput
	err = workflow.ExecuteActivity(ctx, a.GetDownloadLinks, activity.GetDownloadLinksInput(getClipSlugsOutput)).Get(ctx, &getDownloadLinks)
	if err != nil {
		return err
	}

	for _, url := range getDownloadLinks.DownloadLinks {
		input := activity.DownloadClipInput{
			Streamer: input.Streamer,
			ClipID:   uuid.New().String(),
			ClipURL:  url,
		}
		err = workflow.ExecuteActivity(ctx, a.DownloadClip, input).Get(ctx, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
