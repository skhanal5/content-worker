package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

type EditSingleWorkflowInput struct {
	InputPath string `json:"input_path,required"`
	OutputPath string `json:"output_path,required"`
	Strategy string `json:"strategy,required"`
	Title string `json:"title,required"`
}

func EditSingleWorkflow(ctx workflow.Context, input EditSingleWorkflowInput) error {

	ao := workflow.ActivityOptions{
		RetryPolicy: retryPolicy,
        StartToCloseTimeout:    time.Second * 180,
        HeartbeatTimeout:       time.Second * 15,
	}
	
	ctx = workflow.WithActivityOptions(ctx, ao)
	var a activity.Activity
	log.Printf("Kicking off Single Edit Workflow with payload: %s", input)
	
	editInput := activity.EditVideoInput{
		InputPath: input.InputPath,
		OutputPath:  input.OutputPath,
		Strategy: activity.VideoStrategyType(input.Strategy),
		Title: input.Title,
	}

	err := workflow.ExecuteActivity(ctx, a.EditVideo, editInput).Get(ctx, nil)
	if err != nil {
		return err
	}
	log.Println("Completed Single Edit Workflow")
	return nil
}
