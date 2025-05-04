package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"go.temporal.io/sdk/workflow"
)

type DownloadClipInput struct {
	ID  string `json:"id,required"`
	URL string `json:"url,required"`
}

type DownloadClipOutput struct {	
	OutputPath string `json:"output_path,omitempty"`
}


func DownloadClipWorkflow(ctx workflow.Context, input DownloadClipInput) (*DownloadClipOutput, error) {
	ctx = workflow.WithActivityOptions(ctx, ActivityOptions)

	var a activity.Activity

	downloadInput := activity.DownloadClipInput{
		ID:  input.ID,
		URL: input.URL,
	}
	output := &DownloadClipOutput{}
	err := workflow.ExecuteActivity(ctx, a.DownloadClip, downloadInput).Get(ctx, &output.OutputPath)
	if err != nil {
		return nil, err
	}
	return output, nil
}