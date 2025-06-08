package workflow

import (
	"clip-farmer-workflow/internal/activity"
	"log"
	"os"
	"time"

	"go.temporal.io/sdk/workflow"
)

type EditWorkflowInput struct {
	InputDirectory string `json:"input_directory,required"`
	OutputDirectory string `json:"output_directory,required"`
	Strategy string `json:"strategy,required"`
}


func EditWorkflow(ctx workflow.Context, input EditWorkflowInput) error {
	
	ao := workflow.ActivityOptions{
		RetryPolicy: retryPolicy,
        StartToCloseTimeout:    time.Second * 180,
        HeartbeatTimeout:       time.Second * 10,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	var a activity.Activity

	log.Println("Output Directory:", input.OutputDirectory)

	inputDir := input.InputDirectory
	outputDir := input.OutputDirectory

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Printf("Cannot find the output directory: %v", err)
		return err
	}

	filesDir, err := os.ReadDir(inputDir)
	if err != nil {
		log.Printf("Cannot find the input directory: %v", err)
		return err
	}

	for _, file := range filesDir {
		editInput := activity.EditVideoInput{
			InputPath: inputDir + "/" + file.Name(),
			OutputPath:  outputDir + "/" + file.Name(),
			Strategy: activity.VideoStrategyType(input.Strategy),
		}
		err := workflow.ExecuteActivity(ctx, a.EditVideo, editInput).Get(ctx, nil)
		if err != nil {
			return err
		}
	}
	return nil
}