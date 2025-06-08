package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"time"

	"go.temporal.io/sdk/workflow"
)

type EditWorkflowInput struct {
	InputPath string `json:"input_path,required"`
	OutputPath string `json:"output_path,required"`
	Strategy string `json:"strategy,required"`
}


func EditWorkflow(ctx workflow.Context, input EditWorkflowInput) error {
	
	ao := workflow.ActivityOptions{
        ScheduleToCloseTimeout: time.Second * 500,
        StartToCloseTimeout:    time.Second * 180,
        HeartbeatTimeout:       time.Second * 10,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	var a activity.Activity

	editInput := activity.EditVideoInput{
		InputPath: input.InputPath,
		OutputPath:  input.OutputPath,
		Strategy: activity.VideoStrategyType(input.Strategy),
	}
	err := workflow.ExecuteActivity(ctx, a.EditVideo, editInput).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}