package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"strings"
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
	for _, clip := range result.Clips {
		mp4Url := getMP4URL(clip.ThumbnailURL)
		input := DownloadClipInput{
			ID:  clip.ID,
			URL: mp4Url,
		}
		output := &DownloadClipOutput{}
		err := workflow.ExecuteActivity(ctx, a.DownloadClip, input).Get(ctx, &output.OutputPath)
		if err != nil {
			return nil
		}
	}
	return nil
}


/*
	Given a thumbnail url, replace the preview portion
	of it with a ".mp4" extension
	i.e., https://clips-media-assets2.twitch.tv/jTk1-Xmig5ji1Dll05ivnA/AT-cm%7CjTk1-Xmig5ji1Dll05ivnA-preview-260x147.jpg
	becomes https://clips-media-assets2.twitch.tv/jTk1-Xmig5ji1Dll05ivnA/AT-cm%7CjTk1-Xmig5ji1Dll05ivnA.mp4
*/
func getMP4URL(thumbnailURL string) string {
	index := strings.Index(thumbnailURL, "-preview")
	rawUrl := thumbnailURL[:index]
	rawUrl += ".mp4"
	return rawUrl
}