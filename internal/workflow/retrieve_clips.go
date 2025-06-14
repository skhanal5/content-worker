package workflow

import (
	"clip-farmer-workflow/internal/activity"

	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

type RetrieveClipsWorkflowInput struct {
	Streamer string `json:"streamer,required"`
}

func RetrieveClipsWorkflow(ctx workflow.Context, input RetrieveClipsWorkflowInput) (error) {
	ctx = workflow.WithActivityOptions(ctx, ActivityOptions)

	var a activity.Activity

	getTwitchUserInput := activity.GetTwitchUserInput{
		Username: input.Streamer,
	}
	var userOutput activity.GetTwitchUserOutput
	err := workflow.ExecuteActivity(ctx, a.GetTwitchUser, getTwitchUserInput).Get(ctx, &userOutput)
	if err != nil {
		return err
	}

	var getClipSlugsOutput activity.GetClipSlugsOutput
	err = workflow.ExecuteActivity(ctx, a.GetClipSlugs, activity.GetClipSlugsInput{
		BroadcasterID: userOutput.BroadcasterID,
	}).Get(ctx, &getClipSlugsOutput)
	if err != nil {
		return err
	}

	var getDownloadLinks activity.GetDownloadLinksOutput
	err = workflow.ExecuteActivity(ctx, a.GetDownloadLinks, activity.GetDownloadLinksInput{
		ClipIds: getClipSlugsOutput.ClipIds,
	}).Get(ctx, &getDownloadLinks)
	if err != nil {
		return err
	}

	
	for _, url := range getDownloadLinks.DownloadLinks {
		input := activity.DownloadClipInput{
			Streamer: input.Streamer,
			ClipID:  uuid.New().String(),
			ClipURL: url,
		}
		err = workflow.ExecuteActivity(ctx, a.DownloadClip, input).Get(ctx, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
