package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"time"

	"go.temporal.io/sdk/workflow"
)

type HelloWorldInput struct {
	Name string `json:"name,required"`
}

type HelloWorldOutput struct {
	Message string `json:"message,omitempty"`
}

func HelloWorldWorkflow(ctx workflow.Context, input HelloWorldInput) (*HelloWorldOutput, error) {
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	var output HelloWorldOutput
	err := workflow.ExecuteActivity(ctx, activity.HelloWorldActivity, input.Name).Get(ctx, &output.Message)
	if err != nil {
		return nil, err
	}
	return &output, err
}
