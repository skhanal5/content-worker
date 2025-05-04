package workflow

import (
	"clip-farmer-workflow/internal/activity"

	"go.temporal.io/sdk/workflow"
)

type RetrieveClipsInput struct {
	Username string `json:"username,required"`
}

type Clip struct {
	ID string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
	ViewCount int	`json:"view_count,omitempty"`
	Duration float32	`json:"duration,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
}	

type RetrieveClipsOutput struct {
	Clips []Clip `json:"clips,omitempty"`
}

func RetrieveClipsWorkflow(ctx workflow.Context, input RetrieveClipsInput) (*RetrieveClipsOutput, error) {
	ctx = workflow.WithActivityOptions(ctx, ActivityOptions)

	var a activity.Activity

	getTwitchUserInput := activity.GetTwitchUserInput{
		Username: input.Username,
	}
	var userOutput activity.GetTwitchUserOutput
	err := workflow.ExecuteActivity(ctx, a.GetTwitchUser, getTwitchUserInput).Get(ctx, &userOutput)
	if err != nil {
		return nil, err
	}

	getClipsInput := activity.GetClipsInput{
		BroadcasterID: userOutput.BroadcasterID,
	}
	var getClipsOutput activity.GetClipsOutput
	err = workflow.ExecuteActivity(ctx, a.GetClipsFromUser, getClipsInput).Get(ctx, &getClipsOutput)
	if err != nil {
		return nil, err
	}

	output := RetrieveClipsOutput{
		Clips: []Clip{},
	}

	for _, clip := range getClipsOutput.Clips {
		output.Clips = append(output.Clips, Clip{
			ID: clip.ID,
			Title: clip.Title,
			URL: clip.URL,
			ViewCount: clip.ViewCount,
			Duration: clip.Duration,
			CreatedAt: clip.CreatedAt,
			ThumbnailURL: clip.ThumbnailURL,
		})
	}
	return &output, err
}
