package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"go.temporal.io/sdk/workflow"
)

type DownloadClipInput struct {
	Streamer string `json:"streamer,required"`
	ClipID string `json:"clip_id,required"`
	ClipURL string `json:"clip_url,required"`
}

type DownloadClipOutput struct {	
	OutputPath string `json:"output_path,omitempty"`
}


func DownloadClipWorkflow(ctx workflow.Context, input DownloadClipInput) (*DownloadClipOutput, error) {
	ctx = workflow.WithActivityOptions(ctx, ActivityOptions)

	var a activity.Activity

	downloadInput := activity.DownloadClipInput{
		Streamer: input.Streamer,
		ClipID:  input.ClipID,
		ClipURL: input.ClipURL,
	}
	output := &DownloadClipOutput{}
	err := workflow.ExecuteActivity(ctx, a.DownloadClip, downloadInput).Get(ctx, &output.OutputPath)
	if err != nil {
		return nil, err
	}
	return output, nil
}