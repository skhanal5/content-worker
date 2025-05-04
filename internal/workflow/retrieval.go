package workflow

import (
	"clip-farmer-workflow/internal/activity"

	"go.temporal.io/sdk/workflow"
)

type RetrieveClipsInput struct {
	Username string `json:"username,required"`
}

type RetrieveClipsOutput struct {
	ClipURLs []string `json:"clip_urls,omitempty"`
}

func RetrieveClipsWorkflow(ctx workflow.Context, input RetrieveClipsInput) (*RetrieveClipsOutput, error) {
	ctx = workflow.WithActivityOptions(ctx, ActivityOptions)

	var a activity.Activity

	var userOutput activity.GetTwitchUserOutput
	err := workflow.ExecuteActivity(ctx, a.GetTwitchUser, input.Username).Get(ctx, &userOutput)
	if err != nil {
		return nil, err
	}

	var clipsOutput RetrieveClipsOutput
	err = workflow.ExecuteActivity(ctx, a.GetClipsFromUser, userOutput.BroadcasterID).Get(ctx, &clipsOutput.ClipURLs)
	if err != nil {
		return nil, err
	}
	return &clipsOutput, err
}
