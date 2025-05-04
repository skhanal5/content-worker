package workflow

import (
	"clip-farmer-workflow/internal/activity"

	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

type PublishClipsInput struct {
	Username string `json:"username,required"`
}

// Note: Temporal doesn't really recommend using child workflows for this use-case.
// I am using this workflow to orchestrate the major components of my application logic
// for now. Later, I will consolidate each child worklow into this single workflow.
func PublishClipsWorkflow(ctx workflow.Context, input PublishClipsInput) (error) {
	
	ctx = workflow.WithActivityOptions(ctx, ActivityOptions)

	// inherit the same options in the child workflow
	childWorkflowOptions := workflow.ChildWorkflowOptions{}
  	ctx = workflow.WithChildOptions(ctx, childWorkflowOptions)

	childInput := RetrieveClipsInput{
		Username: input.Username,
	}
	var result RetrieveClipsOutput
	err := workflow.ExecuteChildWorkflow(ctx, RetrieveClipsWorkflow, childInput).Get(ctx, &result)
	if err != nil {
		return  err
	}
	
	var a activity.Activity
	for _, clipURL := range result.ClipURLs {
		input := DownloadClipInput{
			ID:  uuid.New().String(), // TODO: use original clip id or title name
			URL: clipURL,
		}
		output := &DownloadClipOutput{}
		err := workflow.ExecuteActivity(ctx, a.DownloadClip, input).Get(ctx, &output.OutputPath)
		if err != nil {
			return nil
		}
	}
	return nil
}